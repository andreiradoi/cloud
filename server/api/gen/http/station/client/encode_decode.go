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

	station "github.com/fieldkit/cloud/server/api/gen/station"
	stationviews "github.com/fieldkit/cloud/server/api/gen/station/views"
	goahttp "goa.design/goa/v3/http"
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
	req, err := http.NewRequest("POST", u.String(), nil)
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
		ID:       v.ID,
		Name:     v.Name,
		Position: v.Position,
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

	return res
}
