// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// feedgen client
//
// Command:
// $ goa gen github.com/danlock/go-rss-gen/design

package feedgen

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "feedgen" service client.
type Client struct {
	MangaEndpoint goa.Endpoint
}

// NewClient initializes a "feedgen" service client given the endpoints.
func NewClient(manga goa.Endpoint) *Client {
	return &Client{
		MangaEndpoint: manga,
	}
}

// Manga calls the "manga" endpoint of the "feedgen" service.
func (c *Client) Manga(ctx context.Context, p *MangaPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.MangaEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
