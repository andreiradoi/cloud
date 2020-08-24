// Code generated by goa v3.1.2, DO NOT EDIT.
//
// export HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	export "github.com/fieldkit/cloud/server/api/gen/export"
	exportviews "github.com/fieldkit/cloud/server/api/gen/export/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListMineRequest instantiates a HTTP request object with method and path
// set to call the "export" service "list mine" endpoint
func (c *Client) BuildListMineRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMineExportPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("export", "list mine", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListMineRequest returns an encoder for requests sent to the export
// list mine server.
func EncodeListMineRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*export.ListMinePayload)
		if !ok {
			return goahttp.ErrInvalidType("export", "list mine", "*export.ListMinePayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListMineResponse returns a decoder for responses returned by the
// export list mine endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListMineResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad-request" (type *goa.ServiceError): http.StatusBadRequest
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
				return nil, goahttp.ErrDecodingError("export", "list mine", err)
			}
			p := NewListMineUserExportsOK(&body)
			view := "default"
			vres := &exportviews.UserExports{Projected: p, View: view}
			if err = exportviews.ValidateUserExports(vres); err != nil {
				return nil, goahttp.ErrValidationError("export", "list mine", err)
			}
			res := export.NewUserExports(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListMineUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "list mine", err)
			}
			err = ValidateListMineUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "list mine", err)
			}
			return nil, NewListMineUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ListMineForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "list mine", err)
			}
			err = ValidateListMineForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "list mine", err)
			}
			return nil, NewListMineForbidden(&body)
		case http.StatusNotFound:
			var (
				body ListMineNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "list mine", err)
			}
			err = ValidateListMineNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "list mine", err)
			}
			return nil, NewListMineNotFound(&body)
		case http.StatusBadRequest:
			var (
				body ListMineBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "list mine", err)
			}
			err = ValidateListMineBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "list mine", err)
			}
			return nil, NewListMineBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("export", "list mine", resp.StatusCode, string(body))
		}
	}
}

// BuildStatusRequest instantiates a HTTP request object with method and path
// set to call the "export" service "status" endpoint
func (c *Client) BuildStatusRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*export.StatusPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("export", "status", "*export.StatusPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatusExportPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("export", "status", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStatusRequest returns an encoder for requests sent to the export
// status server.
func EncodeStatusRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*export.StatusPayload)
		if !ok {
			return goahttp.ErrInvalidType("export", "status", "*export.StatusPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeStatusResponse returns a decoder for responses returned by the export
// status endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeStatusResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad-request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body StatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "status", err)
			}
			p := NewStatusExportStatusOK(&body)
			view := "default"
			vres := &exportviews.ExportStatus{Projected: p, View: view}
			if err = exportviews.ValidateExportStatus(vres); err != nil {
				return nil, goahttp.ErrValidationError("export", "status", err)
			}
			res := export.NewExportStatus(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body StatusUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "status", err)
			}
			err = ValidateStatusUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "status", err)
			}
			return nil, NewStatusUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body StatusForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "status", err)
			}
			err = ValidateStatusForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "status", err)
			}
			return nil, NewStatusForbidden(&body)
		case http.StatusNotFound:
			var (
				body StatusNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "status", err)
			}
			err = ValidateStatusNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "status", err)
			}
			return nil, NewStatusNotFound(&body)
		case http.StatusBadRequest:
			var (
				body StatusBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "status", err)
			}
			err = ValidateStatusBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "status", err)
			}
			return nil, NewStatusBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("export", "status", resp.StatusCode, string(body))
		}
	}
}

// BuildDownloadRequest instantiates a HTTP request object with method and path
// set to call the "export" service "download" endpoint
func (c *Client) BuildDownloadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*export.DownloadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("export", "download", "*export.DownloadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DownloadExportPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("export", "download", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDownloadRequest returns an encoder for requests sent to the export
// download server.
func EncodeDownloadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*export.DownloadPayload)
		if !ok {
			return goahttp.ErrInvalidType("export", "download", "*export.DownloadPayload", v)
		}
		values := req.URL.Query()
		values.Add("auth", p.Auth)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeDownloadResponse returns a decoder for responses returned by the
// export download endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDownloadResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad-request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeDownloadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				length             int64
				contentType        string
				contentDisposition string
				err                error
			)
			{
				lengthRaw := resp.Header.Get("Content-Length")
				if lengthRaw == "" {
					return nil, goahttp.ErrValidationError("export", "download", goa.MissingFieldError("Content-Length", "header"))
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
			contentDispositionRaw := resp.Header.Get("Content-Disposition")
			if contentDispositionRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Disposition", "header"))
			}
			contentDisposition = contentDispositionRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "download", err)
			}
			res := NewDownloadResultOK(length, contentType, contentDisposition)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body DownloadUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "download", err)
			}
			err = ValidateDownloadUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "download", err)
			}
			return nil, NewDownloadUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body DownloadForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "download", err)
			}
			err = ValidateDownloadForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "download", err)
			}
			return nil, NewDownloadForbidden(&body)
		case http.StatusNotFound:
			var (
				body DownloadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "download", err)
			}
			err = ValidateDownloadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "download", err)
			}
			return nil, NewDownloadNotFound(&body)
		case http.StatusBadRequest:
			var (
				body DownloadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "download", err)
			}
			err = ValidateDownloadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "download", err)
			}
			return nil, NewDownloadBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("export", "download", resp.StatusCode, string(body))
		}
	}
}

