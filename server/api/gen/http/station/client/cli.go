// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	station "github.com/fieldkit/cloud/server/api/gen/station"
)

// BuildAddPayload builds the payload for the station add endpoint from CLI
// flags.
func BuildAddPayload(stationAddBody string, stationAddAuth string) (*station.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(stationAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"deviceId\": \"Blanditiis adipisci illo.\",\n      \"locationName\": \"Et rerum iste.\",\n      \"name\": \"Cupiditate non.\",\n      \"statusPb\": \"Doloribus sed fuga expedita.\"\n   }'")
		}
	}
	var auth string
	{
		auth = stationAddAuth
	}
	v := &station.AddPayload{
		Name:         body.Name,
		DeviceID:     body.DeviceID,
		LocationName: body.LocationName,
		StatusPb:     body.StatusPb,
	}
	v.Auth = auth

	return v, nil
}

// BuildGetPayload builds the payload for the station get endpoint from CLI
// flags.
func BuildGetPayload(stationGetID string, stationGetAuth string) (*station.GetPayload, error) {
	var err error
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationGetID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationGetAuth
	}
	v := &station.GetPayload{}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildUpdatePayload builds the payload for the station update endpoint from
// CLI flags.
func BuildUpdatePayload(stationUpdateBody string, stationUpdateID string, stationUpdateAuth string) (*station.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(stationUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"locationName\": \"Dolores velit aperiam quod numquam.\",\n      \"name\": \"Quis rerum rerum illum.\",\n      \"statusPb\": \"Quae vitae aut.\"\n   }'")
		}
	}
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationUpdateID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationUpdateAuth
	}
	v := &station.UpdatePayload{
		Name:         body.Name,
		LocationName: body.LocationName,
		StatusPb:     body.StatusPb,
	}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildListMinePayload builds the payload for the station list mine endpoint
// from CLI flags.
func BuildListMinePayload(stationListMineAuth string) (*station.ListMinePayload, error) {
	var auth string
	{
		auth = stationListMineAuth
	}
	v := &station.ListMinePayload{}
	v.Auth = auth

	return v, nil
}

// BuildListProjectPayload builds the payload for the station list project
// endpoint from CLI flags.
func BuildListProjectPayload(stationListProjectID string, stationListProjectAuth string) (*station.ListProjectPayload, error) {
	var err error
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationListProjectID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth *string
	{
		if stationListProjectAuth != "" {
			auth = &stationListProjectAuth
		}
	}
	v := &station.ListProjectPayload{}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildDownloadPhotoPayload builds the payload for the station download photo
// endpoint from CLI flags.
func BuildDownloadPhotoPayload(stationDownloadPhotoStationID string, stationDownloadPhotoSize string, stationDownloadPhotoIfNoneMatch string, stationDownloadPhotoAuth string) (*station.DownloadPhotoPayload, error) {
	var err error
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(stationDownloadPhotoStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var size *int32
	{
		if stationDownloadPhotoSize != "" {
			var v int64
			v, err = strconv.ParseInt(stationDownloadPhotoSize, 10, 32)
			val := int32(v)
			size = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for size, must be INT32")
			}
		}
	}
	var ifNoneMatch *string
	{
		if stationDownloadPhotoIfNoneMatch != "" {
			ifNoneMatch = &stationDownloadPhotoIfNoneMatch
		}
	}
	var auth string
	{
		auth = stationDownloadPhotoAuth
	}
	v := &station.DownloadPhotoPayload{}
	v.StationID = stationID
	v.Size = size
	v.IfNoneMatch = ifNoneMatch
	v.Auth = auth

	return v, nil
}

// BuildListAllPayload builds the payload for the station list all endpoint
// from CLI flags.
func BuildListAllPayload(stationListAllPage string, stationListAllPageSize string, stationListAllOwnerID string, stationListAllQuery string, stationListAllSortBy string, stationListAllAuth string) (*station.ListAllPayload, error) {
	var err error
	var page *int32
	{
		if stationListAllPage != "" {
			var v int64
			v, err = strconv.ParseInt(stationListAllPage, 10, 32)
			val := int32(v)
			page = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for page, must be INT32")
			}
		}
	}
	var pageSize *int32
	{
		if stationListAllPageSize != "" {
			var v int64
			v, err = strconv.ParseInt(stationListAllPageSize, 10, 32)
			val := int32(v)
			pageSize = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageSize, must be INT32")
			}
		}
	}
	var ownerID *int32
	{
		if stationListAllOwnerID != "" {
			var v int64
			v, err = strconv.ParseInt(stationListAllOwnerID, 10, 32)
			val := int32(v)
			ownerID = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for ownerID, must be INT32")
			}
		}
	}
	var query *string
	{
		if stationListAllQuery != "" {
			query = &stationListAllQuery
		}
	}
	var sortBy *string
	{
		if stationListAllSortBy != "" {
			sortBy = &stationListAllSortBy
		}
	}
	var auth string
	{
		auth = stationListAllAuth
	}
	v := &station.ListAllPayload{}
	v.Page = page
	v.PageSize = pageSize
	v.OwnerID = ownerID
	v.Query = query
	v.SortBy = sortBy
	v.Auth = auth

	return v, nil
}

// BuildDeletePayload builds the payload for the station delete endpoint from
// CLI flags.
func BuildDeletePayload(stationDeleteStationID string, stationDeleteAuth string) (*station.DeletePayload, error) {
	var err error
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(stationDeleteStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = stationDeleteAuth
	}
	v := &station.DeletePayload{}
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}
