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
type StringTypesEnumsRequestInlineEnumProp string

// List of StringTypesEnumsRequestInlineEnumProp values.
const (
	StringTypesEnumsRequestInlineEnumPropVALUE1 StringTypesEnumsRequestInlineEnumProp = "VALUE1"
	StringTypesEnumsRequestInlineEnumPropVALUE2 StringTypesEnumsRequestInlineEnumProp = "VALUE2"
	StringTypesEnumsRequestInlineEnumPropVALUE3 StringTypesEnumsRequestInlineEnumProp = "VALUE3"
)

func(v StringTypesEnumsRequestInlineEnumProp) IsVALUE1() bool {
  return v == StringTypesEnumsRequestInlineEnumPropVALUE1
}

func(v StringTypesEnumsRequestInlineEnumProp) IsVALUE2() bool {
  return v == StringTypesEnumsRequestInlineEnumPropVALUE2
}

func(v StringTypesEnumsRequestInlineEnumProp) IsVALUE3() bool {
  return v == StringTypesEnumsRequestInlineEnumPropVALUE3
}

func(v StringTypesEnumsRequestInlineEnumProp) String() string {
	return string(v)
}

type assignableStringTypesEnumsRequestInlineEnumProp interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsRequestInlineEnumProp(v assignableStringTypesEnumsRequestInlineEnumProp) (StringTypesEnumsRequestInlineEnumProp) {
	return StringTypesEnumsRequestInlineEnumProp(v.String())
}

func ParseStringTypesEnumsRequestInlineEnumProp(str string, target *StringTypesEnumsRequestInlineEnumProp) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsRequestInlineEnumPropVALUE1
	case "VALUE2":
		*target = StringTypesEnumsRequestInlineEnumPropVALUE2
	case "VALUE3":
		*target = StringTypesEnumsRequestInlineEnumPropVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsRequestInlineEnumProp value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsRequestInlineEnumProp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsRequestInlineEnumProp(str, v)
}

// All allowed values of StringTypesEnumsRequestInlineEnumProp enum.
var AllowableStringTypesEnumsRequestInlineEnumPropValues = []StringTypesEnumsRequestInlineEnumProp{
	StringTypesEnumsRequestInlineEnumPropVALUE1,
	StringTypesEnumsRequestInlineEnumPropVALUE2,
	StringTypesEnumsRequestInlineEnumPropVALUE3,
}
type StringTypesEnumsRequestOptionalInlineEnumProp string

// List of StringTypesEnumsRequestOptionalInlineEnumProp values.
const (
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE1 StringTypesEnumsRequestOptionalInlineEnumProp = "VALUE1"
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE2 StringTypesEnumsRequestOptionalInlineEnumProp = "VALUE2"
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE3 StringTypesEnumsRequestOptionalInlineEnumProp = "VALUE3"
)

func(v StringTypesEnumsRequestOptionalInlineEnumProp) IsVALUE1() bool {
  return v == StringTypesEnumsRequestOptionalInlineEnumPropVALUE1
}

func(v StringTypesEnumsRequestOptionalInlineEnumProp) IsVALUE2() bool {
  return v == StringTypesEnumsRequestOptionalInlineEnumPropVALUE2
}

func(v StringTypesEnumsRequestOptionalInlineEnumProp) IsVALUE3() bool {
  return v == StringTypesEnumsRequestOptionalInlineEnumPropVALUE3
}

func(v StringTypesEnumsRequestOptionalInlineEnumProp) String() string {
	return string(v)
}

type assignableStringTypesEnumsRequestOptionalInlineEnumProp interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsRequestOptionalInlineEnumProp(v assignableStringTypesEnumsRequestOptionalInlineEnumProp) (StringTypesEnumsRequestOptionalInlineEnumProp) {
	return StringTypesEnumsRequestOptionalInlineEnumProp(v.String())
}

func ParseStringTypesEnumsRequestOptionalInlineEnumProp(str string, target *StringTypesEnumsRequestOptionalInlineEnumProp) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsRequestOptionalInlineEnumPropVALUE1
	case "VALUE2":
		*target = StringTypesEnumsRequestOptionalInlineEnumPropVALUE2
	case "VALUE3":
		*target = StringTypesEnumsRequestOptionalInlineEnumPropVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsRequestOptionalInlineEnumProp value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsRequestOptionalInlineEnumProp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsRequestOptionalInlineEnumProp(str, v)
}

// All allowed values of StringTypesEnumsRequestOptionalInlineEnumProp enum.
var AllowableStringTypesEnumsRequestOptionalInlineEnumPropValues = []StringTypesEnumsRequestOptionalInlineEnumProp{
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE1,
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE2,
	StringTypesEnumsRequestOptionalInlineEnumPropVALUE3,
}
type StringTypesEnumsRequestNullableInlineEnumProp string

// List of StringTypesEnumsRequestNullableInlineEnumProp values.
const (
	StringTypesEnumsRequestNullableInlineEnumPropVALUE1 StringTypesEnumsRequestNullableInlineEnumProp = "VALUE1"
	StringTypesEnumsRequestNullableInlineEnumPropVALUE2 StringTypesEnumsRequestNullableInlineEnumProp = "VALUE2"
	StringTypesEnumsRequestNullableInlineEnumPropVALUE3 StringTypesEnumsRequestNullableInlineEnumProp = "VALUE3"
)

