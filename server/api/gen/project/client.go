// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package project

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "project" service client.
type Client struct {
	AddUpdateEndpoint     goa.Endpoint
	DeleteUpdateEndpoint  goa.Endpoint
	ModifyUpdateEndpoint  goa.Endpoint
	InvitesEndpoint       goa.Endpoint
	LookupInviteEndpoint  goa.Endpoint
	AcceptInviteEndpoint  goa.Endpoint
	RejectInviteEndpoint  goa.Endpoint
	UploadMediaEndpoint   goa.Endpoint
	DownloadMediaEndpoint goa.Endpoint
}

// NewClient initializes a "project" service client given the endpoints.
func NewClient(addUpdate, deleteUpdate, modifyUpdate, invites, lookupInvite, acceptInvite, rejectInvite, uploadMedia, downloadMedia goa.Endpoint) *Client {
	return &Client{
		AddUpdateEndpoint:     addUpdate,
		DeleteUpdateEndpoint:  deleteUpdate,
		ModifyUpdateEndpoint:  modifyUpdate,
		InvitesEndpoint:       invites,
		LookupInviteEndpoint:  lookupInvite,
		AcceptInviteEndpoint:  acceptInvite,
		RejectInviteEndpoint:  rejectInvite,
		UploadMediaEndpoint:   uploadMedia,
		DownloadMediaEndpoint: downloadMedia,
	}
}

// AddUpdate calls the "add update" endpoint of the "project" service.
func (c *Client) AddUpdate(ctx context.Context, p *AddUpdatePayload) (res *ProjectUpdate, err error) {
	var ires interface{}
	ires, err = c.AddUpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ProjectUpdate), nil
}

// DeleteUpdate calls the "delete update" endpoint of the "project" service.
func (c *Client) DeleteUpdate(ctx context.Context, p *DeleteUpdatePayload) (err error) {
	_, err = c.DeleteUpdateEndpoint(ctx, p)
	return
}

// ModifyUpdate calls the "modify update" endpoint of the "project" service.
func (c *Client) ModifyUpdate(ctx context.Context, p *ModifyUpdatePayload) (res *ProjectUpdate, err error) {
	var ires interface{}
	ires, err = c.ModifyUpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ProjectUpdate), nil
}

// Invites calls the "invites" endpoint of the "project" service.
func (c *Client) Invites(ctx context.Context, p *InvitesPayload) (res *PendingInvites, err error) {
	var ires interface{}
	ires, err = c.InvitesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PendingInvites), nil
}

// LookupInvite calls the "lookup invite" endpoint of the "project" service.
func (c *Client) LookupInvite(ctx context.Context, p *LookupInvitePayload) (res *PendingInvites, err error) {
	var ires interface{}
	ires, err = c.LookupInviteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PendingInvites), nil
}

// AcceptInvite calls the "accept invite" endpoint of the "project" service.
func (c *Client) AcceptInvite(ctx context.Context, p *AcceptInvitePayload) (err error) {
	_, err = c.AcceptInviteEndpoint(ctx, p)
	return
}

// RejectInvite calls the "reject invite" endpoint of the "project" service.
func (c *Client) RejectInvite(ctx context.Context, p *RejectInvitePayload) (err error) {
	_, err = c.RejectInviteEndpoint(ctx, p)
	return
}

// UploadMedia calls the "upload media" endpoint of the "project" service.
func (c *Client) UploadMedia(ctx context.Context, p *UploadMediaPayload, req io.ReadCloser) (err error) {
	_, err = c.UploadMediaEndpoint(ctx, &UploadMediaRequestData{Payload: p, Body: req})
	return
}

// DownloadMedia calls the "download media" endpoint of the "project" service.
func (c *Client) DownloadMedia(ctx context.Context, p *DownloadMediaPayload) (res *DownloadMediaResult, resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.DownloadMediaEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*DownloadMediaResponseData)
	return o.Result, o.Body, nil
}
