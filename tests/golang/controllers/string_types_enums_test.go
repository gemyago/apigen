package controllers

import (
	"encoding/json"
	"math/rand/v2"
	"slices"
	"strconv"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type standardBasicEnumProperties interface {
	IsVALUE1() bool
	IsVALUE2() bool
	IsVALUE3() bool
	String() string
}

type basicStringEnumTestCaseParams[
	TEnum standardBasicEnumProperties,
] struct {
	allowableValues []TEnum
	asFn            func(TEnum) TEnum
	parserFn        func(string, *TEnum) error
}

func newBasicStringEnumTestCase[TEnum standardBasicEnumProperties](
	tc basicStringEnumTestCaseParams[TEnum],
) func(t *testing.T) {
	fake := faker.New()
	return func(t *testing.T) {
		t.Run("IsMethods", func(t *testing.T) {
			t.Run("IsVALUE1", func(t *testing.T) {
				v := tc.allowableValues[0]
				assert.True(t, v.IsVALUE1())
				assert.False(t, v.IsVALUE2())
				assert.False(t, v.IsVALUE3())
			})
			t.Run("IsVALUE2", func(t *testing.T) {
				v := tc.allowableValues[1]
				assert.False(t, v.IsVALUE1())
				assert.True(t, v.IsVALUE2())
				assert.False(t, v.IsVALUE3())
			})
			t.Run("IsVALUE3", func(t *testing.T) {
				v := tc.allowableValues[2]
				assert.False(t, v.IsVALUE1())
				assert.False(t, v.IsVALUE2())
				assert.True(t, v.IsVALUE3())
			})
		})
		t.Run("AsEnum", func(t *testing.T) {
			t.Run("assign compatible value", func(t *testing.T) {
				val1 := lo.Shuffle(slices.Clone(tc.allowableValues))[0]
				val2 := tc.asFn(val1)
				assert.Equal(t, val1, val2)
			})
		})
		t.Run("ParseEnum", func(t *testing.T) {
			t.Run("valid value", func(t *testing.T) {
				for _, v := range tc.allowableValues {
					var val TEnum
					err := tc.parserFn(v.String(), &val)
					require.NoError(t, err)
					assert.Equal(t, v, val)
				}
			})
			t.Run("invalid value", func(t *testing.T) {
				var val TEnum
				err := tc.parserFn("invalid", &val)
				assert.Error(t, err)
			})
		})
		t.Run("UnmarshalJSON", func(t *testing.T) {
			t.Run("valid value", func(t *testing.T) {
				for _, v := range tc.allowableValues {
					data := []byte(`"` + v.String() + `"`)
					var val TEnum
					err := json.Unmarshal(data, &val)
					require.NoError(t, err)
					assert.Equal(t, v, val)
				}
			})
			t.Run("invalid value", func(t *testing.T) {
				data := []byte(`"` + fake.UUID().V4() + `"`)
				var val TEnum
				err := json.Unmarshal(data, &val)
				require.Error(t, err)
			})
			t.Run("invalid json", func(t *testing.T) {
				data := []byte(`{"value": ` + strconv.Itoa(rand.Int()) + `}`)
				var val TEnum
				err := json.Unmarshal(data, &val)
				require.Error(t, err)
			})
		})
	}
}

func TestStringTypesEnums(t *testing.T) {
	t.Run("BasicStringEnum", newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.BasicStringEnum]{
		allowableValues: models.AllowableBasicStringEnumValues,
		asFn: func(val models.BasicStringEnum) models.BasicStringEnum {
			return models.AsBasicStringEnum(val)
		},
		parserFn: models.ParseBasicStringEnum,
	}))
	t.Run("StringTypesEnumsRequestInlineEnumProp",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesEnumsRequestInlineEnumProp]{
			allowableValues: models.AllowableStringTypesEnumsRequestInlineEnumPropValues,
			asFn: func(val models.StringTypesEnumsRequestInlineEnumProp) models.StringTypesEnumsRequestInlineEnumProp {
				return models.AsStringTypesEnumsRequestInlineEnumProp(val)
			},
			parserFn: models.ParseStringTypesEnumsRequestInlineEnumProp,
		}))
	t.Run("StringTypesEnumsRequestOptionalInlineEnumProp",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesEnumsRequestOptionalInlineEnumProp]{
			allowableValues: models.AllowableStringTypesEnumsRequestOptionalInlineEnumPropValues,
			asFn: func(
				val models.StringTypesEnumsRequestOptionalInlineEnumProp,
			) models.StringTypesEnumsRequestOptionalInlineEnumProp {
				return models.AsStringTypesEnumsRequestOptionalInlineEnumProp(val)
			},
			parserFn: models.ParseStringTypesEnumsRequestOptionalInlineEnumProp,
		}))
	t.Run("StringTypesEnumsRequestNullableInlineEnumProp",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesEnumsRequestNullableInlineEnumProp]{
			allowableValues: models.AllowableStringTypesEnumsRequestNullableInlineEnumPropValues,
			asFn: func(
				val models.StringTypesEnumsRequestNullableInlineEnumProp,
			) models.StringTypesEnumsRequestNullableInlineEnumProp {
				return models.AsStringTypesEnumsRequestNullableInlineEnumProp(val)
			},
			parserFn: models.ParseStringTypesEnumsRequestNullableInlineEnumProp,
		}))
	t.Run("StringTypesEnumsRequestOptionalNullableInlineEnumProp",
		newBasicStringEnumTestCase(
			basicStringEnumTestCaseParams[models.StringTypesEnumsRequestOptionalNullableInlineEnumProp]{
				allowableValues: models.AllowableStringTypesEnumsRequestOptionalNullableInlineEnumPropValues,
				asFn: func(
					val models.StringTypesEnumsRequestOptionalNullableInlineEnumProp,
				) models.StringTypesEnumsRequestOptionalNullableInlineEnumProp {
					return models.AsStringTypesEnumsRequestOptionalNullableInlineEnumProp(val)
				},
				parserFn: models.ParseStringTypesEnumsRequestOptionalNullableInlineEnumProp,
			}))
	t.Run("NullableStringEnum",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.NullableStringEnum]{
			allowableValues: models.AllowableNullableStringEnumValues,
			asFn: func(val models.NullableStringEnum) models.NullableStringEnum {
				return models.AsNullableStringEnum(val)
			},
			parserFn: models.ParseNullableStringEnum,
		}))
	t.Run("models.StringTypesStringTypesEnumsInlineEnumParam",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsInlineEnumParam]{
			allowableValues: models.AllowableStringTypesStringTypesEnumsInlineEnumParamValues,
			asFn: func(
				val models.StringTypesStringTypesEnumsInlineEnumParam,
			) models.StringTypesStringTypesEnumsInlineEnumParam {
				return models.AsStringTypesStringTypesEnumsInlineEnumParam(val)
			},
			parserFn: models.ParseStringTypesStringTypesEnumsInlineEnumParam,
		}))
	t.Run("StringTypesStringTypesEnumsNullableInlineEnumParam",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsNullableInlineEnumParam]{
			allowableValues: models.AllowableStringTypesStringTypesEnumsNullableInlineEnumParamValues,
			asFn: func(
				val models.StringTypesStringTypesEnumsNullableInlineEnumParam,
			) models.StringTypesStringTypesEnumsNullableInlineEnumParam {
				return models.AsStringTypesStringTypesEnumsNullableInlineEnumParam(val)
			},
			parserFn: models.ParseStringTypesStringTypesEnumsNullableInlineEnumParam,
		}))
	t.Run("StringTypesStringTypesEnumsInlineEnumParamInQuery",
		newBasicStringEnumTestCase(basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsInlineEnumParamInQuery]{
			allowableValues: models.AllowableStringTypesStringTypesEnumsInlineEnumParamInQueryValues,
			asFn: func(
				val models.StringTypesStringTypesEnumsInlineEnumParamInQuery,
			) models.StringTypesStringTypesEnumsInlineEnumParamInQuery {
				return models.AsStringTypesStringTypesEnumsInlineEnumParamInQuery(val)
			},
			parserFn: models.ParseStringTypesStringTypesEnumsInlineEnumParamInQuery,
		}))
	t.Run("StringTypesStringTypesEnumsNullableInlineEnumParamInQuery",
		newBasicStringEnumTestCase(
			basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsNullableInlineEnumParamInQuery]{
				allowableValues: models.AllowableStringTypesStringTypesEnumsNullableInlineEnumParamInQueryValues,
				asFn: func(
					val models.StringTypesStringTypesEnumsNullableInlineEnumParamInQuery,
				) models.StringTypesStringTypesEnumsNullableInlineEnumParamInQuery {
					return models.AsStringTypesStringTypesEnumsNullableInlineEnumParamInQuery(val)
				},
				parserFn: models.ParseStringTypesStringTypesEnumsNullableInlineEnumParamInQuery,
			}))
	t.Run("StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery",
		newBasicStringEnumTestCase(
			basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery]{
				allowableValues: models.AllowableStringTypesStringTypesEnumsOptionalInlineEnumParamInQueryValues,
				asFn: func(
					val models.StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery,
				) models.StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery {
					return models.AsStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery(val)
				},
				parserFn: models.ParseStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery,
			}))
	t.Run("StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery",
		newBasicStringEnumTestCase(
			basicStringEnumTestCaseParams[models.StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery]{
				allowableValues: models.AllowableStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryValues,
				asFn: func(
					val models.StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery,
				) models.StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery {
					return models.AsStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery(val)
				},
				parserFn: models.ParseStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery,
			}))
}
