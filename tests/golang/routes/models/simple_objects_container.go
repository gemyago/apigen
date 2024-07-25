package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type SimpleObjectsContainer struct { 
	SimpleField1 string `json:"simpleField1"`
	SimpleField2 string `json:"simpleField2"`
	SimpleObject1 *SimpleObject `json:"simpleObject1"`
	SimpleObject2 *SimpleObject `json:"simpleObject2"`
	SimpleNullableObject1 *SimpleNullableObject `json:"simpleNullableObject1"`
	SimpleNullableObject2 *SimpleNullableObject `json:"simpleNullableObject2"`
	OptionalSimpleObject1 *SimpleObject `json:"optionalSimpleObject1,omitempty"`
	OptionalSimpleObject2 *SimpleObject `json:"optionalSimpleObject2,omitempty"`
	OptionalNullableSimpleObject1 *SimpleNullableObject `json:"optionalNullableSimpleObject1,omitempty"`
	OptionalNullableSimpleObject2 *SimpleNullableObject `json:"optionalNullableSimpleObject2,omitempty"`
}
