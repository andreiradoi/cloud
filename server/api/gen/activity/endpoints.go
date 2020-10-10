// Code generated by goa v3.2.4, DO NOT EDIT.
//
// activity endpoints
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package activity

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "activity" service endpoints.
type Endpoints struct {
	Station goa.Endpoint
	Project goa.Endpoint
}

// NewEndpoints wraps the methods of the "activity" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Station: NewStationEndpoint(s, a.JWTAuth),
		Project: NewProjectEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "activity" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Station = m(e.Station)
	e.Project = m(e.Project)
}

// NewStationEndpoint returns an endpoint function that calls the method
// "station" of service "activity".
func NewStationEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StationPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access", "api:admin", "api:ingestion"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Auth != nil {
			token = *p.Auth
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Station(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationActivityPage(res, "default")
		return vres, nil
	}
}

// NewProjectEndpoint returns an endpoint function that calls the method
// "project" of service "activity".
func NewProjectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ProjectPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access", "api:admin", "api:ingestion"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Auth != nil {
			token = *p.Auth
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Project(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProjectActivityPage(res, "default")
		return vres, nil
	}
}
