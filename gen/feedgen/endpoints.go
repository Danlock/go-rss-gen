// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// feedgen endpoints
//
// Command:
// $ goa gen github.com/danlock/go-rss-gen/design

package feedgen

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "feedgen" service endpoints.
type Endpoints struct {
	Manga goa.Endpoint
}

// NewEndpoints wraps the methods of the "feedgen" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Manga: NewMangaEndpoint(s),
	}
}

// Use applies the given middleware to all the "feedgen" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Manga = m(e.Manga)
}

// NewMangaEndpoint returns an endpoint function that calls the method "manga"
// of service "feedgen".
func NewMangaEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MangaPayload)
		return s.Manga(ctx, p)
	}
}