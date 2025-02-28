package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Unused imports workaround.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

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

// StringTypesStringTypesArrayItemsRangeValidationParams represents params for stringTypesArrayItemsRangeValidation operation
//
// Request: POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArrayItemsRangeValidationParams struct {
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

// StringTypesStringTypesArraysParsingParams represents params for stringTypesArraysParsing operation
//
// Request: POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArraysParsingParams struct {
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

// StringTypesStringTypesEnumsParams represents params for stringTypesEnums operation
//
// Request: POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}.
type StringTypesStringTypesEnumsParams struct {
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

// StringTypesStringTypesNullableArrayItemsParams represents params for stringTypesNullableArrayItems operation
//
// Request: POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableArrayItemsParams struct {
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

// StringTypesStringTypesNullableParsingParams represents params for stringTypesNullableParsing operation
//
// Request: POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableParsingParams struct {
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

// StringTypesStringTypesNullableRequiredValidationParams represents params for stringTypesNullableRequiredValidation operation
//
// Request: POST /string-types/nullable-required-validation.
type StringTypesStringTypesNullableRequiredValidationParams struct {
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery *string
	// Payload is parsed from request body and declared as payload.
	Payload *StringTypesNullableRequiredValidationRequest
	// OptionalUnformattedStrInQuery is parsed from request query and declared as optionalUnformattedStrInQuery.
	OptionalUnformattedStrInQuery *string
}

// StringTypesStringTypesParsingParams represents params for stringTypesParsing operation
//
// Request: POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesParsingParams struct {
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

// StringTypesStringTypesPatternValidationParams represents params for stringTypesPatternValidation operation
//
// Request: POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}.
type StringTypesStringTypesPatternValidationParams struct {
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

// StringTypesStringTypesRangeValidationParams represents params for stringTypesRangeValidation operation
//
// Request: POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesRangeValidationParams struct {
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

// StringTypesStringTypesRequiredValidationParams represents params for stringTypesRequiredValidation operation
//
// Request: POST /string-types/required-validation.
type StringTypesStringTypesRequiredValidationParams struct {
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
