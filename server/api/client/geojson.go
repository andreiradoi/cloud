// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "fieldkit": GeoJSON Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListGeoJSONPath computes a request path to the list action of GeoJSON.
func ListGeoJSONPath(project string, expedition string) string {
	param0 := project
	param1 := expedition

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/geojson", param0, param1)
}

// List a expedition's features in a GeoJSON.
func (c *Client) ListGeoJSON(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListGeoJSONRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListGeoJSONRequest create the request corresponding to the list action endpoint of the GeoJSON resource.
func (c *Client) NewListGeoJSONRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListByIDGeoJSONPath computes a request path to the list by id action of GeoJSON.
func ListByIDGeoJSONPath(featureID int) string {
	param0 := strconv.Itoa(featureID)

	return fmt.Sprintf("/features/%s/geojson", param0)
}

// List a feature's GeoJSON by id.
func (c *Client) ListByIDGeoJSON(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListByIDGeoJSONRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListByIDGeoJSONRequest create the request corresponding to the list by id action endpoint of the GeoJSON resource.
func (c *Client) NewListByIDGeoJSONRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListByInputGeoJSONPath computes a request path to the list by input action of GeoJSON.
func ListByInputGeoJSONPath(inputID int) string {
	param0 := strconv.Itoa(inputID)

	return fmt.Sprintf("/inputs/%s/geojson", param0)
}

// List an input's features in a GeoJSON.
func (c *Client) ListByInputGeoJSON(ctx context.Context, path string, descending *bool) (*http.Response, error) {
	req, err := c.NewListByInputGeoJSONRequest(ctx, path, descending)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListByInputGeoJSONRequest create the request corresponding to the list by input action endpoint of the GeoJSON resource.
func (c *Client) NewListByInputGeoJSONRequest(ctx context.Context, path string, descending *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if descending != nil {
		tmp118 := strconv.FormatBool(*descending)
		values.Set("descending", tmp118)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
