package controllers

import (
	"context"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type testCall[TParams any] struct {
	params TParams
}

type numericTypesControllerTestActions struct {
	getNumberAnySimpleCalls []testCall[*handlers.NumericTypesNumberAnySimpleRequest]
}

func (c *numericTypesControllerTestActions) GetNumberAnySimple(
	ctx context.Context, params *handlers.NumericTypesNumberAnySimpleRequest,
) error {
	c.getNumberAnySimpleCalls = append(c.getNumberAnySimpleCalls, testCall[*handlers.NumericTypesNumberAnySimpleRequest]{
		params: params,
	})
	return nil
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumberAnySimple.With(testActions.GetNumberAnySimple).
		Finalize()
}
