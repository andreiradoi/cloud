// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": jsonData Resource Client
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

// GetJSONDataPath computes a request path to the get action of jsonData.
func GetJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data/json", param0)
}

// Retrieve data
func (c *Client) GetJSONData(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Response, error) {
	req, err := c.NewGetJSONDataRequest(ctx, path, end, internal, page, pageSize, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetJSONDataRequest create the request corresponding to the get action endpoint of the jsonData resource.
func (c *Client) NewGetJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp127 := strconv.Itoa(*end)
		values.Set("end", tmp127)
	}
	if internal != nil {
		tmp128 := strconv.FormatBool(*internal)
		values.Set("internal", tmp128)
	}
	if page != nil {
		tmp129 := strconv.Itoa(*page)
		values.Set("page", tmp129)
	}
	if pageSize != nil {
		tmp130 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp130)
	}
	if start != nil {
		tmp131 := strconv.Itoa(*start)
		values.Set("start", tmp131)
	}
	u.RawQuery = values.Encode()
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

// SummaryJSONDataPath computes a request path to the summary action of jsonData.
func SummaryJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/summary/json", param0)
}

// Retrieve summarized data
func (c *Client) SummaryJSONData(ctx context.Context, path string, end *int, internal *bool, interval *int, page *int, pageSize *int, resolution *int, start *int) (*http.Response, error) {
	req, err := c.NewSummaryJSONDataRequest(ctx, path, end, internal, interval, page, pageSize, resolution, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSummaryJSONDataRequest create the request corresponding to the summary action endpoint of the jsonData resource.
func (c *Client) NewSummaryJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, interval *int, page *int, pageSize *int, resolution *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp132 := strconv.Itoa(*end)
		values.Set("end", tmp132)
	}
	if internal != nil {
		tmp133 := strconv.FormatBool(*internal)
		values.Set("internal", tmp133)
	}
	if interval != nil {
		tmp134 := strconv.Itoa(*interval)
		values.Set("interval", tmp134)
	}
	if page != nil {
		tmp135 := strconv.Itoa(*page)
		values.Set("page", tmp135)
	}
	if pageSize != nil {
		tmp136 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp136)
	}
	if resolution != nil {
		tmp137 := strconv.Itoa(*resolution)
		values.Set("resolution", tmp137)
	}
	if start != nil {
		tmp138 := strconv.Itoa(*start)
		values.Set("start", tmp138)
	}
	u.RawQuery = values.Encode()
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
