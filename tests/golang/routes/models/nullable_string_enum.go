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
type NullableStringEnum string

// List of NullableStringEnum values.
const (
	NullableStringEnumVALUE1 NullableStringEnum = "VALUE1"
	NullableStringEnumVALUE2 NullableStringEnum = "VALUE2"
	NullableStringEnumVALUE3 NullableStringEnum = "VALUE3"
)

func(v NullableStringEnum) IsVALUE1() bool {
  return v == NullableStringEnumVALUE1
}

func(v NullableStringEnum) IsVALUE2() bool {
  return v == NullableStringEnumVALUE2
}

func(v NullableStringEnum) IsVALUE3() bool {
  return v == NullableStringEnumVALUE3
}

func(v NullableStringEnum) String() string {
	return string(v)
}

type assignableNullableStringEnum interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsNullableStringEnum(v assignableNullableStringEnum) (NullableStringEnum) {
	return NullableStringEnum(v.String())
}

func ParseNullableStringEnum(str string, target *NullableStringEnum) error {
	switch str {
	case "VALUE1":
		*target = NullableStringEnumVALUE1
	case "VALUE2":
		*target = NullableStringEnumVALUE2
	case "VALUE3":
		*target = NullableStringEnumVALUE3
	default:
		return fmt.Errorf("unexpected NullableStringEnum value: %s", str)
	}
	return nil
}

func (v *NullableStringEnum) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseNullableStringEnum(str, v)
}

// All allowed values of NullableStringEnum enum.
var AllowableNullableStringEnumValues = []NullableStringEnum{
	NullableStringEnumVALUE1,
	NullableStringEnumVALUE2,
	NullableStringEnumVALUE3,
}

