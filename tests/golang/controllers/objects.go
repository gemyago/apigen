package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type objectsControllerTestActions struct {
	objectsArrayParsingBodyDirect mockAction[*handlers.ObjectsObjectsArrayParsingBodyDirectRequest]
	objectsArrayParsingBodyNested mockAction[*handlers.ObjectsObjectsArrayParsingBodyNestedRequest]
	objectsNested                 mockAction[*handlers.ObjectsObjectsNestedRequest]
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
		HandleObjectsArrayParsingBodyDirect.With(testActions.objectsArrayParsingBodyDirect.action).
		HandleObjectsArrayParsingBodyNested.With(testActions.objectsArrayParsingBodyNested.action).
		HandleObjectsNested.With(testActions.objectsNested.action).
		HandleObjectsNullableOptionalBody.With(testActions.objectsNullableOptionalBody.action).
		HandleObjectsNullableRequiredBody.With(testActions.objectsNullableRequiredBody.action).
		HandleObjectsOptionalBody.With(testActions.objectsOptionalBody.action).
		HandleObjectsRequiredBody.With(testActions.objectsRequiredBody.action).
		HandleObjectsRequiredNestedObjects.With(testActions.objectsRequiredNestedObjects.action).
		Finalize()
}
