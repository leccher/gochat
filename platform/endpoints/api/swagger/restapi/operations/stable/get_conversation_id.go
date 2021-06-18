// Code generated by go-swagger; DO NOT EDIT.

package stable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gochat/platform/endpoints/api/swagger/models"
)

// GetConversationIDHandlerFunc turns a function with the right signature into a get conversation ID handler
type GetConversationIDHandlerFunc func(GetConversationIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConversationIDHandlerFunc) Handle(params GetConversationIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetConversationIDHandler interface for that can handle valid get conversation ID params
type GetConversationIDHandler interface {
	Handle(GetConversationIDParams, *models.Principal) middleware.Responder
}

// NewGetConversationID creates a new http.Handler for the get conversation ID operation
func NewGetConversationID(ctx *middleware.Context, handler GetConversationIDHandler) *GetConversationID {
	return &GetConversationID{Context: ctx, Handler: handler}
}

/* GetConversationID swagger:route GET /conversation/{id} stable getConversationId

Fetches an existing conversation.

*/
type GetConversationID struct {
	Context *middleware.Context
	Handler GetConversationIDHandler
}

func (o *GetConversationID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConversationIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
