// Code generated by apigen DO NOT EDIT.

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/v1routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() Ping200Response

// PingPingRequest represents params for ping operation
//
// Request: GET /ping.
type PingPingRequest struct {
	// Message is parsed from request query and declared as message.
	Message string
}

type pingControllerBuilder struct {
	// GET /ping
	//
	// Request type: PingPingRequest,
	//
	// Response type: Ping200Response
	Ping genericHandlerBuilder[
		*PingPingRequest,
		*Ping200Response,
		func(context.Context, *PingPingRequest) (*Ping200Response, error),
		func(http.ResponseWriter, *http.Request, *PingPingRequest) (*Ping200Response, error),
	]
}

func newPingControllerBuilder(app *HTTPApp) *pingControllerBuilder {
	return &pingControllerBuilder{
		// GET /ping
		Ping: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*PingPingRequest,
				*Ping200Response,
			](),
			newHTTPHandlerAdapter[
				*PingPingRequest,
				*Ping200Response,
			](),
			makeActionBuilderParams[
				*PingPingRequest,
				*Ping200Response,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserPingPing(app),
			},
		),
	}
}

type PingController interface {
	// GET /ping
	//
	// Request type: PingPingRequest,
	//
	// Response type: Ping200Response
	Ping(HandlerBuilder[
		*PingPingRequest,
		*Ping200Response,
	]) http.Handler
}

func RegisterPingRoutes(controller PingController, app *HTTPApp) {
	builder := newPingControllerBuilder(app)
	app.router.HandleRoute("GET", "/ping", controller.Ping(builder.Ping))
}