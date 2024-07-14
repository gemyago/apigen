package controllers

import (
	"context"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type testCall[TParams any] struct {
	params TParams
}

type testAction[TParams any] struct {
	calls []testCall[TParams]
}

func (c *testAction[TParams]) action(
	_ context.Context, params TParams,
) error {
	c.calls = append(c.calls, testCall[TParams]{
		params: params,
	})
	return nil
}

type numericTypesControllerTestActions struct {
	numericTypesParsing                  testAction[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesRangeValidation          testAction[*handlers.NumericTypesNumericTypesRangeValidationRequest]
	numericTypesRangeValidationExclusive testAction[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest]
	numericTypesRequiredValidation       testAction[*handlers.NumericTypesNumericTypesRequiredValidationRequest]
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.numericTypesParsing.action).
		HandleNumericTypesRangeValidation.With(testActions.numericTypesRangeValidation.action).
		HandleNumericTypesRangeValidationExclusive.With(testActions.numericTypesRangeValidationExclusive.action).
		HandleNumericTypesRequiredValidation.With(testActions.numericTypesRequiredValidation.action).
		Finalize()
}