// BuildCsvRequest instantiates a HTTP request object with method and path set
// to call the "export" service "csv" endpoint
func (c *Client) BuildCsvRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CsvExportPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("export", "csv", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCsvRequest returns an encoder for requests sent to the export csv
// server.
func EncodeCsvRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*export.CsvPayload)
		if !ok {
			return goahttp.ErrInvalidType("export", "csv", "*export.CsvPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		if p.Start != nil {
			values.Add("start", fmt.Sprintf("%v", *p.Start))
		}
		if p.End != nil {
			values.Add("end", fmt.Sprintf("%v", *p.End))
		}
		if p.Stations != nil {
			values.Add("stations", *p.Stations)
		}
		if p.Sensors != nil {
			values.Add("sensors", *p.Sensors)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeCsvResponse returns a decoder for responses returned by the export csv
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeCsvResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad-request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCsvResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusFound:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Location", "header"))
			}
			location = locationRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "csv", err)
			}
			res := NewCsvResultFound(location)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CsvUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "csv", err)
			}
			err = ValidateCsvUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "csv", err)
			}
			return nil, NewCsvUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body CsvForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "csv", err)
			}
			err = ValidateCsvForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "csv", err)
			}
			return nil, NewCsvForbidden(&body)
		case http.StatusNotFound:
			var (
				body CsvNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "csv", err)
			}
			err = ValidateCsvNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "csv", err)
			}
			return nil, NewCsvNotFound(&body)
		case http.StatusBadRequest:
			var (
				body CsvBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "csv", err)
			}
			err = ValidateCsvBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "csv", err)
			}
			return nil, NewCsvBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("export", "csv", resp.StatusCode, string(body))
		}
	}
}

// BuildJSONLinesRequest instantiates a HTTP request object with method and
// path set to call the "export" service "json lines" endpoint
func (c *Client) BuildJSONLinesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: JSONLinesExportPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("export", "json lines", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeJSONLinesRequest returns an encoder for requests sent to the export
// json lines server.
func EncodeJSONLinesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*export.JSONLinesPayload)
		if !ok {
			return goahttp.ErrInvalidType("export", "json lines", "*export.JSONLinesPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		if p.Start != nil {
			values.Add("start", fmt.Sprintf("%v", *p.Start))
		}
		if p.End != nil {
			values.Add("end", fmt.Sprintf("%v", *p.End))
		}
		if p.Stations != nil {
			values.Add("stations", *p.Stations)
		}
		if p.Sensors != nil {
			values.Add("sensors", *p.Sensors)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeJSONLinesResponse returns a decoder for responses returned by the
// export json lines endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeJSONLinesResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "bad-request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeJSONLinesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusFound:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Location", "header"))
			}
			location = locationRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "json lines", err)
			}
			res := NewJSONLinesResultFound(location)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body JSONLinesUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "json lines", err)
			}
			err = ValidateJSONLinesUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "json lines", err)
			}
			return nil, NewJSONLinesUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body JSONLinesForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "json lines", err)
			}
			err = ValidateJSONLinesForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "json lines", err)
			}
			return nil, NewJSONLinesForbidden(&body)
		case http.StatusNotFound:
			var (
				body JSONLinesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "json lines", err)
			}
			err = ValidateJSONLinesNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "json lines", err)
			}
			return nil, NewJSONLinesNotFound(&body)
		case http.StatusBadRequest:
			var (
				body JSONLinesBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("export", "json lines", err)
			}
			err = ValidateJSONLinesBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("export", "json lines", err)
			}
			return nil, NewJSONLinesBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("export", "json lines", resp.StatusCode, string(body))
		}
	}
}

// unmarshalExportStatusResponseBodyToExportviewsExportStatusView builds a
// value of type *exportviews.ExportStatusView from a value of type
// *ExportStatusResponseBody.
func unmarshalExportStatusResponseBodyToExportviewsExportStatusView(v *ExportStatusResponseBody) *exportviews.ExportStatusView {
	res := &exportviews.ExportStatusView{
		ID:          v.ID,
		Token:       v.Token,
		CreatedAt:   v.CreatedAt,
		CompletedAt: v.CompletedAt,
		Kind:        v.Kind,
		Progress:    v.Progress,
		StatusURL:   v.StatusURL,
		DownloadURL: v.DownloadURL,
		Size:        v.Size,
		Args:        v.Args,
	}

	return res
}
