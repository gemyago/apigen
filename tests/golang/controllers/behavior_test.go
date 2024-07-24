package controllers

import (
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

func TestBehavior(t *testing.T) {
	t.Run("controller builder", func(t *testing.T) {
		t.Run("should panic if actions are not initialized", func(t *testing.T) {
			assert.PanicsWithError(t, "behaviorNoParamsNoResponse action has not been initialized", func() {
				handlers.BuildBehaviorController().Finalize()
			})
		})

		t.Run("should build the controller if all actions are initialized", func(t *testing.T) {
			assert.NotPanics(t, func() {
				newBehaviorController()
			})
		})
	})
}
