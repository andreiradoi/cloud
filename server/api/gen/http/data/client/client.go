// Code generated by goa v3.1.2, DO NOT EDIT.
//
// data client HTTP transport
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the data service endpoint HTTP clients.
type Client struct {
	// DeviceSummary Doer is the HTTP client used to make requests to the device
	// summary endpoint.
	DeviceSummaryDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the data service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		DeviceSummaryDoer:   doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// DeviceSummary returns an endpoint that makes HTTP requests to the data
// service device summary server.
func (c *Client) DeviceSummary() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeviceSummaryRequest(c.encoder)
		decodeResponse = DecodeDeviceSummaryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeviceSummaryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeviceSummaryDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("data", "device summary", err)
		}
		return decodeResponse(resp)
	}
}
