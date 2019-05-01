// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListServicesHandlerFunc turns a function with the right signature into a list services handler
type ListServicesHandlerFunc func(ListServicesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListServicesHandlerFunc) Handle(params ListServicesParams) middleware.Responder {
	return fn(params)
}

// ListServicesHandler interface for that can handle valid list services params
type ListServicesHandler interface {
	Handle(ListServicesParams) middleware.Responder
}

// NewListServices creates a new http.Handler for the list services operation
func NewListServices(ctx *middleware.Context, handler ListServicesHandler) *ListServices {
	return &ListServices{Context: ctx, Handler: handler}
}

/*ListServices swagger:route GET /astrolabe listServices

List available services

This returns the list of services that this Astrolabe server supports


*/
type ListServices struct {
	Context *middleware.Context
	Handler ListServicesHandler
}

func (o *ListServices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListServicesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
