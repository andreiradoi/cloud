// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": member Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AddMemberPath computes a request path to the add action of member.
func AddMemberPath(teamID int) string {
	param0 := strconv.Itoa(teamID)

	return fmt.Sprintf("/teams/%s/members", param0)
}

// Add a member to a team
func (c *Client) AddMember(ctx context.Context, path string, payload *AddMemberPayload) (*http.Response, error) {
	req, err := c.NewAddMemberRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddMemberRequest create the request corresponding to the add action endpoint of the member resource.
func (c *Client) NewAddMemberRequest(ctx context.Context, path string, payload *AddMemberPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteMemberPath computes a request path to the delete action of member.
func DeleteMemberPath(teamID int, userID int) string {
	param0 := strconv.Itoa(teamID)
	param1 := strconv.Itoa(userID)

	return fmt.Sprintf("/teams/%s/members/%s", param0, param1)
}

// Remove a member from a team
func (c *Client) DeleteMember(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteMemberRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteMemberRequest create the request corresponding to the delete action endpoint of the member resource.
func (c *Client) NewDeleteMemberRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetMemberPath computes a request path to the get action of member.
func GetMemberPath(project string, expedition string, team string, email string) string {
	param0 := project
	param1 := expedition
	param2 := team
	param3 := email

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/teams/@/%s/members/@/%s", param0, param1, param2, param3)
}

// Get a member
func (c *Client) GetMember(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetMemberRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMemberRequest create the request corresponding to the get action endpoint of the member resource.
func (c *Client) NewGetMemberRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetIDMemberPath computes a request path to the get id action of member.
func GetIDMemberPath(teamID int, userID int) string {
	param0 := strconv.Itoa(teamID)
	param1 := strconv.Itoa(userID)

	return fmt.Sprintf("/teams/%s/members/%s", param0, param1)
}

// Get a member
func (c *Client) GetIDMember(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetIDMemberRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetIDMemberRequest create the request corresponding to the get id action endpoint of the member resource.
func (c *Client) NewGetIDMemberRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListMemberPath computes a request path to the list action of member.
func ListMemberPath(project string, expedition string, team string) string {
	param0 := project
	param1 := expedition
	param2 := team

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/teams/@/%s/members", param0, param1, param2)
}

// List an team's members
func (c *Client) ListMember(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListMemberRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListMemberRequest create the request corresponding to the list action endpoint of the member resource.
func (c *Client) NewListMemberRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListIDMemberPath computes a request path to the list id action of member.
func ListIDMemberPath(teamID int) string {
	param0 := strconv.Itoa(teamID)

	return fmt.Sprintf("/teams/%s/members", param0)
}

// List an teams's members
func (c *Client) ListIDMember(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListIDMemberRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListIDMemberRequest create the request corresponding to the list id action endpoint of the member resource.
func (c *Client) NewListIDMemberRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateMemberPath computes a request path to the update action of member.
func UpdateMemberPath(teamID int, userID int) string {
	param0 := strconv.Itoa(teamID)
	param1 := strconv.Itoa(userID)

	return fmt.Sprintf("/teams/%s/members/%s", param0, param1)
}

// Update a member
func (c *Client) UpdateMember(ctx context.Context, path string, payload *UpdateMemberPayload) (*http.Response, error) {
	req, err := c.NewUpdateMemberRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateMemberRequest create the request corresponding to the update action endpoint of the member resource.
func (c *Client) NewUpdateMemberRequest(ctx context.Context, path string, payload *UpdateMemberPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
