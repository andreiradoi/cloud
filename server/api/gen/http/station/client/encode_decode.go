// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	station "github.com/fieldkit/cloud/server/api/gen/station"
	stationviews "github.com/fieldkit/cloud/server/api/gen/station/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "station" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddStationPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the station add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "add", "*station.AddPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("station", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the station
// add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeAddResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AddResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "add", err)
			}
			p := NewAddStationFullOK(&body)
			view := "default"
			vres := &stationviews.StationFull{Projected: p, View: view}
			if err = stationviews.ValidateStationFull(vres); err != nil {
				return nil, goahttp.ErrValidationError("station", "add", err)
			}
			res := station.NewStationFull(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body AddBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "add", err)
			}
			return nil, NewAddBadRequest(body)
		case http.StatusForbidden:
			var (
				body AddForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "add", err)
			}
			return nil, NewAddForbidden(body)
		case http.StatusNotFound:
			var (
				body AddNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "add", err)
			}
			return nil, NewAddNotFound(body)
		case http.StatusUnauthorized:
			var (
				body AddUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "add", err)
			}
			return nil, NewAddUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "station" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int32
	)
	{
		p, ok := v.(*station.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("station", "get", "*station.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetStationPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the station get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "get", "*station.GetPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the station
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "get", err)
			}
			p := NewGetStationFullOK(&body)
			view := "default"
			vres := &stationviews.StationFull{Projected: p, View: view}
			if err = stationviews.ValidateStationFull(vres); err != nil {
				return nil, goahttp.ErrValidationError("station", "get", err)
			}
			res := station.NewStationFull(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "get", err)
			}
			return nil, NewGetBadRequest(body)
		case http.StatusForbidden:
			var (
				body GetForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "get", err)
			}
			return nil, NewGetForbidden(body)
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "get", err)
			}
			return nil, NewGetNotFound(body)
		case http.StatusUnauthorized:
			var (
				body GetUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "get", err)
			}
			return nil, NewGetUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "station" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int32
	)
	{
		p, ok := v.(*station.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("station", "update", "*station.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateStationPath(id)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the station
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "update", "*station.UpdatePayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("station", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the station
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "update", err)
			}
			p := NewUpdateStationFullOK(&body)
			view := "default"
			vres := &stationviews.StationFull{Projected: p, View: view}
			if err = stationviews.ValidateStationFull(vres); err != nil {
				return nil, goahttp.ErrValidationError("station", "update", err)
			}
			res := station.NewStationFull(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "update", err)
			}
			return nil, NewUpdateBadRequest(body)
		case http.StatusForbidden:
			var (
				body UpdateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "update", err)
			}
			return nil, NewUpdateForbidden(body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "update", err)
			}
			return nil, NewUpdateNotFound(body)
		case http.StatusUnauthorized:
			var (
				body UpdateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "update", err)
			}
			return nil, NewUpdateUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildListMineRequest instantiates a HTTP request object with method and path
// set to call the "station" service "list mine" endpoint
func (c *Client) BuildListMineRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMineStationPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "list mine", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListMineRequest returns an encoder for requests sent to the station
// list mine server.
func EncodeListMineRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.ListMinePayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "list mine", "*station.ListMinePayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListMineResponse returns a decoder for responses returned by the
// station list mine endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListMineResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListMineResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListMineResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list mine", err)
			}
			p := NewListMineStationsFullOK(&body)
			view := "default"
			vres := &stationviews.StationsFull{Projected: p, View: view}
			if err = stationviews.ValidateStationsFull(vres); err != nil {
				return nil, goahttp.ErrValidationError("station", "list mine", err)
			}
			res := station.NewStationsFull(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListMineBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list mine", err)
			}
			return nil, NewListMineBadRequest(body)
		case http.StatusForbidden:
			var (
				body ListMineForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list mine", err)
			}
			return nil, NewListMineForbidden(body)
		case http.StatusNotFound:
			var (
				body ListMineNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list mine", err)
			}
			return nil, NewListMineNotFound(body)
		case http.StatusUnauthorized:
			var (
				body ListMineUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list mine", err)
			}
			return nil, NewListMineUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "list mine", resp.StatusCode, string(body))
		}
	}
}

