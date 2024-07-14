package controllers

import (
	"context"
	"fmt"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type testCall[TParams any] struct {
	params TParams
}

type numericTypesControllerTestActions struct {
	numericTypesParsing []testCall[*handlers.NumericTypesNumericTypesParsingRequest]
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
	return fmt.Errorf("not implemented")
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.NumericTypesParsing).
		HandleNumericTypesRangeValidation.With(testActions.NumericTypesRangeValidation).
		Finalize()
}
