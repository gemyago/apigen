package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() BasicStringEnum

type StringTypesStringTypesEnumsInlineEnumParam string

// List of StringTypesStringTypesEnumsInlineEnumParam values.
const (
	StringTypesStringTypesEnumsInlineEnumParamVALUE1 StringTypesStringTypesEnumsInlineEnumParam = "VALUE1"
	StringTypesStringTypesEnumsInlineEnumParamVALUE2 StringTypesStringTypesEnumsInlineEnumParam = "VALUE2"
	StringTypesStringTypesEnumsInlineEnumParamVALUE3 StringTypesStringTypesEnumsInlineEnumParam = "VALUE3"
)

func(v StringTypesStringTypesEnumsInlineEnumParam) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamVALUE1
}

func(v StringTypesStringTypesEnumsInlineEnumParam) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamVALUE2
}

func(v StringTypesStringTypesEnumsInlineEnumParam) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamVALUE3
}

func(v StringTypesStringTypesEnumsInlineEnumParam) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsInlineEnumParam interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsInlineEnumParam(v assignableStringTypesStringTypesEnumsInlineEnumParam) (StringTypesStringTypesEnumsInlineEnumParam) {
	return StringTypesStringTypesEnumsInlineEnumParam(v.String())
}

func ParseStringTypesStringTypesEnumsInlineEnumParam(str string, target *StringTypesStringTypesEnumsInlineEnumParam) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsInlineEnumParamVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsInlineEnumParamVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsInlineEnumParamVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsInlineEnumParam value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsInlineEnumParam) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsInlineEnumParam(str, v)
}

// All allowed values of StringTypesStringTypesEnumsInlineEnumParam enum.
var AllowableStringTypesStringTypesEnumsInlineEnumParamValues = []StringTypesStringTypesEnumsInlineEnumParam{
	StringTypesStringTypesEnumsInlineEnumParamVALUE1,
	StringTypesStringTypesEnumsInlineEnumParamVALUE2,
	StringTypesStringTypesEnumsInlineEnumParamVALUE3,
}

type StringTypesStringTypesEnumsNullableInlineEnumParam string

// List of StringTypesStringTypesEnumsNullableInlineEnumParam values.
const (
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE1 StringTypesStringTypesEnumsNullableInlineEnumParam = "VALUE1"
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE2 StringTypesStringTypesEnumsNullableInlineEnumParam = "VALUE2"
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE3 StringTypesStringTypesEnumsNullableInlineEnumParam = "VALUE3"
)

func(v StringTypesStringTypesEnumsNullableInlineEnumParam) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamVALUE1
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParam) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamVALUE2
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParam) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamVALUE3
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParam) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsNullableInlineEnumParam interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsNullableInlineEnumParam(v assignableStringTypesStringTypesEnumsNullableInlineEnumParam) (StringTypesStringTypesEnumsNullableInlineEnumParam) {
	return StringTypesStringTypesEnumsNullableInlineEnumParam(v.String())
}

func ParseStringTypesStringTypesEnumsNullableInlineEnumParam(str string, target *StringTypesStringTypesEnumsNullableInlineEnumParam) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsNullableInlineEnumParam value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsNullableInlineEnumParam) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsNullableInlineEnumParam(str, v)
}

// All allowed values of StringTypesStringTypesEnumsNullableInlineEnumParam enum.
var AllowableStringTypesStringTypesEnumsNullableInlineEnumParamValues = []StringTypesStringTypesEnumsNullableInlineEnumParam{
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE1,
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE2,
	StringTypesStringTypesEnumsNullableInlineEnumParamVALUE3,
}

type StringTypesStringTypesEnumsInlineEnumParamInQuery string

// List of StringTypesStringTypesEnumsInlineEnumParamInQuery values.
const (
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE1 StringTypesStringTypesEnumsInlineEnumParamInQuery = "VALUE1"
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE2 StringTypesStringTypesEnumsInlineEnumParamInQuery = "VALUE2"
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE3 StringTypesStringTypesEnumsInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesStringTypesEnumsInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE1
}

func(v StringTypesStringTypesEnumsInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE2
}

func(v StringTypesStringTypesEnumsInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE3
}

func(v StringTypesStringTypesEnumsInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsInlineEnumParamInQuery(v assignableStringTypesStringTypesEnumsInlineEnumParamInQuery) (StringTypesStringTypesEnumsInlineEnumParamInQuery) {
	return StringTypesStringTypesEnumsInlineEnumParamInQuery(v.String())
}

func ParseStringTypesStringTypesEnumsInlineEnumParamInQuery(str string, target *StringTypesStringTypesEnumsInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesStringTypesEnumsInlineEnumParamInQuery enum.
var AllowableStringTypesStringTypesEnumsInlineEnumParamInQueryValues = []StringTypesStringTypesEnumsInlineEnumParamInQuery{
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE1,
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE2,
	StringTypesStringTypesEnumsInlineEnumParamInQueryVALUE3,
}

type StringTypesStringTypesEnumsNullableInlineEnumParamInQuery string

// List of StringTypesStringTypesEnumsNullableInlineEnumParamInQuery values.
const (
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE1 StringTypesStringTypesEnumsNullableInlineEnumParamInQuery = "VALUE1"
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE2 StringTypesStringTypesEnumsNullableInlineEnumParamInQuery = "VALUE2"
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE3 StringTypesStringTypesEnumsNullableInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE1
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE2
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE3
}

func(v StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsNullableInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsNullableInlineEnumParamInQuery(v assignableStringTypesStringTypesEnumsNullableInlineEnumParamInQuery) (StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) {
	return StringTypesStringTypesEnumsNullableInlineEnumParamInQuery(v.String())
}

func ParseStringTypesStringTypesEnumsNullableInlineEnumParamInQuery(str string, target *StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsNullableInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsNullableInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsNullableInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesStringTypesEnumsNullableInlineEnumParamInQuery enum.
var AllowableStringTypesStringTypesEnumsNullableInlineEnumParamInQueryValues = []StringTypesStringTypesEnumsNullableInlineEnumParamInQuery{
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE1,
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE2,
	StringTypesStringTypesEnumsNullableInlineEnumParamInQueryVALUE3,
}

type StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery string

// List of StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery values.
const (
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE1 StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery = "VALUE1"
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE2 StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery = "VALUE2"
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE3 StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE1
}

func(v StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE2
}

func(v StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE3
}

func(v StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery(v assignableStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) (StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) {
	return StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery(v.String())
}

func ParseStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery(str string, target *StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery enum.
var AllowableStringTypesStringTypesEnumsOptionalInlineEnumParamInQueryValues = []StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery{
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE1,
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE2,
	StringTypesStringTypesEnumsOptionalInlineEnumParamInQueryVALUE3,
}

type StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery string

// List of StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery values.
const (
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE1 StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery = "VALUE1"
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE2 StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery = "VALUE2"
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE3 StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE1
}

func(v StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE2
}

func(v StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE3
}

func(v StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery(v assignableStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) (StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) {
	return StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery(v.String())
}

func ParseStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery(str string, target *StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery enum.
var AllowableStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryValues = []StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery{
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE1,
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE2,
	StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryVALUE3,
}

// StringTypesStringTypesArrayItemsRangeValidationRequest represents params for stringTypesArrayItemsRangeValidation operation
//
// Request: POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArrayItemsRangeValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesArrayItemsRangeValidationRequest
}

// StringTypesStringTypesArraysParsingRequest represents params for stringTypesArraysParsing operation
//
// Request: POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArraysParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesArraysParsingRequest
}

// StringTypesStringTypesEnumsRequest represents params for stringTypesEnums operation
//
// Request: POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}.
type StringTypesStringTypesEnumsRequest struct {
	// InlineEnumParam is parsed from request path and declared as inlineEnumParam.
	InlineEnumParam StringTypesStringTypesEnumsInlineEnumParam
	// NullableInlineEnumParam is parsed from request path and declared as nullableInlineEnumParam.
	NullableInlineEnumParam *StringTypesStringTypesEnumsNullableInlineEnumParam
	// RefEnumParam is parsed from request path and declared as refEnumParam.
	RefEnumParam BasicStringEnum
	// NullableRefEnumParam is parsed from request path and declared as nullableRefEnumParam.
	NullableRefEnumParam *NullableStringEnum
	// InlineEnumParamInQuery is parsed from request query and declared as inlineEnumParamInQuery.
	InlineEnumParamInQuery StringTypesStringTypesEnumsInlineEnumParamInQuery
	// NullableInlineEnumParamInQuery is parsed from request query and declared as nullableInlineEnumParamInQuery.
	NullableInlineEnumParamInQuery *StringTypesStringTypesEnumsNullableInlineEnumParamInQuery
	// RefEnumParamInQuery is parsed from request query and declared as refEnumParamInQuery.
	RefEnumParamInQuery BasicStringEnum
	// NullableRefEnumParamInQuery is parsed from request query and declared as nullableRefEnumParamInQuery.
	NullableRefEnumParamInQuery *NullableStringEnum
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesEnumsRequest
	// OptionalInlineEnumParamInQuery is parsed from request query and declared as optionalInlineEnumParamInQuery.
	OptionalInlineEnumParamInQuery StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery
	// OptionalNullableInlineEnumParamInQuery is parsed from request query and declared as optionalNullableInlineEnumParamInQuery.
	OptionalNullableInlineEnumParamInQuery *StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery
	// OptionalRefEnumParamInQuery is parsed from request query and declared as optionalRefEnumParamInQuery.
	OptionalRefEnumParamInQuery BasicStringEnum
	// OptionalNullableRefEnumParamInQuery is parsed from request query and declared as optionalNullableRefEnumParamInQuery.
	OptionalNullableRefEnumParamInQuery *NullableStringEnum
}

// StringTypesStringTypesNullableArrayItemsRequest represents params for stringTypesNullableArrayItems operation
//
// Request: POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableArrayItemsRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []*string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []*string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []*time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []*time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []*string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []*string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []*string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []*time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []*time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []*string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesNullableArrayItemsRequest
}

// StringTypesStringTypesNullableParsingRequest represents params for stringTypesNullableParsing operation
//
// Request: POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr *string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr *string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr *time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr *time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr *string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery *string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery *string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery *time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery *time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery *string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesNullableParsingRequest
}

// StringTypesStringTypesNullableRequiredValidationRequest represents params for stringTypesNullableRequiredValidation operation
//
// Request: POST /string-types/nullable-required-validation.
type StringTypesStringTypesNullableRequiredValidationRequest struct {
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery *string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesNullableRequiredValidationRequest
	// OptionalUnformattedStrInQuery is parsed from request query and declared as optionalUnformattedStrInQuery.
	OptionalUnformattedStrInQuery *string
}

// StringTypesStringTypesParsingRequest represents params for stringTypesParsing operation
//
// Request: POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesParsingRequest
}

// StringTypesStringTypesPatternValidationRequest represents params for stringTypesPatternValidation operation
//
// Request: POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}.
type StringTypesStringTypesPatternValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesPatternValidationRequest
}

// StringTypesStringTypesRangeValidationRequest represents params for stringTypesRangeValidation operation
//
// Request: POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesRangeValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesRangeValidationRequest
}

// StringTypesStringTypesRequiredValidationRequest represents params for stringTypesRequiredValidation operation
//
// Request: POST /string-types/required-validation.
type StringTypesStringTypesRequiredValidationRequest struct {
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesRequiredValidationRequest
	// OptionalUnformattedStrInQuery is parsed from request query and declared as optionalUnformattedStrInQuery.
	OptionalUnformattedStrInQuery string
	// OptionalCustomFormatStrInQuery is parsed from request query and declared as optionalCustomFormatStrInQuery.
	OptionalCustomFormatStrInQuery string
	// OptionalDateStrInQuery is parsed from request query and declared as optionalDateStrInQuery.
	OptionalDateStrInQuery time.Time
	// OptionalDateTimeStrInQuery is parsed from request query and declared as optionalDateTimeStrInQuery.
	OptionalDateTimeStrInQuery time.Time
	// OptionalByteStrInQuery is parsed from request query and declared as optionalByteStrInQuery.
	OptionalByteStrInQuery string
}

type stringTypesControllerBuilder struct {
	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationRequest,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation genericHandlerBuilder[
		*StringTypesStringTypesArrayItemsRangeValidationRequest,
		void,
		func(context.Context, *StringTypesStringTypesArrayItemsRangeValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesArrayItemsRangeValidationRequest) (error),
	]

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingRequest,
	//
	// Response type: none
	StringTypesArraysParsing genericHandlerBuilder[
		*StringTypesStringTypesArraysParsingRequest,
		void,
		func(context.Context, *StringTypesStringTypesArraysParsingRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesArraysParsingRequest) (error),
	]

	// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
	//
	// Request type: StringTypesStringTypesEnumsRequest,
	//
	// Response type: none
	StringTypesEnums genericHandlerBuilder[
		*StringTypesStringTypesEnumsRequest,
		void,
		func(context.Context, *StringTypesStringTypesEnumsRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesEnumsRequest) (error),
	]

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsRequest,
	//
	// Response type: none
	StringTypesNullableArrayItems genericHandlerBuilder[
		*StringTypesStringTypesNullableArrayItemsRequest,
		void,
		func(context.Context, *StringTypesStringTypesNullableArrayItemsRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesNullableArrayItemsRequest) (error),
	]

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	StringTypesNullableParsing genericHandlerBuilder[
		*StringTypesStringTypesNullableParsingRequest,
		void,
		func(context.Context, *StringTypesStringTypesNullableParsingRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesNullableParsingRequest) (error),
	]

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationRequest,
	//
	// Response type: none
	StringTypesNullableRequiredValidation genericHandlerBuilder[
		*StringTypesStringTypesNullableRequiredValidationRequest,
		void,
		func(context.Context, *StringTypesStringTypesNullableRequiredValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesNullableRequiredValidationRequest) (error),
	]

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	StringTypesParsing genericHandlerBuilder[
		*StringTypesStringTypesParsingRequest,
		void,
		func(context.Context, *StringTypesStringTypesParsingRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesParsingRequest) (error),
	]

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	StringTypesPatternValidation genericHandlerBuilder[
		*StringTypesStringTypesPatternValidationRequest,
		void,
		func(context.Context, *StringTypesStringTypesPatternValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesPatternValidationRequest) (error),
	]

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	StringTypesRangeValidation genericHandlerBuilder[
		*StringTypesStringTypesRangeValidationRequest,
		void,
		func(context.Context, *StringTypesStringTypesRangeValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesRangeValidationRequest) (error),
	]

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	StringTypesRequiredValidation genericHandlerBuilder[
		*StringTypesStringTypesRequiredValidationRequest,
		void,
		func(context.Context, *StringTypesStringTypesRequiredValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *StringTypesStringTypesRequiredValidationRequest) (error),
	]
}

func newStringTypesControllerBuilder(app *RootHandler) *stringTypesControllerBuilder {
	return &stringTypesControllerBuilder{
		// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesArrayItemsRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesArrayItemsRangeValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesArrayItemsRangeValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesArrayItemsRangeValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesArrayItemsRangeValidation(app),
			},
		),

		// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesArraysParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesArraysParsingRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesArraysParsingRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesArraysParsingRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesArraysParsing(app),
			},
		),

		// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
		StringTypesEnums: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesEnumsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesEnumsRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesEnumsRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesEnums(app),
			},
		),

		// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesNullableArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableArrayItemsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableArrayItemsRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableArrayItemsRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableArrayItems(app),
			},
		),

		// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesNullableParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableParsingRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableParsingRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableParsingRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableParsing(app),
			},
		),

		// POST /string-types/nullable-required-validation
		StringTypesNullableRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableRequiredValidation(app),
			},
		),

		// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesParsingRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesParsingRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesParsingRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesParsing(app),
			},
		),

		// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
		StringTypesPatternValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesPatternValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesPatternValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesPatternValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesPatternValidation(app),
			},
		),

		// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesRangeValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesRangeValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesRangeValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesRangeValidation(app),
			},
		),

		// POST /string-types/required-validation
		StringTypesRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesRequiredValidation(app),
			},
		),
	}
}