// BuildListProjectRequest instantiates a HTTP request object with method and
// path set to call the "station" service "list project" endpoint
func (c *Client) BuildListProjectRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int32
	)
	{
		p, ok := v.(*station.ListProjectPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("station", "list project", "*station.ListProjectPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProjectStationPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "list project", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListProjectRequest returns an encoder for requests sent to the station
// list project server.
func EncodeListProjectRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.ListProjectPayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "list project", "*station.ListProjectPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListProjectResponse returns a decoder for responses returned by the
// station list project endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeListProjectResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListProjectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListProjectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list project", err)
			}
			p := NewListProjectStationsFullOK(&body)
			view := "default"
			vres := &stationviews.StationsFull{Projected: p, View: view}
			if err = stationviews.ValidateStationsFull(vres); err != nil {
				return nil, goahttp.ErrValidationError("station", "list project", err)
			}
			res := station.NewStationsFull(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListProjectBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list project", err)
			}
			return nil, NewListProjectBadRequest(body)
		case http.StatusForbidden:
			var (
				body ListProjectForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list project", err)
			}
			return nil, NewListProjectForbidden(body)
		case http.StatusNotFound:
			var (
				body ListProjectNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list project", err)
			}
			return nil, NewListProjectNotFound(body)
		case http.StatusUnauthorized:
			var (
				body ListProjectUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "list project", err)
			}
			return nil, NewListProjectUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "list project", resp.StatusCode, string(body))
		}
	}
}

