// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// feedgen service
//
// Command:
// $ goa gen github.com/danlock/go-rss-gen/design

package feedgen

import (
	"context"
)

// Service is the feedgen service interface.
type Service interface {
	// Manga implements manga.
	Manga(context.Context, *MangaPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "feedgen"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"manga"}

// MangaPayload is the payload type of the feedgen service manga method.
type MangaPayload struct {
	// List of manga titles to subscribe to
	Titles []string
}