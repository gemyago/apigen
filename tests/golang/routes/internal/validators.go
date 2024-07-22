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

type OptionalVal[TVal any] struct {
	Value    TVal
	Assigned bool
}

type ValueValidator[TRawVal any, TTargetVal any] func(OptionalVal[TRawVal], TTargetVal) error

func ValidateNonEmpty[TRawVal any, TTargetVal any](rawVal OptionalVal[TRawVal], _ TTargetVal) error {
	if !rawVal.Assigned {
		return ErrValueRequired
	}
	return nil
}

var _ ValueValidator[string, string] = ValidateNonEmpty

type ModelValidationContext struct {
	Errors []error
}

type ModelValidator[TTargetVal any] func(validationCtx *ModelValidationContext, val TTargetVal)

func NewModelParamValidator[TRawVal any, TTargetVal any](
	validateModel ModelValidator[TTargetVal],
) ValueValidator[TRawVal, TTargetVal] {
	return func(_ OptionalVal[TRawVal], tv TTargetVal) error {
		validationCtx := ModelValidationContext{}
		validateModel(&validationCtx, tv)
		return errors.Join(validationCtx.Errors...)
	}
}

func NewMinMaxValueValidator[TRawVal any, TTargetVal constraints.Ordered](
	threshold TTargetVal,
	exclusive bool,
	isMin bool,
) ValueValidator[TRawVal, TTargetVal] {
	return func(ov OptionalVal[TRawVal], tv TTargetVal) error {
		if !ov.Assigned {
			return nil
		}

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

func NewMinMaxLengthValidator[TRawVal any, TTargetVal string](
	threshold int,
	isMin bool,
) ValueValidator[TRawVal, TTargetVal] {
	return func(ov OptionalVal[TRawVal], tv TTargetVal) error {
		if !ov.Assigned {
			return nil
		}

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

func NewPatternValidator[TRawVal any, TTargetValue string](patternStr string) ValueValidator[TRawVal, string] {
	pattern := regexp.MustCompile(patternStr)
	return func(ov OptionalVal[TRawVal], tv string) error {
		if !ov.Assigned {
			return nil
		}
		if !pattern.MatchString(tv) {
			return fmt.Errorf("value %v does not match pattern %v: %w", tv, patternStr, ErrInvalidValue)
		}
		return nil
	}
}

func NewCompositeValidator[
	TRawVal any, TTargetVal any,
](validators ...ValueValidator[TRawVal, TTargetVal]) ValueValidator[TRawVal, TTargetVal] {
	return func(ov OptionalVal[TRawVal], tv TTargetVal) error {
		for _, v := range validators {
			if err := v(ov, tv); err != nil {
				return err
			}
		}
		return nil
	}
}
