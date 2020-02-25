// Code generated by goa v3.0.10, DO NOT EDIT.
//
// modules client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package modules

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "modules" service client.
type Client struct {
	MetaEndpoint goa.Endpoint
}

// NewClient initializes a "modules" service client given the endpoints.
func NewClient(meta goa.Endpoint) *Client {
	return &Client{
		MetaEndpoint: meta,
	}
}

// Meta calls the "meta" endpoint of the "modules" service.
func (c *Client) Meta(ctx context.Context) (res *MetaResult, err error) {
	var ires interface{}
	ires, err = c.MetaEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*MetaResult), nil
}
