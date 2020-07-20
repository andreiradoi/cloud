// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	goa "goa.design/goa/v3/pkg"
)

// BuildUpdatePayload builds the payload for the notes update endpoint from CLI
// flags.
func BuildUpdatePayload(notesUpdateBody string, notesUpdateStationID string, notesUpdateAuth string) (*notes.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(notesUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"notes\": {\n         \"creating\": [\n            {\n               \"body\": \"Fugiat voluptatem facere.\",\n               \"key\": \"Quia nobis libero quia.\",\n               \"mediaIds\": [\n                  2039732461194112458,\n                  7816285458348604520,\n                  7577586275964425592,\n                  948553046914617935\n               ]\n            },\n            {\n               \"body\": \"Fugiat voluptatem facere.\",\n               \"key\": \"Quia nobis libero quia.\",\n               \"mediaIds\": [\n                  2039732461194112458,\n                  7816285458348604520,\n                  7577586275964425592,\n                  948553046914617935\n               ]\n            },\n            {\n               \"body\": \"Fugiat voluptatem facere.\",\n               \"key\": \"Quia nobis libero quia.\",\n               \"mediaIds\": [\n                  2039732461194112458,\n                  7816285458348604520,\n                  7577586275964425592,\n                  948553046914617935\n               ]\n            },\n            {\n               \"body\": \"Fugiat voluptatem facere.\",\n               \"key\": \"Quia nobis libero quia.\",\n               \"mediaIds\": [\n                  2039732461194112458,\n                  7816285458348604520,\n                  7577586275964425592,\n                  948553046914617935\n               ]\n            }\n         ],\n         \"notes\": [\n            {\n               \"body\": \"Quisquam ea quidem a asperiores consequatur architecto.\",\n               \"id\": 6709774753631418749,\n               \"key\": \"Iusto est reiciendis.\",\n               \"mediaIds\": [\n                  6433030965608229216,\n                  5445315585411334221,\n                  7103160155254992702\n               ]\n            },\n            {\n               \"body\": \"Quisquam ea quidem a asperiores consequatur architecto.\",\n               \"id\": 6709774753631418749,\n               \"key\": \"Iusto est reiciendis.\",\n               \"mediaIds\": [\n                  6433030965608229216,\n                  5445315585411334221,\n                  7103160155254992702\n               ]\n            },\n            {\n               \"body\": \"Quisquam ea quidem a asperiores consequatur architecto.\",\n               \"id\": 6709774753631418749,\n               \"key\": \"Iusto est reiciendis.\",\n               \"mediaIds\": [\n                  6433030965608229216,\n                  5445315585411334221,\n                  7103160155254992702\n               ]\n            }\n         ]\n      }\n   }'")
		}
		if body.Notes == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("notes", "body"))
		}
		if body.Notes != nil {
			if err2 := ValidateFieldNoteUpdateRequestBody(body.Notes); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesUpdateStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesUpdateAuth
	}
	v := &notes.UpdatePayload{}
	if body.Notes != nil {
		v.Notes = marshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate(body.Notes)
	}
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildGetPayload builds the payload for the notes get endpoint from CLI flags.
func BuildGetPayload(notesGetStationID string, notesGetAuth string) (*notes.GetPayload, error) {
	var err error
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesGetStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesGetAuth
	}
	v := &notes.GetPayload{}
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildMediaPayload builds the payload for the notes media endpoint from CLI
// flags.
func BuildMediaPayload(notesMediaMediaID string, notesMediaAuth string) (*notes.MediaPayload, error) {
	var err error
	var mediaID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesMediaMediaID, 10, 32)
		mediaID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for mediaID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesMediaAuth
	}
	v := &notes.MediaPayload{}
	v.MediaID = mediaID
	v.Auth = auth

	return v, nil
}

// BuildUploadPayload builds the payload for the notes upload endpoint from CLI
// flags.
func BuildUploadPayload(notesUploadStationID string, notesUploadKey string, notesUploadContentType string, notesUploadContentLength string, notesUploadAuth string) (*notes.UploadPayload, error) {
	var err error
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesUploadStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var key string
	{
		key = notesUploadKey
	}
	var contentType string
	{
		contentType = notesUploadContentType
	}
	var contentLength int64
	{
		contentLength, err = strconv.ParseInt(notesUploadContentLength, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for contentLength, must be INT64")
		}
	}
	var auth string
	{
		auth = notesUploadAuth
	}
	v := &notes.UploadPayload{}
	v.StationID = stationID
	v.Key = key
	v.ContentType = contentType
	v.ContentLength = contentLength
	v.Auth = auth

	return v, nil
}
