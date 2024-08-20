package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

type BindingError string

const (
	// ErrBadValueFormat error means data provided can not be parsed to a target type.
	ErrBadValueFormat BindingError = "BAD_FORMAT"

	// ErrValueRequired error code indicates that the required value has not been provided.
	ErrValueRequired BindingError = "INVALID_REQUIRED"

	// ErrInvalidValueOutOfRange error code indicates that the value is out of range of allowable values
	// this is usually when number is out of min/max range, or string is outside of limits.
	ErrInvalidValueOutOfRange BindingError = "INVALID_OUT_OF_RANGE"

	// ErrInvalidValue error code a generic validation error.
	ErrInvalidValue BindingError = "INVALID"
)

func (c BindingError) Error() string {
	return string(c)
}

// FieldBindingError occurs at parsing/validation stage and holds
// context on field that the error is related to.
type FieldBindingError struct {
	Location string `json:"location"`
	Err      error  `json:"-"`
	Code     string `json:"code"`
}

func (be FieldBindingError) Error() string {
	return fmt.Sprintf("field %s code=%s, error: %v", be.Location, be.Code, be.Err)
}

type AggregatedBindingError struct {
	Errors []FieldBindingError `json:"errors"`
}

func (c AggregatedBindingError) Error() string {
	errs := make([]error, len(c.Errors))
	for i, err := range c.Errors {
		errs[i] = err
	}
	return errors.Join(errs...).Error()
}

type BindingContext struct {
	parent       *BindingContext
	field        string
	memoizedPath string
	errors       []FieldBindingError
}

func (c *BindingContext) Fork(field string) *BindingContext {
	return &BindingContext{
		parent: c,
		field:  field,
	}
}

func (c *BindingContext) BuildPath() string {
	if c.memoizedPath != "" {
		return c.memoizedPath
	}
	if c.parent == nil {
		return c.field
	}
	parentPath := c.parent.BuildPath()
	if parentPath == "" {
		c.memoizedPath = c.field
	} else {
		c.memoizedPath = parentPath + "." + c.field
	}
	return c.memoizedPath
}

func (c *BindingContext) AppendFieldError(err FieldBindingError) {
	if c.parent != nil {
		c.parent.AppendFieldError(err)
		return
	}
	c.errors = append(c.errors, err)
}

func (c BindingContext) AggregatedError() error {
	if c.parent != nil {
		return c.parent.AggregatedError()
	}
	if len(c.errors) == 0 {
		return nil
	}
	return AggregatedBindingError{Errors: c.errors}
}

type OptionalVal[TVal any] struct {
	Value    TVal
	Assigned bool
}

type ValueValidator[TTargetVal any] func(TTargetVal) error

// EnsureNonDefault will validate if given value is non default for given type.
//
// There is no easy way to make a truly required validation (e.g if field is present)
// without a custom marshaler and shadow models, which will impact performance.
// So keeping a non default validation as a reasonable tradeoff.
func EnsureNonDefault[TTargetVal comparable](val TTargetVal) error {
	var empty TTargetVal
	if val == empty {
		return fmt.Errorf("provided value %v is default for given type and considered empty: %w", val, ErrValueRequired)
	}
	return nil
}

var _ = EnsureNonDefault[int]

// EnsureArrayFieldRequired will validate if given array is not empty.
func EnsureArrayFieldRequired[TTargetVal any](val []TTargetVal) error {
	if len(val) == 0 {
		return fmt.Errorf("provided array is empty: %w", ErrValueRequired)
	}
	return nil
}

var _ = EnsureArrayFieldRequired[int]

func SkipNullValidator[TTargetVal any](target ValueValidator[TTargetVal]) ValueValidator[*TTargetVal] {
	return func(tv *TTargetVal) error {
		if tv == nil {
			return nil
		}

		return target(*tv)
	}
}

func SkipNullFieldValidator[TTargetVal any](target FieldValidator[*TTargetVal]) FieldValidator[*TTargetVal] {
	return func(bindingCtx *BindingContext, value *TTargetVal) {
		if value == nil {
			return
		}

		target(bindingCtx, value)
	}
}

