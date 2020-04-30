package api

import (
	"encoding/hex"
	"fmt"
	"strings"

	"image"

	"github.com/goadesign/goa"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend/repositories"
	"github.com/fieldkit/cloud/server/data"
)

func StationType(p Permissions, station *data.Station, owner *data.User, ingestions []*data.Ingestion, noteMedia []*data.FieldNoteMediaForStation) (*app.Station, error) {
	status, err := station.GetStatus()
	if err != nil {
		return nil, err
	}

	images := make([]*app.ImageRef, 0)
	for _, row := range noteMedia {
		if row.URL != "" && strings.Contains(row.ContentType, "image") {
			images = append(images, &app.ImageRef{
				URL: fmt.Sprintf("/stations/%d/field-note-media/%d", station.ID, row.ID),
			})
		}
	}

	lastUploads := make([]*app.LastUpload, len(ingestions))
	for i, ingestion := range ingestions {
		lastUploads[i] = &app.LastUpload{
			ID:       int(ingestion.ID),
			Time:     ingestion.Time,
			UploadID: ingestion.UploadID,
			Size:     int(ingestion.Size),
			Type:     ingestion.Type,
			URL:      ingestion.URL,
			Blocks:   ingestion.Blocks.ToIntArray(),
		}
	}

	sp, err := p.ForStation(station)
	if err != nil {
		return nil, err
	}

	return &app.Station{
		ID: int(station.ID),
		Owner: &app.StationOwner{
			ID:   int(owner.ID),
			Name: owner.Name,
		},
		DeviceID:    hex.EncodeToString(station.DeviceID),
		Name:        station.Name,
		LastUploads: lastUploads,
		StatusJSON:  status,
		Images:      images,
		ReadOnly:    sp.IsReadOnly(),
		Photos: &app.StationPhotos{
			Small: fmt.Sprintf("/stations/%d/photo", station.ID),
		},
	}, nil
}

func sortByStation(ingestions []*data.Ingestion) map[string][]*data.Ingestion {
	m := make(map[string][]*data.Ingestion)
	for _, i := range ingestions {
		key := hex.EncodeToString(i.DeviceID)
		if m[key] == nil {
			m[key] = make([]*data.Ingestion, 0)
		}
		m[key] = append(m[key], i)
	}
	return m
}

func sortMediaByStation(all []*data.FieldNoteMediaForStation) map[int32][]*data.FieldNoteMediaForStation {
	m := make(map[int32][]*data.FieldNoteMediaForStation)
	for _, row := range all {
		if m[row.StationID] == nil {
			m[row.StationID] = make([]*data.FieldNoteMediaForStation, 0)
		}
		m[row.StationID] = append(m[row.StationID], row)
	}
	return m
}

func sortUsersByID(all []*data.User) map[int32]*data.User {
	m := make(map[int32]*data.User)
	for _, row := range all {
		m[row.ID] = row
	}
	return m
}

func StationsType(p Permissions, stations []*data.Station, owners []*data.User, ingestions []*data.Ingestion, noteMedia []*data.FieldNoteMediaForStation) (*app.Stations, error) {
	stationsCollection := make([]*app.Station, len(stations))

	ownersByID := sortUsersByID(owners)
	noteMediaByStation := sortMediaByStation(noteMedia)
	ingestionsByStation := sortByStation(ingestions)
	for i, station := range stations {
		key := hex.EncodeToString(station.DeviceID)
		ingestionsByStation := ingestionsByStation[key]
		if ingestionsByStation == nil {
			ingestionsByStation = make([]*data.Ingestion, 0)
		}
		noteMediaByStation := noteMediaByStation[station.ID]
		if noteMediaByStation == nil {
			noteMediaByStation = make([]*data.FieldNoteMediaForStation, 0)
		}
		owner := ownersByID[station.OwnerID]
		appStation, err := StationType(p, station, owner, ingestionsByStation, noteMediaByStation)
		if err != nil {
			return nil, err
		}
		stationsCollection[i] = appStation
	}

	return &app.Stations{
		Stations: stationsCollection,
	}, nil
}

type StationController struct {
	*goa.Controller
	options *ControllerOptions
}

func NewStationController(service *goa.Service, options *ControllerOptions) *StationController {
	return &StationController{
		Controller: service.NewController("StationController"),
		options:    options,
	}
}

