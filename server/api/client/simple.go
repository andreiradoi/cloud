// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": simple Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// MyCsvDataSimplePath computes a request path to the my csv data action of simple.
func MyCsvDataSimplePath() string {

	return fmt.Sprintf("/my/simple/data/csv")
}

// MyCsvDataSimple makes a request to the my csv data action endpoint of the simple resource
func (c *Client) MyCsvDataSimple(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMyCsvDataSimpleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMyCsvDataSimpleRequest create the request corresponding to the my csv data action endpoint of the simple resource.
func (c *Client) NewMyCsvDataSimpleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// MyFeaturesSimplePath computes a request path to the my features action of simple.
func MyFeaturesSimplePath() string {

	return fmt.Sprintf("/my/simple/features")
}

// MyFeaturesSimple makes a request to the my features action endpoint of the simple resource
func (c *Client) MyFeaturesSimple(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMyFeaturesSimpleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMyFeaturesSimpleRequest create the request corresponding to the my features action endpoint of the simple resource.
func (c *Client) NewMyFeaturesSimpleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
