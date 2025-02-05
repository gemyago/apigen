package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type errorHandlingController struct{}

func (c *errorHandlingController) ErrorHandlingParsingErrors(
	builder *handlers.ErrorHandlingControllerBuilderV2,
) http.Handler {
	return builder.ErrorHandlingParsingErrors.HandleWith(
		func(_ context.Context, _ *handlers.ErrorHandlingErrorHandlingParsingErrorsRequest) error {
			return errors.New("not implemented")
		},
	)
}

func (c *errorHandlingController) ErrorHandlingValidationErrors(
	builder *handlers.ErrorHandlingControllerBuilderV2,
) http.Handler {
	return builder.ErrorHandlingValidationErrors.HandleWith(
		func(_ context.Context, _ *handlers.ErrorHandlingErrorHandlingValidationErrorsRequest) error {
			return errors.New("not implemented")
		},
	)
}

func (c *errorHandlingController) ErrorHandlingActionErrors(
	builder *handlers.ErrorHandlingControllerBuilderV2,
) http.Handler {
	return builder.ErrorHandlingActionErrors.HandleWith(
		func(_ context.Context) error {
			return errors.New("simulated action error")
		},
	)
}

var _ handlers.ErrorHandlingControllerV2 = &errorHandlingController{}
