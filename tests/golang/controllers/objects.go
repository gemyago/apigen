package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type objectsControllerTestActions struct {
	objectsArrayParsingBodyDirect mockAction[*handlers.ObjectsObjectsArrayBodyDirectRequest]
	objectsArrayParsingBodyNested mockAction[*handlers.ObjectsObjectsArrayBodyNestedRequest]
	objectsDeeplyNested           mockAction[*handlers.ObjectsObjectsDeeplyNestedRequest]
	objectsNullableOptionalBody   mockAction[*handlers.ObjectsObjectsNullableOptionalBodyRequest]
	objectsNullableRequiredBody   mockAction[*handlers.ObjectsObjectsNullableRequiredBodyRequest]
	objectsOptionalBody           mockAction[*handlers.ObjectsObjectsOptionalBodyRequest]
	objectsRequiredBody           mockAction[*handlers.ObjectsObjectsRequiredBodyRequest]
	objectsRequiredNestedObjects  mockAction[*handlers.ObjectsObjectsRequiredNestedObjectsRequest]
}

func newObjectsController(
	testActions *objectsControllerTestActions,
) *handlers.ObjectsController {
	return handlers.BuildObjectsController().
		HandleObjectsArrayBodyDirect.With(testActions.objectsArrayParsingBodyDirect.action).
		HandleObjectsArrayBodyNested.With(testActions.objectsArrayParsingBodyNested.action).
		HandleObjectsDeeplyNested.With(testActions.objectsDeeplyNested.action).
		HandleObjectsNullableOptionalBody.With(testActions.objectsNullableOptionalBody.action).
		HandleObjectsNullableRequiredBody.With(testActions.objectsNullableRequiredBody.action).
		HandleObjectsOptionalBody.With(testActions.objectsOptionalBody.action).
		HandleObjectsRequiredBody.With(testActions.objectsRequiredBody.action).
		HandleObjectsRequiredNestedObjects.With(testActions.objectsRequiredNestedObjects.action).
		Finalize()
}
