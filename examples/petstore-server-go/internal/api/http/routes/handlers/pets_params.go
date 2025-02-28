// Code generated by apigen DO NOT EDIT.

package handlers

import (
	"net/http"
	"time"

	. "github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/routes/models"
	. "github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/routes/internal"
)

// Below is to workaround unused imports if that happens.
var _ = BindingContext{}
var _ = http.MethodGet
var _ = time.Time{}
type _ func() Error

type paramsParserPetsCreatePet struct {
	bindPayload requestParamBinder[*http.Request, *Pet]
}

func (p *paramsParserPetsCreatePet) parse(router httpRouter, req *http.Request) (*CreatePetParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &CreatePetParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsCreatePet(rootHandler *RootHandler) paramsParser[*CreatePetParams] {
	return &paramsParserPetsCreatePet{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *Pet]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*Pet],
			),
			validateValue: NewPetValidator(),
		}),
	}
}

type paramsParserPetsGetPetByID struct {
	bindPetID requestParamBinder[string, int64]
}

func (p *paramsParserPetsGetPetByID) parse(router httpRouter, req *http.Request) (*GetPetByIDParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &GetPetByIDParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindPetID(pathParamsCtx.Fork("petId"), readPathValue("petId", router, req), &reqParams.PetID)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsGetPetByID(rootHandler *RootHandler) paramsParser[*GetPetByIDParams] {
	return &paramsParserPetsGetPetByID{
		bindPetID: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
			),
		}),
	}
}

type paramsParserPetsListPets struct {
	bindLimit requestParamBinder[[]string, int64]
	bindOffset requestParamBinder[[]string, int64]
}

func (p *paramsParserPetsListPets) parse(router httpRouter, req *http.Request) (*ListPetsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ListPetsParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindLimit(queryParamsCtx.Fork("limit"), readQueryValue("limit", query), &reqParams.Limit)
	p.bindOffset(queryParamsCtx.Fork("offset"), readQueryValue("offset", query), &reqParams.Offset)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserPetsListPets(rootHandler *RootHandler) paramsParser[*ListPetsParams] {
	return &paramsParserPetsListPets{
		bindLimit: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](1, false, true),
				NewMinMaxValueValidator[int64](100, false, false),
			),
		}),
		bindOffset: newRequestParamBinder(binderParams[[]string, int64]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](1, false, true),
			),
		}),
	}
}

type petsControllerBuilder struct {
	// POST /pets
	//
	// Request type: CreatePetParams,
	//
	// Response type: none
	CreatePet genericHandlerBuilder[
		*CreatePetParams,
		void,
		handlerActionFuncNoResponse[*CreatePetParams, void],
		httpHandlerActionFuncNoResponse[*CreatePetParams, void],
	]

	// GET /pets/{petId}
	//
	// Request type: GetPetByIDParams,
	//
	// Response type: PetResponse
	GetPetByID genericHandlerBuilder[
		*GetPetByIDParams,
		*PetResponse,
		handlerActionFunc[*GetPetByIDParams, *PetResponse],
		httpHandlerActionFunc[*GetPetByIDParams, *PetResponse],
	]

	// GET /pets
	//
	// Request type: ListPetsParams,
	//
	// Response type: PetsResponse
	ListPets genericHandlerBuilder[
		*ListPetsParams,
		*PetsResponse,
		handlerActionFunc[*ListPetsParams, *PetsResponse],
		httpHandlerActionFunc[*ListPetsParams, *PetsResponse],
	]
}

func newPetsControllerBuilder(app *RootHandler) *petsControllerBuilder {
	return &petsControllerBuilder{
		// POST /pets
		CreatePet: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*CreatePetParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*CreatePetParams,
				void,
			](),
			makeActionBuilderParams[
				*CreatePetParams,
				void,
			]{
				defaultStatus: 201,
				voidResult:    true,
				paramsParser:  newParamsParserPetsCreatePet(app),
			},
		),

		// GET /pets/{petId}
		GetPetByID: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*GetPetByIDParams,
				*PetResponse,
			](),
			newHTTPHandlerAdapter[
				*GetPetByIDParams,
				*PetResponse,
			](),
			makeActionBuilderParams[
				*GetPetByIDParams,
				*PetResponse,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserPetsGetPetByID(app),
			},
		),

		// GET /pets
		ListPets: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*ListPetsParams,
				*PetsResponse,
			](),
			newHTTPHandlerAdapter[
				*ListPetsParams,
				*PetsResponse,
			](),
			makeActionBuilderParams[
				*ListPetsParams,
				*PetsResponse,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserPetsListPets(app),
			},
		),
	}
}
