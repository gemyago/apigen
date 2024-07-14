package controllers

import (
	"context"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type testCall[TParams any] struct {
	params TParams
}

type numericTypesControllerTestActions struct {
	numericTypesParsing         []testCall[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesRangeValidation []testCall[*handlers.NumericTypesNumericTypesRangeValidationRequest]
}

func (c *numericTypesControllerTestActions) NumericTypesParsing(
	ctx context.Context, params *handlers.NumericTypesNumericTypesParsingRequest,
) error {
	c.numericTypesParsing = append(c.numericTypesParsing, testCall[*handlers.NumericTypesNumericTypesParsingRequest]{
		params: params,
	})
	return nil
}

func (c *numericTypesControllerTestActions) NumericTypesRangeValidation(
	ctx context.Context, params *handlers.NumericTypesNumericTypesRangeValidationRequest,
) error {
	c.numericTypesRangeValidation = append(c.numericTypesRangeValidation, testCall[*handlers.NumericTypesNumericTypesRangeValidationRequest]{
		params: params,
	})
	return nil
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.NumericTypesParsing).
		HandleNumericTypesRangeValidation.With(testActions.NumericTypesRangeValidation).
		Finalize()
}
