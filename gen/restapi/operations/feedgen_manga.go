// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FeedgenMangaHandlerFunc turns a function with the right signature into a feedgen manga handler
type FeedgenMangaHandlerFunc func(FeedgenMangaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FeedgenMangaHandlerFunc) Handle(params FeedgenMangaParams) middleware.Responder {
	return fn(params)
}

// FeedgenMangaHandler interface for that can handle valid feedgen manga params
type FeedgenMangaHandler interface {
	Handle(FeedgenMangaParams) middleware.Responder
}

// NewFeedgenManga creates a new http.Handler for the feedgen manga operation
func NewFeedgenManga(ctx *middleware.Context, handler FeedgenMangaHandler) *FeedgenManga {
	return &FeedgenManga{Context: ctx, Handler: handler}
}

/*FeedgenManga swagger:route POST /api/feed/manga feedgenManga

Create feed from manga titles

Creates a URL containing the current feed for the requested manga titles

*/
type FeedgenManga struct {
	Context *middleware.Context
	Handler FeedgenMangaHandler
}

func (o *FeedgenManga) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFeedgenMangaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}