type StringTypesController interface {
	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationRequest,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesArrayItemsRangeValidationRequest,
	]) http.Handler

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingRequest,
	//
	// Response type: none
	StringTypesArraysParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesArraysParsingRequest,
	]) http.Handler

	// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
	//
	// Request type: StringTypesStringTypesEnumsRequest,
	//
	// Response type: none
	StringTypesEnums(NoResponseHandlerBuilder[
		*StringTypesStringTypesEnumsRequest,
	]) http.Handler

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsRequest,
	//
	// Response type: none
	StringTypesNullableArrayItems(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableArrayItemsRequest,
	]) http.Handler

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	StringTypesNullableParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableParsingRequest,
	]) http.Handler

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationRequest,
	//
	// Response type: none
	StringTypesNullableRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableRequiredValidationRequest,
	]) http.Handler

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	StringTypesParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesParsingRequest,
	]) http.Handler

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	StringTypesPatternValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesPatternValidationRequest,
	]) http.Handler

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	StringTypesRangeValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesRangeValidationRequest,
	]) http.Handler

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	StringTypesRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesRequiredValidationRequest,
	]) http.Handler
}

// RegisterStringTypesRoutes will attach the following routes to the root handler:
// 
// - POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
// 
// - POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/nullable-required-validation
// 
// - POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
// 
// - POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/required-validation
// 
// Routes will use provided controller to handle requests.
func RegisterStringTypesRoutes(rootHandler *RootHandler, controller StringTypesController) {
	builder := newStringTypesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArrayItemsRangeValidation(builder.StringTypesArrayItemsRangeValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArraysParsing(builder.StringTypesArraysParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}", controller.StringTypesEnums(builder.StringTypesEnums))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableArrayItems(builder.StringTypesNullableArrayItems))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableParsing(builder.StringTypesNullableParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-required-validation", controller.StringTypesNullableRequiredValidation(builder.StringTypesNullableRequiredValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(builder.StringTypesParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}", controller.StringTypesPatternValidation(builder.StringTypesPatternValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesRangeValidation(builder.StringTypesRangeValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/required-validation", controller.StringTypesRequiredValidation(builder.StringTypesRequiredValidation))
}