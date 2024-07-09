package handlers

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type PetsCreatePetParamsParser struct {
	bindPayload requestParamBinder[*http.Request, models.Pet]
}

func (p *PetsCreatePetParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsCreatePetRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsCreatePetRequest{}
	
	p.bindPayload(&bindingCtx, optionalVal[*http.Request]{value: req, assigned: true}, &reqParams.Payload)
	return reqParams, bindingCtx.Error()
}

func newPetsCreatePetParamsParser() *PetsCreatePetParamsParser {
	return &PetsCreatePetParamsParser{
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
type PetsGetPetByIdParamsParser struct {
	bindPetId requestParamBinder[string, int64]
}

func (p *PetsGetPetByIdParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsGetPetByIdRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsGetPetByIdRequest{}
	
	p.bindPetId(&bindingCtx, readPathValue("petId", router, req), &reqParams.PetId)
	return reqParams, bindingCtx.Error()
}

func newPetsGetPetByIdParamsParser() *PetsGetPetByIdParamsParser {
	return &PetsGetPetByIdParamsParser{
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
type PetsListPetsParamsParser struct {
	bindLimit requestParamBinder[[]string, int64]
	bindOffset requestParamBinder[[]string, int64]
}

func (p *PetsListPetsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsListPetsRequest{}
	query := req.URL.Query()
	p.bindLimit(&bindingCtx, readQueryValue("limit", query), &reqParams.Limit)
	p.bindOffset(&bindingCtx, readQueryValue("offset", query), &reqParams.Offset)
	return reqParams, bindingCtx.Error()
}

func newPetsListPetsParamsParser() *PetsListPetsParamsParser {
	return &PetsListPetsParamsParser{
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
