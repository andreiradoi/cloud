// Code generated by goa v3.1.2, DO NOT EDIT.
//
// export HTTP server
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"regexp"

	export "github.com/fieldkit/cloud/server/api/gen/export"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the export service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	ListMine  http.Handler
	Status    http.Handler
	Download  http.Handler
	Csv       http.Handler
	JSONLines http.Handler
	CORS      http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the export service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *export.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListMine", "GET", "/export"},
			{"Status", "GET", "/export/{id}"},
			{"Download", "GET", "/export/{id}/download"},
			{"Csv", "POST", "/export/csv"},
			{"JSONLines", "POST", "/export/json-lines"},
			{"CORS", "OPTIONS", "/export"},
			{"CORS", "OPTIONS", "/export/{id}"},
			{"CORS", "OPTIONS", "/export/{id}/download"},
			{"CORS", "OPTIONS", "/export/csv"},
			{"CORS", "OPTIONS", "/export/json-lines"},
		},
		ListMine:  NewListMineHandler(e.ListMine, mux, decoder, encoder, errhandler, formatter),
		Status:    NewStatusHandler(e.Status, mux, decoder, encoder, errhandler, formatter),
		Download:  NewDownloadHandler(e.Download, mux, decoder, encoder, errhandler, formatter),
		Csv:       NewCsvHandler(e.Csv, mux, decoder, encoder, errhandler, formatter),
		JSONLines: NewJSONLinesHandler(e.JSONLines, mux, decoder, encoder, errhandler, formatter),
		CORS:      NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "export" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListMine = m(s.ListMine)
	s.Status = m(s.Status)
	s.Download = m(s.Download)
	s.Csv = m(s.Csv)
	s.JSONLines = m(s.JSONLines)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the export endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListMineHandler(mux, h.ListMine)
	MountStatusHandler(mux, h.Status)
	MountDownloadHandler(mux, h.Download)
	MountCsvHandler(mux, h.Csv)
	MountJSONLinesHandler(mux, h.JSONLines)
	MountCORSHandler(mux, h.CORS)
}

// MountListMineHandler configures the mux to serve the "export" service "list
// mine" endpoint.
func MountListMineHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleExportOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/export", f)
}

// NewListMineHandler creates a HTTP handler which loads the HTTP request and
// calls the "export" service "list mine" endpoint.
func NewListMineHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListMineRequest(mux, decoder)
		encodeResponse = EncodeListMineResponse(encoder)
		encodeError    = EncodeListMineError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list mine")
		ctx = context.WithValue(ctx, goa.ServiceKey, "export")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountStatusHandler configures the mux to serve the "export" service "status"
// endpoint.
func MountStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleExportOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/export/{id}", f)
}

// NewStatusHandler creates a HTTP handler which loads the HTTP request and
// calls the "export" service "status" endpoint.
func NewStatusHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeStatusRequest(mux, decoder)
		encodeResponse = EncodeStatusResponse(encoder)
		encodeError    = EncodeStatusError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "status")
		ctx = context.WithValue(ctx, goa.ServiceKey, "export")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDownloadHandler configures the mux to serve the "export" service
// "download" endpoint.
func MountDownloadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleExportOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/export/{id}/download", f)
}

// NewDownloadHandler creates a HTTP handler which loads the HTTP request and
// calls the "export" service "download" endpoint.
func NewDownloadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDownloadRequest(mux, decoder)
		encodeResponse = EncodeDownloadResponse(encoder)
		encodeError    = EncodeDownloadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "download")
		ctx = context.WithValue(ctx, goa.ServiceKey, "export")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		o := res.(*export.DownloadResponseData)
		defer o.Body.Close()
		if err := encodeResponse(ctx, w, o.Result); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, o.Body); err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// MountCsvHandler configures the mux to serve the "export" service "csv"
// endpoint.
func MountCsvHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleExportOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/export/csv", f)
}

// NewCsvHandler creates a HTTP handler which loads the HTTP request and calls
// the "export" service "csv" endpoint.
func NewCsvHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCsvRequest(mux, decoder)
		encodeResponse = EncodeCsvResponse(encoder)
		encodeError    = EncodeCsvError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "csv")
		ctx = context.WithValue(ctx, goa.ServiceKey, "export")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountJSONLinesHandler configures the mux to serve the "export" service "json
// lines" endpoint.
func MountJSONLinesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleExportOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/export/json-lines", f)
}

// NewJSONLinesHandler creates a HTTP handler which loads the HTTP request and
// calls the "export" service "json lines" endpoint.
func NewJSONLinesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeJSONLinesRequest(mux, decoder)
		encodeResponse = EncodeJSONLinesResponse(encoder)
		encodeError    = EncodeJSONLinesError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "json lines")
		ctx = context.WithValue(ctx, goa.ServiceKey, "export")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service export.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleExportOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/export", f)
	mux.Handle("OPTIONS", "/export/{id}", f)
	mux.Handle("OPTIONS", "/export/{id}/download", f)
	mux.Handle("OPTIONS", "/export/csv", f)
	mux.Handle("OPTIONS", "/export/json-lines", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleExportOrigin applies the CORS response headers corresponding to the
// origin for the service export.
func handleExportOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("(.+[.])?127.0.0.1:\\d+")
	spec1 := regexp.MustCompile("(.+[.])?fieldkit.org:\\d+")
	spec2 := regexp.MustCompile("(.+[.])?local.fkdev.org:\\d+")
	spec3 := regexp.MustCompile("(.+[.])?localhost:\\d+")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec1) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec2) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec3) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
