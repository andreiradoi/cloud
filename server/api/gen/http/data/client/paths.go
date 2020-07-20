// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the data service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"fmt"
)

// DeviceSummaryDataPath returns the URL path to the data service device summary HTTP endpoint.
func DeviceSummaryDataPath(deviceID string) string {
	return fmt.Sprintf("/data/devices/%v/summary", deviceID)
}
