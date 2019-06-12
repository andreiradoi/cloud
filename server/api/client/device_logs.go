// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": device_logs Resource Client
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

// AllDeviceLogsPath computes a request path to the all action of device_logs.
func AllDeviceLogsPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/logs", param0)
}

// Get all device logs
func (c *Client) AllDeviceLogs(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAllDeviceLogsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAllDeviceLogsRequest create the request corresponding to the all action endpoint of the device_logs resource.
func (c *Client) NewAllDeviceLogsRequest(ctx context.Context, path string) (*http.Request, error) {
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
