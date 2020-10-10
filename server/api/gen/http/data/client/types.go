// Code generated by goa v3.2.4, DO NOT EDIT.
//
// data HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	dataviews "github.com/fieldkit/cloud/server/api/gen/data/views"
	goa "goa.design/goa/v3/pkg"
)

// DeviceSummaryResponseBody is the type of the "data" service "device summary"
// endpoint HTTP response body.
type DeviceSummaryResponseBody struct {
	Provisions DeviceProvisionSummaryCollectionResponseBody `form:"provisions,omitempty" json:"provisions,omitempty" xml:"provisions,omitempty"`
}

// DeviceSummaryUnauthorizedResponseBody is the type of the "data" service
// "device summary" endpoint HTTP response body for the "unauthorized" error.
type DeviceSummaryUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeviceSummaryForbiddenResponseBody is the type of the "data" service "device
// summary" endpoint HTTP response body for the "forbidden" error.
type DeviceSummaryForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeviceSummaryNotFoundResponseBody is the type of the "data" service "device
// summary" endpoint HTTP response body for the "not-found" error.
type DeviceSummaryNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeviceSummaryBadRequestResponseBody is the type of the "data" service
// "device summary" endpoint HTTP response body for the "bad-request" error.
type DeviceSummaryBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

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

// NewDeviceSummaryUnauthorized builds a data service device summary endpoint
// unauthorized error.
func NewDeviceSummaryUnauthorized(body *DeviceSummaryUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeviceSummaryForbidden builds a data service device summary endpoint
// forbidden error.
func NewDeviceSummaryForbidden(body *DeviceSummaryForbiddenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeviceSummaryNotFound builds a data service device summary endpoint
// not-found error.
func NewDeviceSummaryNotFound(body *DeviceSummaryNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeviceSummaryBadRequest builds a data service device summary endpoint
// bad-request error.
func NewDeviceSummaryBadRequest(body *DeviceSummaryBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateDeviceSummaryUnauthorizedResponseBody runs the validations defined
// on device summary_unauthorized_response_body
func ValidateDeviceSummaryUnauthorizedResponseBody(body *DeviceSummaryUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeviceSummaryForbiddenResponseBody runs the validations defined on
// device summary_forbidden_response_body
func ValidateDeviceSummaryForbiddenResponseBody(body *DeviceSummaryForbiddenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeviceSummaryNotFoundResponseBody runs the validations defined on
// device summary_not-found_response_body
func ValidateDeviceSummaryNotFoundResponseBody(body *DeviceSummaryNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeviceSummaryBadRequestResponseBody runs the validations defined on
// device summary_bad-request_response_body
func ValidateDeviceSummaryBadRequestResponseBody(body *DeviceSummaryBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
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
