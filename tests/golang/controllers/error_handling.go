package controllers

import (
	"context"
	"errors"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

func newErrorHandlingController() *handlers.ErrorHandlingController {
	return handlers.BuildErrorHandlingController().
		HandleErrorHandlingParsingErrors.With(
		func(_ context.Context, _ *handlers.ErrorHandlingErrorHandlingParsingErrorsRequest) error {
			return errors.New("not implemented")
		}).
		HandleErrorHandlingValidationErrors.With(
		func(_ context.Context, _ *handlers.ErrorHandlingErrorHandlingValidationErrorsRequest) error {
			return errors.New("not implemented")
		}).
		Finalize()
}