// BuildPhotoRequest instantiates a HTTP request object with method and path
// set to call the "station" service "photo" endpoint
func (c *Client) BuildPhotoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int32
	)
	{
		p, ok := v.(*station.PhotoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("station", "photo", "*station.PhotoPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PhotoStationPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("station", "photo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePhotoRequest returns an encoder for requests sent to the station photo
// server.
func EncodePhotoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*station.PhotoPayload)
		if !ok {
			return goahttp.ErrInvalidType("station", "photo", "*station.PhotoPayload", v)
		}
		values := req.URL.Query()
		values.Add("token", p.Auth)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodePhotoResponse returns a decoder for responses returned by the station
// photo endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePhotoResponse may return the following errors:
//	- "bad-request" (type station.BadRequest): http.StatusBadRequest
//	- "forbidden" (type station.Forbidden): http.StatusForbidden
//	- "not-found" (type station.NotFound): http.StatusNotFound
//	- "unauthorized" (type station.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodePhotoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				length      int64
				contentType string
				err         error
			)
			{
				lengthRaw := resp.Header.Get("Content-Length")
				if lengthRaw == "" {
					return nil, goahttp.ErrValidationError("station", "photo", goa.MissingFieldError("Content-Length", "header"))
				}
				v, err2 := strconv.ParseInt(lengthRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("length", lengthRaw, "integer"))
				}
				length = v
			}
			contentTypeRaw := resp.Header.Get("Content-Type")
			if contentTypeRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Type", "header"))
			}
			contentType = contentTypeRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("station", "photo", err)
			}
			res := NewPhotoResultOK(length, contentType)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PhotoBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "photo", err)
			}
			return nil, NewPhotoBadRequest(body)
		case http.StatusForbidden:
			var (
				body PhotoForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "photo", err)
			}
			return nil, NewPhotoForbidden(body)
		case http.StatusNotFound:
			var (
				body PhotoNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "photo", err)
			}
			return nil, NewPhotoNotFound(body)
		case http.StatusUnauthorized:
			var (
				body PhotoUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("station", "photo", err)
			}
			return nil, NewPhotoUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("station", "photo", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStationOwnerResponseBodyToStationviewsStationOwnerView builds a
// value of type *stationviews.StationOwnerView from a value of type
// *StationOwnerResponseBody.
func unmarshalStationOwnerResponseBodyToStationviewsStationOwnerView(v *StationOwnerResponseBody) *stationviews.StationOwnerView {
	res := &stationviews.StationOwnerView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalStationUploadResponseBodyToStationviewsStationUploadView builds a
// value of type *stationviews.StationUploadView from a value of type
// *StationUploadResponseBody.
func unmarshalStationUploadResponseBodyToStationviewsStationUploadView(v *StationUploadResponseBody) *stationviews.StationUploadView {
	res := &stationviews.StationUploadView{
		ID:       v.ID,
		Time:     v.Time,
		UploadID: v.UploadID,
		Size:     v.Size,
		URL:      v.URL,
		Type:     v.Type,
	}
	res.Blocks = make([]int64, len(v.Blocks))
	for i, val := range v.Blocks {
		res.Blocks[i] = val
	}

	return res
}

// unmarshalImageRefResponseBodyToStationviewsImageRefView builds a value of
// type *stationviews.ImageRefView from a value of type *ImageRefResponseBody.
func unmarshalImageRefResponseBodyToStationviewsImageRefView(v *ImageRefResponseBody) *stationviews.ImageRefView {
	res := &stationviews.ImageRefView{
		URL: v.URL,
	}

	return res
}

// unmarshalStationPhotosResponseBodyToStationviewsStationPhotosView builds a
// value of type *stationviews.StationPhotosView from a value of type
// *StationPhotosResponseBody.
func unmarshalStationPhotosResponseBodyToStationviewsStationPhotosView(v *StationPhotosResponseBody) *stationviews.StationPhotosView {
	res := &stationviews.StationPhotosView{
		Small: v.Small,
	}

	return res
}

// unmarshalStationModuleResponseBodyToStationviewsStationModuleView builds a
// value of type *stationviews.StationModuleView from a value of type
// *StationModuleResponseBody.
func unmarshalStationModuleResponseBodyToStationviewsStationModuleView(v *StationModuleResponseBody) *stationviews.StationModuleView {
	res := &stationviews.StationModuleView{
		ID:           v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         v.Name,
		Position:     v.Position,
		Flags:        v.Flags,
		Internal:     v.Internal,
	}
	res.Sensors = make([]*stationviews.StationSensorView, len(v.Sensors))
	for i, val := range v.Sensors {
		res.Sensors[i] = unmarshalStationSensorResponseBodyToStationviewsStationSensorView(val)
	}

	return res
}

// unmarshalStationSensorResponseBodyToStationviewsStationSensorView builds a
// value of type *stationviews.StationSensorView from a value of type
// *StationSensorResponseBody.
func unmarshalStationSensorResponseBodyToStationviewsStationSensorView(v *StationSensorResponseBody) *stationviews.StationSensorView {
	res := &stationviews.StationSensorView{
		Name:          v.Name,
		UnitOfMeasure: v.UnitOfMeasure,
	}
	if v.Reading != nil {
		res.Reading = unmarshalSensorReadingResponseBodyToStationviewsSensorReadingView(v.Reading)
	}

	return res
}

// unmarshalSensorReadingResponseBodyToStationviewsSensorReadingView builds a
// value of type *stationviews.SensorReadingView from a value of type
// *SensorReadingResponseBody.
func unmarshalSensorReadingResponseBodyToStationviewsSensorReadingView(v *SensorReadingResponseBody) *stationviews.SensorReadingView {
	if v == nil {
		return nil
	}
	res := &stationviews.SensorReadingView{
		Last: v.Last,
		Time: v.Time,
	}

	return res
}

// unmarshalStationFullResponseBodyToStationviewsStationFullView builds a value
// of type *stationviews.StationFullView from a value of type
// *StationFullResponseBody.
func unmarshalStationFullResponseBodyToStationviewsStationFullView(v *StationFullResponseBody) *stationviews.StationFullView {
	res := &stationviews.StationFullView{
		ID:                 v.ID,
		Name:               v.Name,
		DeviceID:           v.DeviceID,
		ReadOnly:           v.ReadOnly,
		Battery:            v.Battery,
		RecordingStartedAt: v.RecordingStartedAt,
		MemoryUsed:         v.MemoryUsed,
		MemoryAvailable:    v.MemoryAvailable,
		FirmwareNumber:     v.FirmwareNumber,
		FirmwareTime:       v.FirmwareTime,
	}
	res.Owner = unmarshalStationOwnerResponseBodyToStationviewsStationOwnerView(v.Owner)
	res.Uploads = make([]*stationviews.StationUploadView, len(v.Uploads))
	for i, val := range v.Uploads {
		res.Uploads[i] = unmarshalStationUploadResponseBodyToStationviewsStationUploadView(val)
	}
	res.Images = make([]*stationviews.ImageRefView, len(v.Images))
	for i, val := range v.Images {
		res.Images[i] = unmarshalImageRefResponseBodyToStationviewsImageRefView(val)
	}
	res.Photos = unmarshalStationPhotosResponseBodyToStationviewsStationPhotosView(v.Photos)
	res.StatusJSON = make(map[string]interface{}, len(v.StatusJSON))
	for key, val := range v.StatusJSON {
		tk := key
		tv := val
		res.StatusJSON[tk] = tv
	}
	res.Modules = make([]*stationviews.StationModuleView, len(v.Modules))
	for i, val := range v.Modules {
		res.Modules[i] = unmarshalStationModuleResponseBodyToStationviewsStationModuleView(val)
	}

	return res
}
