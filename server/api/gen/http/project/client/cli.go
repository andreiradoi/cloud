// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	project "github.com/fieldkit/cloud/server/api/gen/project"
)

// BuildUpdatePayload builds the payload for the project update endpoint from
// CLI flags.
func BuildUpdatePayload(projectUpdateBody string, projectUpdateID string, projectUpdateAuth string) (*project.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(projectUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"body\": \"Sint quasi possimus et.\"\n   }'")
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(projectUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var auth string
	{
		auth = projectUpdateAuth
	}
	v := &project.UpdatePayload{
		Body: body.Body,
	}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildInvitesPayload builds the payload for the project invites endpoint from
// CLI flags.
func BuildInvitesPayload(projectInvitesAuth string) (*project.InvitesPayload, error) {
	var auth string
	{
		auth = projectInvitesAuth
	}
	v := &project.InvitesPayload{}
	v.Auth = auth

	return v, nil
}

// BuildLookupInvitePayload builds the payload for the project lookup invite
// endpoint from CLI flags.
func BuildLookupInvitePayload(projectLookupInviteToken string, projectLookupInviteAuth string) (*project.LookupInvitePayload, error) {
	var token string
	{
		token = projectLookupInviteToken
	}
	var auth string
	{
		auth = projectLookupInviteAuth
	}
	v := &project.LookupInvitePayload{}
	v.Token = token
	v.Auth = auth

	return v, nil
}

// BuildAcceptInvitePayload builds the payload for the project accept invite
// endpoint from CLI flags.
func BuildAcceptInvitePayload(projectAcceptInviteID string, projectAcceptInviteToken string, projectAcceptInviteAuth string) (*project.AcceptInvitePayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(projectAcceptInviteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token *string
	{
		if projectAcceptInviteToken != "" {
			token = &projectAcceptInviteToken
		}
	}
	var auth string
	{
		auth = projectAcceptInviteAuth
	}
	v := &project.AcceptInvitePayload{}
	v.ID = id
	v.Token = token
	v.Auth = auth

	return v, nil
}

// BuildRejectInvitePayload builds the payload for the project reject invite
// endpoint from CLI flags.
func BuildRejectInvitePayload(projectRejectInviteID string, projectRejectInviteToken string, projectRejectInviteAuth string) (*project.RejectInvitePayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(projectRejectInviteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token *string
	{
		if projectRejectInviteToken != "" {
			token = &projectRejectInviteToken
		}
	}
	var auth string
	{
		auth = projectRejectInviteAuth
	}
	v := &project.RejectInvitePayload{}
	v.ID = id
	v.Token = token
	v.Auth = auth

	return v, nil
}