func(v StringTypesEnumsRequestNullableInlineEnumProp) IsVALUE1() bool {
  return v == StringTypesEnumsRequestNullableInlineEnumPropVALUE1
}

func(v StringTypesEnumsRequestNullableInlineEnumProp) IsVALUE2() bool {
  return v == StringTypesEnumsRequestNullableInlineEnumPropVALUE2
}

func(v StringTypesEnumsRequestNullableInlineEnumProp) IsVALUE3() bool {
  return v == StringTypesEnumsRequestNullableInlineEnumPropVALUE3
}

func(v StringTypesEnumsRequestNullableInlineEnumProp) String() string {
	return string(v)
}

type assignableStringTypesEnumsRequestNullableInlineEnumProp interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsRequestNullableInlineEnumProp(v assignableStringTypesEnumsRequestNullableInlineEnumProp) (StringTypesEnumsRequestNullableInlineEnumProp) {
	return StringTypesEnumsRequestNullableInlineEnumProp(v.String())
}

func ParseStringTypesEnumsRequestNullableInlineEnumProp(str string, target *StringTypesEnumsRequestNullableInlineEnumProp) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsRequestNullableInlineEnumPropVALUE1
	case "VALUE2":
		*target = StringTypesEnumsRequestNullableInlineEnumPropVALUE2
	case "VALUE3":
		*target = StringTypesEnumsRequestNullableInlineEnumPropVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsRequestNullableInlineEnumProp value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsRequestNullableInlineEnumProp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsRequestNullableInlineEnumProp(str, v)
}

// All allowed values of StringTypesEnumsRequestNullableInlineEnumProp enum.
var AllowableStringTypesEnumsRequestNullableInlineEnumPropValues = []StringTypesEnumsRequestNullableInlineEnumProp{
	StringTypesEnumsRequestNullableInlineEnumPropVALUE1,
	StringTypesEnumsRequestNullableInlineEnumPropVALUE2,
	StringTypesEnumsRequestNullableInlineEnumPropVALUE3,
}
type StringTypesEnumsRequestOptionalNullableInlineEnumProp string

// List of StringTypesEnumsRequestOptionalNullableInlineEnumProp values.
const (
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE1 StringTypesEnumsRequestOptionalNullableInlineEnumProp = "VALUE1"
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE2 StringTypesEnumsRequestOptionalNullableInlineEnumProp = "VALUE2"
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE3 StringTypesEnumsRequestOptionalNullableInlineEnumProp = "VALUE3"
)

func(v StringTypesEnumsRequestOptionalNullableInlineEnumProp) IsVALUE1() bool {
  return v == StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE1
}

func(v StringTypesEnumsRequestOptionalNullableInlineEnumProp) IsVALUE2() bool {
  return v == StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE2
}

func(v StringTypesEnumsRequestOptionalNullableInlineEnumProp) IsVALUE3() bool {
  return v == StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE3
}

func(v StringTypesEnumsRequestOptionalNullableInlineEnumProp) String() string {
	return string(v)
}

type assignableStringTypesEnumsRequestOptionalNullableInlineEnumProp interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsStringTypesEnumsRequestOptionalNullableInlineEnumProp(v assignableStringTypesEnumsRequestOptionalNullableInlineEnumProp) (StringTypesEnumsRequestOptionalNullableInlineEnumProp) {
	return StringTypesEnumsRequestOptionalNullableInlineEnumProp(v.String())
}

func ParseStringTypesEnumsRequestOptionalNullableInlineEnumProp(str string, target *StringTypesEnumsRequestOptionalNullableInlineEnumProp) error {
	switch str {
	case "VALUE1":
		*target = StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE1
	case "VALUE2":
		*target = StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE2
	case "VALUE3":
		*target = StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE3
	default:
		return fmt.Errorf("unexpected StringTypesEnumsRequestOptionalNullableInlineEnumProp value: %s", str)
	}
	return nil
}

func (v *StringTypesEnumsRequestOptionalNullableInlineEnumProp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseStringTypesEnumsRequestOptionalNullableInlineEnumProp(str, v)
}

// All allowed values of StringTypesEnumsRequestOptionalNullableInlineEnumProp enum.
var AllowableStringTypesEnumsRequestOptionalNullableInlineEnumPropValues = []StringTypesEnumsRequestOptionalNullableInlineEnumProp{
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE1,
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE2,
	StringTypesEnumsRequestOptionalNullableInlineEnumPropVALUE3,
}

type StringTypesEnumsRequest struct { 
	InlineEnumProp StringTypesEnumsRequestInlineEnumProp `json:"inlineEnumProp"`
	OptionalInlineEnumProp StringTypesEnumsRequestOptionalInlineEnumProp `json:"optionalInlineEnumProp,omitempty"`
	NullableInlineEnumProp *StringTypesEnumsRequestNullableInlineEnumProp `json:"nullableInlineEnumProp"`
	OptionalNullableInlineEnumProp *StringTypesEnumsRequestOptionalNullableInlineEnumProp `json:"optionalNullableInlineEnumProp,omitempty"`
	RefEnumProp BasicStringEnum `json:"refEnumProp"`
	OptionalRefEnumProp BasicStringEnum `json:"optionalRefEnumProp,omitempty"`
	NullableRefEnumProp *NullableStringEnum `json:"nullableRefEnumProp"`
	OptionalNullableRefEnumProp *NullableStringEnum `json:"optionalNullableRefEnumProp,omitempty"`
}
