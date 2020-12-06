// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the station service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// AddStationPath returns the URL path to the station service add HTTP endpoint.
func AddStationPath() string {
	return "/stations"
}

// GetStationPath returns the URL path to the station service get HTTP endpoint.
func GetStationPath(id int32) string {
	return fmt.Sprintf("/stations/%v", id)
}

// TransferStationPath returns the URL path to the station service transfer HTTP endpoint.
func TransferStationPath(id int32, ownerID int32) string {
	return fmt.Sprintf("/stations/%v/transfer/%v", id, ownerID)
}

// UpdateStationPath returns the URL path to the station service update HTTP endpoint.
func UpdateStationPath(id int32) string {
	return fmt.Sprintf("/stations/%v", id)
}

// ListMineStationPath returns the URL path to the station service list mine HTTP endpoint.
func ListMineStationPath() string {
	return "/user/stations"
}

// ListProjectStationPath returns the URL path to the station service list project HTTP endpoint.
func ListProjectStationPath(id int32) string {
	return fmt.Sprintf("/projects/%v/stations", id)
}

// DownloadPhotoStationPath returns the URL path to the station service download photo HTTP endpoint.
func DownloadPhotoStationPath(stationID int32) string {
	return fmt.Sprintf("/stations/%v/photo", stationID)
}

// ListAllStationPath returns the URL path to the station service list all HTTP endpoint.
func ListAllStationPath() string {
	return "/admin/stations"
}

// DeleteStationPath returns the URL path to the station service delete HTTP endpoint.
func DeleteStationPath(stationID int32) string {
	return fmt.Sprintf("/admin/stations/%v", stationID)
}

// AdminSearchStationPath returns the URL path to the station service admin search HTTP endpoint.
func AdminSearchStationPath() string {
	return "/admin/stations/search"
}