func NewMinMaxValueValidator[TTargetVal constraints.Ordered](
	threshold TTargetVal,
	exclusive bool,
	isMin bool,
) ValueValidator[TTargetVal] {
	return func(tv TTargetVal) error {
		// From OpenAPI spec:
		// exclusiveMinimum: false or not included	value ≥ minimum
		// exclusiveMinimum: true	value > minimum
		// exclusiveMaximum: false or not included	value ≤ maximum
		// exclusiveMaximum: true	value < maximum

		if isMin && ((exclusive && tv <= threshold) || (!exclusive && tv < threshold)) {
			return fmt.Errorf("value %v is less than minimum %v: %w", tv, threshold, ErrInvalidValueOutOfRange)
		}
		if !isMin && ((exclusive && tv >= threshold) || (!exclusive && tv > threshold)) {
			return fmt.Errorf("value %v is greater than maximum %v: %w", tv, threshold, ErrInvalidValueOutOfRange)
		}

		return nil
	}
}

func NewMinMaxLengthValidator[TTargetVal string](
	threshold int,
	isMin bool,
) ValueValidator[TTargetVal] {
	return func(tv TTargetVal) error {
		targetLen := len(tv)
		if isMin && targetLen < threshold {
			return fmt.Errorf(
				"value %v has length (%d) less than minimum %v: %w",
				tv, targetLen, threshold, ErrInvalidValueOutOfRange,
			)
		}
		if !isMin && targetLen > threshold {
			return fmt.Errorf(
				"value %v has length (%d) more than maximum %v: %w", tv,
				targetLen, threshold, ErrInvalidValueOutOfRange,
			)
		}

		return nil
	}
}

func NewPatternValidator[TTargetValue string](patternStr string) ValueValidator[string] {
	pattern := regexp.MustCompile(patternStr)
	return func(v string) error {
		if !pattern.MatchString(v) {
			return fmt.Errorf("value %v does not match pattern %v: %w", v, patternStr, ErrInvalidValue)
		}
		return nil
	}
}

type FieldValidator[TValue any] func(
	bindingCtx *BindingContext,
	value TValue,
)

func NewSimpleFieldValidator[
	TValue any,
](validators ...ValueValidator[TValue]) FieldValidator[TValue] {
	return func(
		bindingCtx *BindingContext,
		value TValue,
	) {
		for _, v := range validators {
			if err := v(value); err != nil {
				errCode := ErrInvalidValue
				errors.As(err, &errCode)
				bindingCtx.AppendFieldError(FieldBindingError{
					Location: bindingCtx.BuildPath(),
					Code:     errCode.Error(),
					Err:      err,
				})
				return
			}
		}
	}
}

type ObjectFieldValidatorParams struct {
	Nullable bool
	Required bool
}

func NewObjectFieldValidator[TTargetVal any](
	params ObjectFieldValidatorParams,
	modelValidator FieldValidator[*TTargetVal],
) FieldValidator[*TTargetVal] {
	return func(bindingCtx *BindingContext, value *TTargetVal) {
		if value == nil {
			if params.Required && !params.Nullable {
				bindingCtx.AppendFieldError(FieldBindingError{
					Location: bindingCtx.BuildPath(),
					Code:     ErrValueRequired.Error(),
				})
			}
			return
		}

		modelValidator(bindingCtx, value)
	}
}

func NewArrayValidator[
	TValue any,
](
	validateField FieldValidator[[]TValue],
	validateItems FieldValidator[TValue],
) FieldValidator[[]TValue] {
	return func(
		bindingCtx *BindingContext,
		value []TValue,
	) {
		validateField(bindingCtx, value)
		for i, v := range value {
			// TODO: Consider fmt.Stringer approach, defer conversion and benchmark if makes noticeable difference.
			validateItems(bindingCtx.Fork(strconv.Itoa(i)), v)
		}
	}
}
