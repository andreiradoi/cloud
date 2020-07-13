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

// BuildAddUpdatePayload builds the payload for the project add update endpoint
// from CLI flags.
func BuildAddUpdatePayload(projectAddUpdateBody string, projectAddUpdateProjectID string, projectAddUpdateAuth string) (*project.AddUpdatePayload, error) {
	var err error
	var body AddUpdateRequestBody
	{
		err = json.Unmarshal([]byte(projectAddUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"body\": \"Eveniet dicta fuga atque maiores velit.\"\n   }'")
		}
	}
	var projectID int32
	{
		var v int64
		v, err = strconv.ParseInt(projectAddUpdateProjectID, 10, 32)
		projectID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be INT32")
		}
	}
	var auth string
	{
		auth = projectAddUpdateAuth
	}
	v := &project.AddUpdatePayload{
		Body: body.Body,
	}
	v.ProjectID = projectID
	v.Auth = auth

	return v, nil
}

// BuildDeleteUpdatePayload builds the payload for the project delete update
// endpoint from CLI flags.
func BuildDeleteUpdatePayload(projectDeleteUpdateProjectID string, projectDeleteUpdateUpdateID string, projectDeleteUpdateAuth string) (*project.DeleteUpdatePayload, error) {
	var err error
	var projectID int32
	{
		var v int64
		v, err = strconv.ParseInt(projectDeleteUpdateProjectID, 10, 32)
		projectID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be INT32")
		}
	}
	var updateID int64
	{
		updateID, err = strconv.ParseInt(projectDeleteUpdateUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for updateID, must be INT64")
		}
	}
	var auth string
	{
		auth = projectDeleteUpdateAuth
	}
	v := &project.DeleteUpdatePayload{}
	v.ProjectID = projectID
	v.UpdateID = updateID
	v.Auth = auth

	return v, nil
}

// BuildModifyUpdatePayload builds the payload for the project modify update
// endpoint from CLI flags.
func BuildModifyUpdatePayload(projectModifyUpdateBody string, projectModifyUpdateProjectID string, projectModifyUpdateUpdateID string, projectModifyUpdateAuth string) (*project.ModifyUpdatePayload, error) {
	var err error
	var body ModifyUpdateRequestBody
	{
		err = json.Unmarshal([]byte(projectModifyUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"body\": \"Repudiandae temporibus a facilis earum.\"\n   }'")
		}
	}
	var projectID int32
	{
		var v int64
		v, err = strconv.ParseInt(projectModifyUpdateProjectID, 10, 32)
		projectID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be INT32")
		}
	}
	var updateID int64
	{
		updateID, err = strconv.ParseInt(projectModifyUpdateUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for updateID, must be INT64")
		}
	}
	var auth string
	{
		auth = projectModifyUpdateAuth
	}
	v := &project.ModifyUpdatePayload{
		Body: body.Body,
	}
	v.ProjectID = projectID
	v.UpdateID = updateID
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

// BuildUploadMediaPayload builds the payload for the project upload media
// endpoint from CLI flags.
func BuildUploadMediaPayload(projectUploadMediaProjectID string, projectUploadMediaContentType string, projectUploadMediaContentLength string, projectUploadMediaAuth string) (*project.UploadMediaPayload, error) {
	var err error
	var projectID int32
	{
		var v int64
		v, err = strconv.ParseInt(projectUploadMediaProjectID, 10, 32)
		projectID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be INT32")
		}
	}
	var contentType string
	{
		contentType = projectUploadMediaContentType
	}
	var contentLength int64
	{
		contentLength, err = strconv.ParseInt(projectUploadMediaContentLength, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for contentLength, must be INT64")
		}
	}
	var auth string
	{
		auth = projectUploadMediaAuth
	}
	v := &project.UploadMediaPayload{}
	v.ProjectID = projectID
	v.ContentType = contentType
	v.ContentLength = contentLength
	v.Auth = auth

	return v, nil
}

// BuildDownloadMediaPayload builds the payload for the project download media
// endpoint from CLI flags.
func BuildDownloadMediaPayload(projectDownloadMediaProjectID string) (*project.DownloadMediaPayload, error) {
	var err error
	var projectID int32
	{
		var v int64
		v, err = strconv.ParseInt(projectDownloadMediaProjectID, 10, 32)
		projectID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be INT32")
		}
	}
	v := &project.DownloadMediaPayload{}
	v.ProjectID = projectID

	return v, nil
}
