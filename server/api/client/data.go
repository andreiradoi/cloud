// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": data Resource Client
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

// DeleteDataPath computes a request path to the delete action of data.
func DeleteDataPath(ingestionID int) string {
	param0 := strconv.Itoa(ingestionID)

	return fmt.Sprintf("/data/ingestions/%s", param0)
}

// Delete data
func (c *Client) DeleteData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteDataRequest create the request corresponding to the delete action endpoint of the data resource.
func (c *Client) NewDeleteDataRequest(ctx context.Context, path string) (*http.Request, error) {
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

// DeviceDataDataPath computes a request path to the device data action of data.
func DeviceDataDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data", param0)
}

// Retrieve data
func (c *Client) DeviceDataData(ctx context.Context, path string, firstBlock *int, lastBlock *int, page *int, pageSize *int) (*http.Response, error) {
	req, err := c.NewDeviceDataDataRequest(ctx, path, firstBlock, lastBlock, page, pageSize)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeviceDataDataRequest create the request corresponding to the device data action endpoint of the data resource.
func (c *Client) NewDeviceDataDataRequest(ctx context.Context, path string, firstBlock *int, lastBlock *int, page *int, pageSize *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if firstBlock != nil {
		tmp140 := strconv.Itoa(*firstBlock)
		values.Set("firstBlock", tmp140)
	}
	if lastBlock != nil {
		tmp141 := strconv.Itoa(*lastBlock)
		values.Set("lastBlock", tmp141)
	}
	if page != nil {
		tmp142 := strconv.Itoa(*page)
		values.Set("page", tmp142)
	}
	if pageSize != nil {
		tmp143 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp143)
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

// DeviceSummaryDataPath computes a request path to the device summary action of data.
func DeviceSummaryDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/summary", param0)
}

// Retrieve summary
func (c *Client) DeviceSummaryData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeviceSummaryDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeviceSummaryDataRequest create the request corresponding to the device summary action endpoint of the data resource.
func (c *Client) NewDeviceSummaryDataRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ProcessDataPath computes a request path to the process action of data.
func ProcessDataPath() string {

	return fmt.Sprintf("/data/process")
}

// Process data
func (c *Client) ProcessData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewProcessDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewProcessDataRequest create the request corresponding to the process action endpoint of the data resource.
func (c *Client) NewProcessDataRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
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

// ProcessIngestionDataPath computes a request path to the process ingestion action of data.
func ProcessIngestionDataPath(ingestionID int) string {
	param0 := strconv.Itoa(ingestionID)

	return fmt.Sprintf("/data/ingestions/%s/process", param0)
}

// Process ingestion
func (c *Client) ProcessIngestionData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewProcessIngestionDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewProcessIngestionDataRequest create the request corresponding to the process ingestion action endpoint of the data resource.
func (c *Client) NewProcessIngestionDataRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
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
