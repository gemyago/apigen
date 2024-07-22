package internal

import (
	"errors"
	"fmt"
	"regexp"

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
	Field    string `json:"field"`
	Location string `json:"location"`
	Err      error  `json:"-"`
	Code     string `json:"code"`
}

func (be FieldBindingError) Error() string {
	return fmt.Sprintf("field %s (in %s) code=%s, error: %v", be.Field, be.Location, be.Code, be.Err)
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
	errors []FieldBindingError
}

func (c *BindingContext) AppendFieldError(err FieldBindingError) {
	c.errors = append(c.errors, err)
}

func (c BindingContext) AggregatedError() error {
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
	field string,
	location string,
	value TValue,
)

type SimpleFieldValidatorOpts struct {
	ProhibitEmptyValue bool
}

func NewSimpleFieldValidator[
	TValue any,
](opts SimpleFieldValidatorOpts, validators ...ValueValidator[TValue]) FieldValidator[TValue] {
	return func(
		bindingCtx *BindingContext,
		field string,
		location string,
		value TValue,
	) {
		for _, v := range validators {
			if err := v(value); err != nil {
				errCode := ErrInvalidValue
				errors.As(err, &errCode)
				bindingCtx.AppendFieldError(FieldBindingError{
					Field:    field,
					Location: location,
					Code:     errCode.Error(),
					Err:      err,
				})
				return
			}
		}
	}
}