func (c *StationController) ListProject(ctx *app.ListProjectStationContext) error {
	p, err := NewPermissions(ctx, c.options).Unwrap()
	if err != nil {
		return err
	}

	stations := []*data.Station{}
	if err := c.options.Database.SelectContext(ctx, &stations, `
		SELECT * FROM fieldkit.station WHERE id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1)`, ctx.ProjectID); err != nil {
		return err
	}

	owners := []*data.User{}
	if err := c.options.Database.SelectContext(ctx, &owners, `
		SELECT * FROM fieldkit.user WHERE id IN (SELECT owner_id FROM fieldkit.station
		WHERE id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1))`, ctx.ProjectID); err != nil {
		return err
	}

	ingestions := []*data.Ingestion{}
	if err := c.options.Database.SelectContext(ctx, &ingestions, `
		SELECT * FROM fieldkit.ingestion WHERE device_id IN (SELECT s.device_id FROM fieldkit.station AS s JOIN fieldkit.project_station AS ps ON (s.id = ps.station_id)
		WHERE project_id = $1) ORDER BY time DESC`, ctx.ProjectID); err != nil {
		return err
	}

	noteMedia := []*data.FieldNoteMediaForStation{}
	if err := c.options.Database.SelectContext(ctx, &noteMedia, `
		SELECT s.id AS station_id, fnm.* FROM fieldkit.station AS s JOIN fieldkit.field_note AS fn ON (fn.station_id = s.id) JOIN fieldkit.field_note_media AS fnm ON (fn.media_id = fnm.id)
		WHERE s.id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1) ORDER BY fnm.created DESC`, ctx.ProjectID); err != nil {
		return err
	}

	stationsWm, err := StationsType(p, stations, owners, ingestions, noteMedia)
	if err != nil {
		return err
	}

	return ctx.OK(stationsWm)
}

func (c *StationController) List(ctx *app.ListStationContext) error {
	p, err := NewPermissions(ctx, c.options).Unwrap()
	if err != nil {
		return err
	}

	stations := []*data.Station{}
	if err := c.options.Database.SelectContext(ctx, &stations, `SELECT * FROM fieldkit.station WHERE owner_id = $1`, p.UserID()); err != nil {
		return err
	}

	owners := []*data.User{}
	if err := c.options.Database.SelectContext(ctx, &owners, `SELECT * FROM fieldkit.user WHERE id = $1`, p.UserID()); err != nil {
		return err
	}

	ingestions := []*data.Ingestion{}
	if err := c.options.Database.SelectContext(ctx, &ingestions, `
		SELECT * FROM fieldkit.ingestion WHERE device_id IN (SELECT device_id FROM fieldkit.station WHERE owner_id = $1) ORDER BY time DESC`, p.UserID()); err != nil {
		return err
	}

	noteMedia := []*data.FieldNoteMediaForStation{}
	if err := c.options.Database.SelectContext(ctx, &noteMedia, `
		SELECT s.id AS station_id, fnm.* FROM fieldkit.station AS s JOIN fieldkit.field_note AS fn ON (fn.station_id = s.id) JOIN fieldkit.field_note_media AS fnm ON (fn.media_id = fnm.id)
		WHERE s.owner_id = $1 ORDER BY fnm.created DESC`, p.UserID()); err != nil {
		return err
	}

	stationsWm, err := StationsType(p, stations, owners, ingestions, noteMedia)
	if err != nil {
		return err
	}

	return ctx.OK(stationsWm)
}

func (c *StationController) Photo(ctx *app.PhotoStationContext) error {
	x := uint(124)
	y := uint(100)

	defaultPhotoContentType := "image/png"
	defaultPhoto, err := StationDefaultPicture(int64(ctx.StationID))
	if err != nil {
		// NOTE This, hopefully never happens because we've got no image to send back.
		return err
	}

	allMedia := []*data.FieldNoteMediaForStation{}
	if err := c.options.Database.SelectContext(ctx, &allMedia, `
		SELECT s.id AS station_id, fnm.* FROM fieldkit.station AS s JOIN fieldkit.field_note AS fn ON (fn.station_id = s.id) JOIN fieldkit.field_note_media AS fnm ON (fn.media_id = fnm.id)
		WHERE s.id = $1 ORDER BY fnm.created DESC`, ctx.StationID); err != nil {
		return LogErrorAndSendData(ctx, ctx.ResponseData, err, defaultPhotoContentType, defaultPhoto)
	}

	if len(allMedia) == 0 {
		return SendData(ctx.ResponseData, defaultPhotoContentType, defaultPhoto)
	}

	mr := repositories.NewMediaRepository(c.options.Session, c.options.Buckets.Media)

	lm, err := mr.LoadByURL(ctx, allMedia[0].URL)
	if err != nil {
		return LogErrorAndSendData(ctx, ctx.ResponseData, err, defaultPhotoContentType, defaultPhoto)
	}

	original, _, err := image.Decode(lm.Reader)
	if err != nil {
		return LogErrorAndSendData(ctx, ctx.ResponseData, err, defaultPhotoContentType, defaultPhoto)
	}

	cropped, err := SmartCrop(original, x, y)
	if err != nil {
		return LogErrorAndSendData(ctx, ctx.ResponseData, err, defaultPhotoContentType, defaultPhoto)
	}

	err = SendImage(ctx.ResponseData, cropped)
	if err != nil {
		return LogErrorAndSendData(ctx, ctx.ResponseData, err, defaultPhotoContentType, defaultPhoto)
	}

	return nil
}
