package handlers

type NumericTypesNumberNumberAnySimpleRequest struct {
	PathParam1 float32
	PathParam2 float32
	RequiredQuery1 float32
	RequiredQuery2 float32
	OptionalQuery1 float32
	OptionalQuery2 float32
}

type NumericTypesNumberController struct {
	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	//
	// Request type: NumericTypesNumberNumberAnySimpleRequest,
	//
	// Response type: none
	NumberAnySimple httpHandlerFactory
}

type NumericTypesNumberControllerBuilder struct {
	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	//
	// Request type: NumericTypesNumberNumberAnySimpleRequest,
	//
	// Response type: none
	HandleNumberAnySimple actionBuilderVoidResult[*NumericTypesNumberControllerBuilder, *NumericTypesNumberNumberAnySimpleRequest]
}

func (c *NumericTypesNumberControllerBuilder) Finalize() *NumericTypesNumberController {
	// TODO: panic if any handler is null
	return &NumericTypesNumberController{
		NumberAnySimple: c.HandleNumberAnySimple.httpHandlerFactory,
	}
}

func BuildNumericTypesNumberController() *NumericTypesNumberControllerBuilder {
	controllerBuilder := &NumericTypesNumberControllerBuilder{}

	// GET /numeric-types/number/any/{pathParam1}/{pathParam2}
	controllerBuilder.HandleNumberAnySimple.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumberAnySimple.defaultStatusCode = 204
	controllerBuilder.HandleNumberAnySimple.voidResult = true
	controllerBuilder.HandleNumberAnySimple.paramsParser = newNumericTypesNumberNumberAnySimpleParamsParser()

	return controllerBuilder
}

func MountNumericTypesNumberRoutes(controller *NumericTypesNumberController, r httpRouter) {
	r.HandleRoute("GET", "/numeric-types/number/any/{pathParam1}/{pathParam2}", controller.NumberAnySimple(r))
}
