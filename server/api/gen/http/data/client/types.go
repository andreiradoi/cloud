// Code generated by goa v3.1.2, DO NOT EDIT.
//
// data HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	data "github.com/fieldkit/cloud/server/api/gen/data"
	dataviews "github.com/fieldkit/cloud/server/api/gen/data/views"
	goa "goa.design/goa/v3/pkg"
)

// DeviceSummaryResponseBody is the type of the "data" service "device summary"
// endpoint HTTP response body.
type DeviceSummaryResponseBody struct {
	Provisions DeviceProvisionSummaryCollectionResponseBody `form:"provisions,omitempty" json:"provisions,omitempty" xml:"provisions,omitempty"`
}

// DeviceSummaryBadRequestResponseBody is the type of the "data" service
// "device summary" endpoint HTTP response body for the "bad-request" error.
type DeviceSummaryBadRequestResponseBody string

// DeviceSummaryForbiddenResponseBody is the type of the "data" service "device
// summary" endpoint HTTP response body for the "forbidden" error.
type DeviceSummaryForbiddenResponseBody string

// DeviceSummaryNotFoundResponseBody is the type of the "data" service "device
// summary" endpoint HTTP response body for the "not-found" error.
type DeviceSummaryNotFoundResponseBody string

// DeviceSummaryUnauthorizedResponseBody is the type of the "data" service
// "device summary" endpoint HTTP response body for the "unauthorized" error.
type DeviceSummaryUnauthorizedResponseBody string

// DeviceProvisionSummaryCollectionResponseBody is used to define fields on
// response body types.
type DeviceProvisionSummaryCollectionResponseBody []*DeviceProvisionSummaryResponseBody

// DeviceProvisionSummaryResponseBody is used to define fields on response body
// types.
type DeviceProvisionSummaryResponseBody struct {
	Generation *string                        `form:"generation,omitempty" json:"generation,omitempty" xml:"generation,omitempty"`
	Created    *int64                         `form:"created,omitempty" json:"created,omitempty" xml:"created,omitempty"`
	Updated    *int64                         `form:"updated,omitempty" json:"updated,omitempty" xml:"updated,omitempty"`
	Meta       *DeviceMetaSummaryResponseBody `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Data       *DeviceDataSummaryResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// DeviceMetaSummaryResponseBody is used to define fields on response body
// types.
type DeviceMetaSummaryResponseBody struct {
	Size  *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	First *int64 `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Last  *int64 `form:"last,omitempty" json:"last,omitempty" xml:"last,omitempty"`
}

// DeviceDataSummaryResponseBody is used to define fields on response body
// types.
type DeviceDataSummaryResponseBody struct {
	Size  *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	First *int64 `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Last  *int64 `form:"last,omitempty" json:"last,omitempty" xml:"last,omitempty"`
}

// NewDeviceSummaryDeviceDataSummaryResponseOK builds a "data" service "device
// summary" endpoint result from a HTTP "OK" response.
func NewDeviceSummaryDeviceDataSummaryResponseOK(body *DeviceSummaryResponseBody) *dataviews.DeviceDataSummaryResponseView {
	v := &dataviews.DeviceDataSummaryResponseView{}
	v.Provisions = make([]*dataviews.DeviceProvisionSummaryView, len(body.Provisions))
	for i, val := range body.Provisions {
		v.Provisions[i] = unmarshalDeviceProvisionSummaryResponseBodyToDataviewsDeviceProvisionSummaryView(val)
	}

	return v
}

// NewDeviceSummaryBadRequest builds a data service device summary endpoint
// bad-request error.
func NewDeviceSummaryBadRequest(body DeviceSummaryBadRequestResponseBody) data.BadRequest {
	v := data.BadRequest(body)
	return v
}

// NewDeviceSummaryForbidden builds a data service device summary endpoint
// forbidden error.
func NewDeviceSummaryForbidden(body DeviceSummaryForbiddenResponseBody) data.Forbidden {
	v := data.Forbidden(body)
	return v
}

// NewDeviceSummaryNotFound builds a data service device summary endpoint
// not-found error.
func NewDeviceSummaryNotFound(body DeviceSummaryNotFoundResponseBody) data.NotFound {
	v := data.NotFound(body)
	return v
}

// NewDeviceSummaryUnauthorized builds a data service device summary endpoint
// unauthorized error.
func NewDeviceSummaryUnauthorized(body DeviceSummaryUnauthorizedResponseBody) data.Unauthorized {
	v := data.Unauthorized(body)
	return v
}

// ValidateDeviceProvisionSummaryCollectionResponseBody runs the validations
// defined on DeviceProvisionSummaryCollectionResponseBody
func ValidateDeviceProvisionSummaryCollectionResponseBody(body DeviceProvisionSummaryCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateDeviceProvisionSummaryResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateDeviceProvisionSummaryResponseBody runs the validations defined on
// DeviceProvisionSummaryResponseBody
func ValidateDeviceProvisionSummaryResponseBody(body *DeviceProvisionSummaryResponseBody) (err error) {
	if body.Generation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("generation", "body"))
	}
	if body.Created == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created", "body"))
	}
	if body.Updated == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated", "body"))
	}
	if body.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Meta != nil {
		if err2 := ValidateDeviceMetaSummaryResponseBody(body.Meta); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Data != nil {
		if err2 := ValidateDeviceDataSummaryResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeviceMetaSummaryResponseBody runs the validations defined on
// DeviceMetaSummaryResponseBody
func ValidateDeviceMetaSummaryResponseBody(body *DeviceMetaSummaryResponseBody) (err error) {
	if body.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "body"))
	}
	if body.First == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first", "body"))
	}
	if body.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "body"))
	}
	return
}

// ValidateDeviceDataSummaryResponseBody runs the validations defined on
// DeviceDataSummaryResponseBody
func ValidateDeviceDataSummaryResponseBody(body *DeviceDataSummaryResponseBody) (err error) {
	if body.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "body"))
	}
	if body.First == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first", "body"))
	}
	if body.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "body"))
	}
	return
}
