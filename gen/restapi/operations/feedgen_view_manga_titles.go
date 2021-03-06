// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// FeedgenViewMangaTitlesHandlerFunc turns a function with the right signature into a feedgen view manga titles handler
type FeedgenViewMangaTitlesHandlerFunc func(FeedgenViewMangaTitlesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FeedgenViewMangaTitlesHandlerFunc) Handle(params FeedgenViewMangaTitlesParams) middleware.Responder {
	return fn(params)
}

// FeedgenViewMangaTitlesHandler interface for that can handle valid feedgen view manga titles params
type FeedgenViewMangaTitlesHandler interface {
	Handle(FeedgenViewMangaTitlesParams) middleware.Responder
}

// NewFeedgenViewMangaTitles creates a new http.Handler for the feedgen view manga titles operation
func NewFeedgenViewMangaTitles(ctx *middleware.Context, handler FeedgenViewMangaTitlesHandler) *FeedgenViewMangaTitles {
	return &FeedgenViewMangaTitles{Context: ctx, Handler: handler}
}

/*FeedgenViewMangaTitles swagger:route GET /api/feed/manga/{hash}/titles/ feedgenViewMangaTitles

Get manga titles inside feed

*/
type FeedgenViewMangaTitles struct {
	Context *middleware.Context
	Handler FeedgenViewMangaTitlesHandler
}

func (o *FeedgenViewMangaTitles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFeedgenViewMangaTitlesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FeedgenViewMangaTitlesOKBody feedgen view manga titles o k body
// swagger:model FeedgenViewMangaTitlesOKBody
type FeedgenViewMangaTitlesOKBody struct {

	// titles
	Titles []string `json:"titles"`
}

// Validate validates this feedgen view manga titles o k body
func (o *FeedgenViewMangaTitlesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FeedgenViewMangaTitlesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FeedgenViewMangaTitlesOKBody) UnmarshalBinary(b []byte) error {
	var res FeedgenViewMangaTitlesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
