// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// feedgen HTTP client types
//
// Command:
// $ goa gen github.com/danlock/go-rss-gen/design

package client

import (
	feedgen "github.com/danlock/go-rss-gen/gen/feedgen"
)

// MangaRequestBody is the type of the "feedgen" service "manga" endpoint HTTP
// request body.
type MangaRequestBody struct {
	// List of manga titles to subscribe to
	Titles []string `form:"titles" json:"titles" xml:"titles"`
}

// NewMangaRequestBody builds the HTTP request body from the payload of the
// "manga" endpoint of the "feedgen" service.
func NewMangaRequestBody(p *feedgen.MangaPayload) *MangaRequestBody {
	body := &MangaRequestBody{}
	if p.Titles != nil {
		body.Titles = make([]string, len(p.Titles))
		for i, val := range p.Titles {
			body.Titles[i] = val
		}
	}
	return body
}
