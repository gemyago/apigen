package controllers

import (
	"context"
	"errors"

	"github.com/gemyago/apigen/tests/golang/routes/error_handling/handlers"
)

func newErrorHandlingController() *handlers.ErrorHandlingController {
	return handlers.BuildErrorHandlingController().
		HandleParsingErrors.With(func(_ context.Context, _ *handlers.ErrorHandlingParsingErrorsRequest) error {
		return errors.New("not implemented")
	}).Finalize()
}
