// Code generated by go-swagger; DO NOT EDIT.

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateBindHandlerFunc turns a function with the right signature into a create bind handler
type CreateBindHandlerFunc func(CreateBindParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBindHandlerFunc) Handle(params CreateBindParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateBindHandler interface for that can handle valid create bind params
type CreateBindHandler interface {
	Handle(CreateBindParams, interface{}) middleware.Responder
}

// NewCreateBind creates a new http.Handler for the create bind operation
func NewCreateBind(ctx *middleware.Context, handler CreateBindHandler) *CreateBind {
	return &CreateBind{Context: ctx, Handler: handler}
}

/*CreateBind swagger:route POST /services/haproxy/configuration/binds Bind createBind

Add a new bind

Adds a new bind in the specified frontend in the configuration file.

*/
type CreateBind struct {
	Context *middleware.Context
	Handler CreateBindHandler
}

func (o *CreateBind) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateBindParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}