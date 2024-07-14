package handlers

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type paramsParserPetsCreatePet struct {
	bindPayload requestParamBinder[*http.Request, models.Pet]
}

func (p *paramsParserPetsCreatePet) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsCreatePetRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsCreatePetRequest{}
	// body params
	p.bindPayload(&bindingCtx, optionalVal[*http.Request]{value: req, assigned: true}, &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsCreatePet() *paramsParserPetsCreatePet {
	return &paramsParserPetsCreatePet{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, models.Pet]{
			field: "payload",
			location: "body",
			parseValue: parseJsonPayload[models.Pet],
			validateValue: newCompositeValidator[*http.Request, models.Pet](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserPetsGetPetById struct {
	bindPetId requestParamBinder[string, int64]
}

func (p *paramsParserPetsGetPetById) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsGetPetByIdRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsGetPetByIdRequest{}
	// path params
	p.bindPetId(&bindingCtx, readPathValue("petId", router, req), &reqParams.PetId)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsGetPetById() *paramsParserPetsGetPetById {
	return &paramsParserPetsGetPetById{
		bindPetId: newRequestParamBinder(binderParams[string, int64]{
			field: "petId",
			location: "path",
			parseValue: knownParsers.int64_in_path,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserPetsListPets struct {
	bindLimit requestParamBinder[[]string, int64]
	bindOffset requestParamBinder[[]string, int64]
}

func (p *paramsParserPetsListPets) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsListPetsRequest{}
	// query params
	query := req.URL.Query()
	p.bindLimit(&bindingCtx, readQueryValue("limit", query), &reqParams.Limit)
	p.bindOffset(&bindingCtx, readQueryValue("offset", query), &reqParams.Offset)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsListPets() *paramsParserPetsListPets {
	return &paramsParserPetsListPets{
		bindLimit: newRequestParamBinder(binderParams[[]string, int64]{
			field: "limit",
			location: "query",
			parseValue: knownParsers.int64_in_query,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int64](1, false, true),
				newMinMaxValueValidator[[]string, int64](100, false, false),
			),
		}),
		bindOffset: newRequestParamBinder(binderParams[[]string, int64]{
			field: "offset",
			location: "query",
			parseValue: knownParsers.int64_in_query,
			validateValue: newCompositeValidator[[]string, int64](
				newMinMaxValueValidator[[]string, int64](1, false, true),
			),
		}),
	}
}
