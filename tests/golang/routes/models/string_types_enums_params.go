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
type StringTypesEnumsParamsInlineEnumParam string

// List of StringTypesEnumsParamsInlineEnumParam values.
const (
	StringTypesEnumsParamsInlineEnumParamVALUE1 StringTypesEnumsParamsInlineEnumParam = "VALUE1"
	StringTypesEnumsParamsInlineEnumParamVALUE2 StringTypesEnumsParamsInlineEnumParam = "VALUE2"
	StringTypesEnumsParamsInlineEnumParamVALUE3 StringTypesEnumsParamsInlineEnumParam = "VALUE3"
)

func(v StringTypesEnumsParamsInlineEnumParam) IsVALUE1() bool {
  return v == StringTypesEnumsParamsInlineEnumParamVALUE1
}

func(v StringTypesEnumsParamsInlineEnumParam) IsVALUE2() bool {
  return v == StringTypesEnumsParamsInlineEnumParamVALUE2
}

func(v StringTypesEnumsParamsInlineEnumParam) IsVALUE3() bool {
  return v == StringTypesEnumsParamsInlineEnumParamVALUE3
}

func(v StringTypesEnumsParamsInlineEnumParam) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsInlineEnumParam interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsInlineEnumParam(v assignableStringTypesEnumsParamsInlineEnumParam) (StringTypesEnumsParamsInlineEnumParam) {
	return StringTypesEnumsParamsInlineEnumParam(v.String())
}

func ParseStringTypesEnumsParamsInlineEnumParam(str string, target *StringTypesEnumsParamsInlineEnumParam) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsInlineEnumParamVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsInlineEnumParamVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsInlineEnumParamVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsInlineEnumParam value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsInlineEnumParam) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsInlineEnumParam(str, v)
}

// All allowed values of StringTypesEnumsParamsInlineEnumParam enum.
var AllowableStringTypesEnumsParamsInlineEnumParamValues = []StringTypesEnumsParamsInlineEnumParam{
	StringTypesEnumsParamsInlineEnumParamVALUE1,
	StringTypesEnumsParamsInlineEnumParamVALUE2,
	StringTypesEnumsParamsInlineEnumParamVALUE3,
}
type StringTypesEnumsParamsNullableInlineEnumParam string

// List of StringTypesEnumsParamsNullableInlineEnumParam values.
const (
	StringTypesEnumsParamsNullableInlineEnumParamVALUE1 StringTypesEnumsParamsNullableInlineEnumParam = "VALUE1"
	StringTypesEnumsParamsNullableInlineEnumParamVALUE2 StringTypesEnumsParamsNullableInlineEnumParam = "VALUE2"
	StringTypesEnumsParamsNullableInlineEnumParamVALUE3 StringTypesEnumsParamsNullableInlineEnumParam = "VALUE3"
)

func(v StringTypesEnumsParamsNullableInlineEnumParam) IsVALUE1() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamVALUE1
}

func(v StringTypesEnumsParamsNullableInlineEnumParam) IsVALUE2() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamVALUE2
}

func(v StringTypesEnumsParamsNullableInlineEnumParam) IsVALUE3() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamVALUE3
}

func(v StringTypesEnumsParamsNullableInlineEnumParam) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsNullableInlineEnumParam interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsNullableInlineEnumParam(v assignableStringTypesEnumsParamsNullableInlineEnumParam) (StringTypesEnumsParamsNullableInlineEnumParam) {
	return StringTypesEnumsParamsNullableInlineEnumParam(v.String())
}

func ParseStringTypesEnumsParamsNullableInlineEnumParam(str string, target *StringTypesEnumsParamsNullableInlineEnumParam) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsNullableInlineEnumParamVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsNullableInlineEnumParamVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsNullableInlineEnumParamVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsNullableInlineEnumParam value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsNullableInlineEnumParam) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsNullableInlineEnumParam(str, v)
}

// All allowed values of StringTypesEnumsParamsNullableInlineEnumParam enum.
var AllowableStringTypesEnumsParamsNullableInlineEnumParamValues = []StringTypesEnumsParamsNullableInlineEnumParam{
	StringTypesEnumsParamsNullableInlineEnumParamVALUE1,
	StringTypesEnumsParamsNullableInlineEnumParamVALUE2,
	StringTypesEnumsParamsNullableInlineEnumParamVALUE3,
}
type StringTypesEnumsParamsInlineEnumParamInQuery string

// List of StringTypesEnumsParamsInlineEnumParamInQuery values.
const (
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE1 StringTypesEnumsParamsInlineEnumParamInQuery = "VALUE1"
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE2 StringTypesEnumsParamsInlineEnumParamInQuery = "VALUE2"
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE3 StringTypesEnumsParamsInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesEnumsParamsInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesEnumsParamsInlineEnumParamInQueryVALUE1
}

func(v StringTypesEnumsParamsInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesEnumsParamsInlineEnumParamInQueryVALUE2
}

func(v StringTypesEnumsParamsInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesEnumsParamsInlineEnumParamInQueryVALUE3
}

func(v StringTypesEnumsParamsInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsInlineEnumParamInQuery(v assignableStringTypesEnumsParamsInlineEnumParamInQuery) (StringTypesEnumsParamsInlineEnumParamInQuery) {
	return StringTypesEnumsParamsInlineEnumParamInQuery(v.String())
}

func ParseStringTypesEnumsParamsInlineEnumParamInQuery(str string, target *StringTypesEnumsParamsInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesEnumsParamsInlineEnumParamInQuery enum.
var AllowableStringTypesEnumsParamsInlineEnumParamInQueryValues = []StringTypesEnumsParamsInlineEnumParamInQuery{
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE1,
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE2,
	StringTypesEnumsParamsInlineEnumParamInQueryVALUE3,
}
type StringTypesEnumsParamsOptionalInlineEnumParamInQuery string

// List of StringTypesEnumsParamsOptionalInlineEnumParamInQuery values.
const (
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE1 StringTypesEnumsParamsOptionalInlineEnumParamInQuery = "VALUE1"
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE2 StringTypesEnumsParamsOptionalInlineEnumParamInQuery = "VALUE2"
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE3 StringTypesEnumsParamsOptionalInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesEnumsParamsOptionalInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE1
}

func(v StringTypesEnumsParamsOptionalInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE2
}

func(v StringTypesEnumsParamsOptionalInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE3
}

func(v StringTypesEnumsParamsOptionalInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsOptionalInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsOptionalInlineEnumParamInQuery(v assignableStringTypesEnumsParamsOptionalInlineEnumParamInQuery) (StringTypesEnumsParamsOptionalInlineEnumParamInQuery) {
	return StringTypesEnumsParamsOptionalInlineEnumParamInQuery(v.String())
}

func ParseStringTypesEnumsParamsOptionalInlineEnumParamInQuery(str string, target *StringTypesEnumsParamsOptionalInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsOptionalInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsOptionalInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsOptionalInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesEnumsParamsOptionalInlineEnumParamInQuery enum.
var AllowableStringTypesEnumsParamsOptionalInlineEnumParamInQueryValues = []StringTypesEnumsParamsOptionalInlineEnumParamInQuery{
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE1,
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE2,
	StringTypesEnumsParamsOptionalInlineEnumParamInQueryVALUE3,
}
type StringTypesEnumsParamsNullableInlineEnumParamInQuery string

// List of StringTypesEnumsParamsNullableInlineEnumParamInQuery values.
const (
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE1 StringTypesEnumsParamsNullableInlineEnumParamInQuery = "VALUE1"
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE2 StringTypesEnumsParamsNullableInlineEnumParamInQuery = "VALUE2"
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE3 StringTypesEnumsParamsNullableInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesEnumsParamsNullableInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE1
}

func(v StringTypesEnumsParamsNullableInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE2
}

func(v StringTypesEnumsParamsNullableInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE3
}

func(v StringTypesEnumsParamsNullableInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsNullableInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsNullableInlineEnumParamInQuery(v assignableStringTypesEnumsParamsNullableInlineEnumParamInQuery) (StringTypesEnumsParamsNullableInlineEnumParamInQuery) {
	return StringTypesEnumsParamsNullableInlineEnumParamInQuery(v.String())
}

func ParseStringTypesEnumsParamsNullableInlineEnumParamInQuery(str string, target *StringTypesEnumsParamsNullableInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsNullableInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsNullableInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsNullableInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesEnumsParamsNullableInlineEnumParamInQuery enum.
var AllowableStringTypesEnumsParamsNullableInlineEnumParamInQueryValues = []StringTypesEnumsParamsNullableInlineEnumParamInQuery{
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE1,
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE2,
	StringTypesEnumsParamsNullableInlineEnumParamInQueryVALUE3,
}
type StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery string

// List of StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery values.
const (
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE1 StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery = "VALUE1"
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE2 StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery = "VALUE2"
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE3 StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery = "VALUE3"
)

func(v StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) IsVALUE1() bool {
  return v == StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE1
}

func(v StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) IsVALUE2() bool {
  return v == StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE2
}

func(v StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) IsVALUE3() bool {
  return v == StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE3
}

func(v StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) String() string {
	return string(v)
}

type assignableStringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery(v assignableStringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) (StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) {
	return StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery(v.String())
}

func ParseStringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery(str string, target *StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE1
	case "VALUE2":
		*target = StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE2
	case "VALUE3":
		*target = StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery(str, v)
}

// All allowed values of StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery enum.
var AllowableStringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryValues = []StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery{
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE1,
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE2,
	StringTypesEnumsParamsOptionalNullableInlineEnumParamInQueryVALUE3,
}

// StringTypesEnumsParams - Parameters for the stringTypesEnums operation.
type StringTypesEnumsParams struct { 
	InlineEnumParam StringTypesEnumsParamsInlineEnumParam `json:"inlineEnumParam,omitempty"`
	NullableInlineEnumParam *StringTypesEnumsParamsNullableInlineEnumParam `json:"nullableInlineEnumParam,omitempty"`
	RefEnumParam BasicStringEnum `json:"refEnumParam,omitempty"`
	NullableRefEnumParam *NullableStringEnum `json:"nullableRefEnumParam,omitempty"`
	InlineEnumParamInQuery StringTypesEnumsParamsInlineEnumParamInQuery `json:"inlineEnumParamInQuery,omitempty"`
	OptionalInlineEnumParamInQuery StringTypesEnumsParamsOptionalInlineEnumParamInQuery `json:"optionalInlineEnumParamInQuery,omitempty"`
	NullableInlineEnumParamInQuery *StringTypesEnumsParamsNullableInlineEnumParamInQuery `json:"nullableInlineEnumParamInQuery,omitempty"`
	OptionalNullableInlineEnumParamInQuery *StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery `json:"optionalNullableInlineEnumParamInQuery,omitempty"`
	RefEnumParamInQuery BasicStringEnum `json:"refEnumParamInQuery,omitempty"`
	NullableRefEnumParamInQuery *NullableStringEnum `json:"nullableRefEnumParamInQuery,omitempty"`
	OptionalRefEnumParamInQuery BasicStringEnum `json:"optionalRefEnumParamInQuery,omitempty"`
	OptionalNullableRefEnumParamInQuery *NullableStringEnum `json:"optionalNullableRefEnumParamInQuery,omitempty"`
	Payload *StringTypesEnumsRequest `json:"payload,omitempty"`
}
