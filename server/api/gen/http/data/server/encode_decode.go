// Code generated by goa v3.1.2, DO NOT EDIT.
//
// data HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"net/http"
	"strings"

	data "github.com/fieldkit/cloud/server/api/gen/data"
	dataviews "github.com/fieldkit/cloud/server/api/gen/data/views"
	goahttp "goa.design/goa/v3/http"
)

// EncodeDeviceSummaryResponse returns an encoder for responses returned by the
// data device summary endpoint.
func EncodeDeviceSummaryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*dataviews.DeviceDataSummaryResponse)
		enc := encoder(ctx, w)
		body := NewDeviceSummaryResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeviceSummaryRequest returns a decoder for requests sent to the data
// device summary endpoint.
func DecodeDeviceSummaryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			deviceID string
			auth     *string

			params = mux.Vars(r)
		)
		deviceID = params["deviceId"]
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		payload := NewDeviceSummaryPayload(deviceID, auth)
		if payload.Auth != nil {
			if strings.Contains(*payload.Auth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Auth, " ", 2)[1]
				payload.Auth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeviceSummaryError returns an encoder for errors returned by the
// device summary data endpoint.
func EncodeDeviceSummaryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(data.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeviceSummaryBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(data.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeviceSummaryForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(data.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeviceSummaryNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(data.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeviceSummaryUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalDataviewsDeviceProvisionSummaryViewToDeviceProvisionSummaryResponseBody
// builds a value of type *DeviceProvisionSummaryResponseBody from a value of
// type *dataviews.DeviceProvisionSummaryView.
func marshalDataviewsDeviceProvisionSummaryViewToDeviceProvisionSummaryResponseBody(v *dataviews.DeviceProvisionSummaryView) *DeviceProvisionSummaryResponseBody {
	res := &DeviceProvisionSummaryResponseBody{
		Generation: *v.Generation,
		Created:    *v.Created,
		Updated:    *v.Updated,
	}
	if v.Meta != nil {
		res.Meta = marshalDataviewsDeviceMetaSummaryViewToDeviceMetaSummaryResponseBody(v.Meta)
	}
	if v.Data != nil {
		res.Data = marshalDataviewsDeviceDataSummaryViewToDeviceDataSummaryResponseBody(v.Data)
	}

	return res
}

// marshalDataviewsDeviceMetaSummaryViewToDeviceMetaSummaryResponseBody builds
// a value of type *DeviceMetaSummaryResponseBody from a value of type
// *dataviews.DeviceMetaSummaryView.
func marshalDataviewsDeviceMetaSummaryViewToDeviceMetaSummaryResponseBody(v *dataviews.DeviceMetaSummaryView) *DeviceMetaSummaryResponseBody {
	res := &DeviceMetaSummaryResponseBody{
		Size:  *v.Size,
		First: *v.First,
		Last:  *v.Last,
	}

	return res
}

// marshalDataviewsDeviceDataSummaryViewToDeviceDataSummaryResponseBody builds
// a value of type *DeviceDataSummaryResponseBody from a value of type
// *dataviews.DeviceDataSummaryView.
func marshalDataviewsDeviceDataSummaryViewToDeviceDataSummaryResponseBody(v *dataviews.DeviceDataSummaryView) *DeviceDataSummaryResponseBody {
	res := &DeviceDataSummaryResponseBody{
		Size:  *v.Size,
		First: *v.First,
		Last:  *v.Last,
	}

	return res
}
