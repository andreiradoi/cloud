// Code generated by goa v3.1.1, DO NOT EDIT.
//
// activity HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	activity "github.com/fieldkit/cloud/server/api/gen/activity"
	activityviews "github.com/fieldkit/cloud/server/api/gen/activity/views"
)

// StationResponseBody is the type of the "activity" service "station" endpoint
// HTTP response body.
type StationResponseBody struct {
	Activities StationActivityResponseBodyCollection `form:"activities" json:"activities" xml:"activities"`
	Total      int32                                 `form:"total" json:"total" xml:"total"`
	Page       int32                                 `form:"page" json:"page" xml:"page"`
}

// ProjectResponseBody is the type of the "activity" service "project" endpoint
// HTTP response body.
type ProjectResponseBody struct {
	Activities ProjectActivityResponseBodyCollection `form:"activities" json:"activities" xml:"activities"`
	Total      int32                                 `form:"total" json:"total" xml:"total"`
	Page       int32                                 `form:"page" json:"page" xml:"page"`
}

// StationActivityResponseBodyCollection is used to define fields on response
// body types.
type StationActivityResponseBodyCollection []*StationActivityResponseBody

// StationActivityResponseBody is used to define fields on response body types.
type StationActivityResponseBody struct {
	ID        int64                       `form:"id" json:"id" xml:"id"`
	Station   *StationSummaryResponseBody `form:"station" json:"station" xml:"station"`
	CreatedAt int64                       `form:"created_at" json:"created_at" xml:"created_at"`
	Type      string                      `form:"type" json:"type" xml:"type"`
	Meta      interface{}                 `form:"meta" json:"meta" xml:"meta"`
}

// StationSummaryResponseBody is used to define fields on response body types.
type StationSummaryResponseBody struct {
	ID   int64  `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// ProjectActivityResponseBodyCollection is used to define fields on response
// body types.
type ProjectActivityResponseBodyCollection []*ProjectActivityResponseBody

// ProjectActivityResponseBody is used to define fields on response body types.
type ProjectActivityResponseBody struct {
	ID        int64                       `form:"id" json:"id" xml:"id"`
	Project   *ProjectSummaryResponseBody `form:"project" json:"project" xml:"project"`
	CreatedAt int64                       `form:"created_at" json:"created_at" xml:"created_at"`
	Type      string                      `form:"type" json:"type" xml:"type"`
	Meta      interface{}                 `form:"meta" json:"meta" xml:"meta"`
}

// ProjectSummaryResponseBody is used to define fields on response body types.
type ProjectSummaryResponseBody struct {
	ID   int64  `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// NewStationResponseBody builds the HTTP response body from the result of the
// "station" endpoint of the "activity" service.
func NewStationResponseBody(res *activityviews.StationActivityPageView) *StationResponseBody {
	body := &StationResponseBody{
		Total: *res.Total,
		Page:  *res.Page,
	}
	if res.Activities != nil {
		body.Activities = make([]*StationActivityResponseBody, len(res.Activities))
		for i, val := range res.Activities {
			body.Activities[i] = marshalActivityviewsStationActivityViewToStationActivityResponseBody(val)
		}
	}
	return body
}

// NewProjectResponseBody builds the HTTP response body from the result of the
// "project" endpoint of the "activity" service.
func NewProjectResponseBody(res *activityviews.ProjectActivityPageView) *ProjectResponseBody {
	body := &ProjectResponseBody{
		Total: *res.Total,
		Page:  *res.Page,
	}
	if res.Activities != nil {
		body.Activities = make([]*ProjectActivityResponseBody, len(res.Activities))
		for i, val := range res.Activities {
			body.Activities[i] = marshalActivityviewsProjectActivityViewToProjectActivityResponseBody(val)
		}
	}
	return body
}

// NewStationPayload builds a activity service station endpoint payload.
func NewStationPayload(id int64, page *int64, auth *string) *activity.StationPayload {
	v := &activity.StationPayload{}
	v.ID = &id
	v.Page = page
	v.Auth = auth

	return v
}

// NewProjectPayload builds a activity service project endpoint payload.
func NewProjectPayload(id int64, page *int64, auth *string) *activity.ProjectPayload {
	v := &activity.ProjectPayload{}
	v.ID = &id
	v.Page = page
	v.Auth = auth

	return v
}
