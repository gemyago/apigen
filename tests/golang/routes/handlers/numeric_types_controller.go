package handlers

type NumericTypesNumberAnySimpleRequest struct {
	PathParam1 float32
	PathParam2 float32
	RequiredQuery1 float32
	RequiredQuery2 float32
	OptionalQuery1 float32
	OptionalQuery2 float32
}

type NumericTypesController struct {
	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	//
	// Request type: NumericTypesNumberAnySimpleRequest,
	//
	// Response type: none
	NumberAnySimple httpHandlerFactory
}

type NumericTypesControllerBuilder struct {
	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	//
	// Request type: NumericTypesNumberAnySimpleRequest,
	//
	// Response type: none
	HandleNumberAnySimple actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumberAnySimpleRequest]
}

func (c *NumericTypesControllerBuilder) Finalize() *NumericTypesController {
	// TODO: panic if any handler is null
	return &NumericTypesController{
		NumberAnySimple: c.HandleNumberAnySimple.httpHandlerFactory,
	}
}

func BuildNumericTypesController() *NumericTypesControllerBuilder {
	controllerBuilder := &NumericTypesControllerBuilder{}

	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	controllerBuilder.HandleNumberAnySimple.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumberAnySimple.defaultStatusCode = 204
	controllerBuilder.HandleNumberAnySimple.voidResult = true
	controllerBuilder.HandleNumberAnySimple.paramsParser = newNumericTypesNumberAnySimpleParamsParser()

	return controllerBuilder
}

func MountNumericTypesRoutes(controller *NumericTypesController, r httpRouter) {
	r.HandleRoute("GET", "/numeric-types/number/any/{pathParam1}/{pathParam2}", controller.NumberAnySimple(r))
}
