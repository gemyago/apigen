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
type BasicStringEnum string

// List of BasicStringEnum values.
const (
	BasicStringEnumVALUE1 BasicStringEnum = "VALUE1"
	BasicStringEnumVALUE2 BasicStringEnum = "VALUE2"
	BasicStringEnumVALUE3 BasicStringEnum = "VALUE3"
)

func(v BasicStringEnum) IsVALUE1() bool {
  return v == BasicStringEnumVALUE1
}

func(v BasicStringEnum) IsVALUE2() bool {
  return v == BasicStringEnumVALUE2
}

func(v BasicStringEnum) IsVALUE3() bool {
  return v == BasicStringEnumVALUE3
}

func(v BasicStringEnum) String() string {
	return string(v)
}

type assignableBasicStringEnum interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

func AsBasicStringEnum(v assignableBasicStringEnum) (BasicStringEnum) {
	return BasicStringEnum(v.String())
}

func ParseBasicStringEnum(str string, target *BasicStringEnum) error {
	switch str {
	case "VALUE1":
		*target = BasicStringEnumVALUE1
	case "VALUE2":
		*target = BasicStringEnumVALUE2
	case "VALUE3":
		*target = BasicStringEnumVALUE3
	default:
		return fmt.Errorf("unexpected BasicStringEnum value: %s", str)
	}
	return nil
}

func (v *BasicStringEnum) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return ParseBasicStringEnum(str, v)
}

// All allowed values of BasicStringEnum enum.
var AllowableBasicStringEnumValues = []BasicStringEnum{
	BasicStringEnumVALUE1,
	BasicStringEnumVALUE2,
	BasicStringEnumVALUE3,
}

