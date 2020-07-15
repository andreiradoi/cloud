// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": Application Media Types
//
// Command:
// $ main

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// AdminDeletePayload media type (default view)
//
// Identifier: application/vnd.app.admin.user.delete+json; view=default
type AdminDeletePayload struct {
	Email    string `form:"email" json:"email" yaml:"email" xml:"email"`
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate validates the AdminDeletePayload media type instance.
func (mt *AdminDeletePayload) Validate() (err error) {
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	return
}

// DeviceDataRecordsResponse media type (default view)
//
// Identifier: application/vnd.app.device.data+json; view=default
type DeviceDataRecordsResponse struct {
	Data DeviceDataRecordCollection `form:"data" json:"data" yaml:"data" xml:"data"`
	Meta DeviceMetaRecordCollection `form:"meta" json:"meta" yaml:"meta" xml:"meta"`
}

// Validate validates the DeviceDataRecordsResponse media type instance.
func (mt *DeviceDataRecordsResponse) Validate() (err error) {
	if mt.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "meta"))
	}
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if err2 := mt.Data.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Meta.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DeviceDataRecord media type (default view)
//
// Identifier: application/vnd.app.device.data.record+json; view=default
type DeviceDataRecord struct {
	Data     map[string]interface{} `form:"data" json:"data" yaml:"data" xml:"data"`
	ID       int                    `form:"id" json:"id" yaml:"id" xml:"id"`
	Location []float64              `form:"location" json:"location" yaml:"location" xml:"location"`
	Meta     int                    `form:"meta" json:"meta" yaml:"meta" xml:"meta"`
	Record   int                    `form:"record" json:"record" yaml:"record" xml:"record"`
	Time     time.Time              `form:"time" json:"time" yaml:"time" xml:"time"`
}

// Validate validates the DeviceDataRecord media type instance.
func (mt *DeviceDataRecord) Validate() (err error) {

	if mt.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	return
}

// DeviceDataRecordCollection is the media type for an array of DeviceDataRecord (default view)
//
// Identifier: application/vnd.app.device.data.record+json; type=collection; view=default
type DeviceDataRecordCollection []*DeviceDataRecord

// Validate validates the DeviceDataRecordCollection media type instance.
func (mt DeviceDataRecordCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DeviceDataSummary media type (default view)
//
// Identifier: application/vnd.app.device.data.summary+json; view=default
type DeviceDataSummary struct {
	First int `form:"first" json:"first" yaml:"first" xml:"first"`
	Last  int `form:"last" json:"last" yaml:"last" xml:"last"`
	Size  int `form:"size" json:"size" yaml:"size" xml:"size"`
}

// Validate validates the DeviceDataSummary media type instance.
func (mt *DeviceDataSummary) Validate() (err error) {

	return
}

// JSONDataResponse media type (default view)
//
// Identifier: application/vnd.app.device.json.data+json; view=default
type JSONDataResponse struct {
	Versions []*JSONDataVersion `form:"versions" json:"versions" yaml:"versions" xml:"versions"`
}

// Validate validates the JSONDataResponse media type instance.
func (mt *JSONDataResponse) Validate() (err error) {
	if mt.Versions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "versions"))
	}
	for _, e := range mt.Versions {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// JSONDataSummaryResponse media type (default view)
//
// Identifier: application/vnd.app.device.json.data.summary+json; view=default
type JSONDataSummaryResponse struct {
	Data       []*JSONDataRow        `form:"data" json:"data" yaml:"data" xml:"data"`
	Modules    []*JSONDataMetaModule `form:"modules" json:"modules" yaml:"modules" xml:"modules"`
	Statistics *JSONDataStatistics   `form:"statistics" json:"statistics" yaml:"statistics" xml:"statistics"`
}

// Validate validates the JSONDataSummaryResponse media type instance.
func (mt *JSONDataSummaryResponse) Validate() (err error) {
	if mt.Modules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "modules"))
	}
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Statistics == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "statistics"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range mt.Modules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DeviceMetaRecord media type (default view)
//
// Identifier: application/vnd.app.device.meta.record+json; view=default
type DeviceMetaRecord struct {
	Data   map[string]interface{} `form:"data" json:"data" yaml:"data" xml:"data"`
	ID     int                    `form:"id" json:"id" yaml:"id" xml:"id"`
	Record int                    `form:"record" json:"record" yaml:"record" xml:"record"`
	Time   time.Time              `form:"time" json:"time" yaml:"time" xml:"time"`
}

// Validate validates the DeviceMetaRecord media type instance.
func (mt *DeviceMetaRecord) Validate() (err error) {

	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	return
}

// DeviceMetaRecordCollection is the media type for an array of DeviceMetaRecord (default view)
//
// Identifier: application/vnd.app.device.meta.record+json; type=collection; view=default
type DeviceMetaRecordCollection []*DeviceMetaRecord

// Validate validates the DeviceMetaRecordCollection media type instance.
func (mt DeviceMetaRecordCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DeviceMetaSummary media type (default view)
//
// Identifier: application/vnd.app.device.meta.summary+json; view=default
type DeviceMetaSummary struct {
	First int `form:"first" json:"first" yaml:"first" xml:"first"`
	Last  int `form:"last" json:"last" yaml:"last" xml:"last"`
	Size  int `form:"size" json:"size" yaml:"size" xml:"size"`
}

// Validate validates the DeviceMetaSummary media type instance.
func (mt *DeviceMetaSummary) Validate() (err error) {

	return
}

// DeviceProvisionSummary media type (default view)
//
// Identifier: application/vnd.app.device.provision.summary+json; view=default
type DeviceProvisionSummary struct {
	Created    time.Time          `form:"created" json:"created" yaml:"created" xml:"created"`
	Data       *DeviceDataSummary `form:"data" json:"data" yaml:"data" xml:"data"`
	Generation string             `form:"generation" json:"generation" yaml:"generation" xml:"generation"`
	Meta       *DeviceMetaSummary `form:"meta" json:"meta" yaml:"meta" xml:"meta"`
	Updated    time.Time          `form:"updated" json:"updated" yaml:"updated" xml:"updated"`
}

// Validate validates the DeviceProvisionSummary media type instance.
func (mt *DeviceProvisionSummary) Validate() (err error) {
	if mt.Generation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "generation"))
	}

	if mt.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "meta"))
	}
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	return
}

// DeviceProvisionSummaryCollection is the media type for an array of DeviceProvisionSummary (default view)
//
// Identifier: application/vnd.app.device.provision.summary+json; type=collection; view=default
type DeviceProvisionSummaryCollection []*DeviceProvisionSummary

// Validate validates the DeviceProvisionSummaryCollection media type instance.
func (mt DeviceProvisionSummaryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DeviceDataSummaryResponse media type (default view)
//
// Identifier: application/vnd.app.device.summary+json; view=default
type DeviceDataSummaryResponse struct {
	Provisions DeviceProvisionSummaryCollection `form:"provisions" json:"provisions" yaml:"provisions" xml:"provisions"`
}

// Validate validates the DeviceDataSummaryResponse media type instance.
func (mt *DeviceDataSummaryResponse) Validate() (err error) {
	if mt.Provisions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "provisions"))
	}
	if err2 := mt.Provisions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// FieldNoteMedia media type (default view)
//
// Identifier: application/vnd.app.field_note_media+json; view=default
type FieldNoteMedia struct {
	ContentType string    `form:"content_type" json:"content_type" yaml:"content_type" xml:"content_type"`
	Created     time.Time `form:"created" json:"created" yaml:"created" xml:"created"`
	ID          int       `form:"id" json:"id" yaml:"id" xml:"id"`
	URL         string    `form:"url" json:"url" yaml:"url" xml:"url"`
	UserID      int       `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the FieldNoteMedia media type instance.
func (mt *FieldNoteMedia) Validate() (err error) {

	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.ContentType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "content_type"))
	}
	return
}

// FieldNoteQueryResult media type (default view)
//
// Identifier: application/vnd.app.field_note_result+json; view=default
type FieldNoteQueryResult struct {
	CategoryKey      string    `form:"category_key" json:"category_key" yaml:"category_key" xml:"category_key"`
	Created          time.Time `form:"created" json:"created" yaml:"created" xml:"created"`
	Creator          string    `form:"creator" json:"creator" yaml:"creator" xml:"creator"`
	ID               int       `form:"id" json:"id" yaml:"id" xml:"id"`
	MediaContentType *string   `form:"media_content_type,omitempty" json:"media_content_type,omitempty" yaml:"media_content_type,omitempty" xml:"media_content_type,omitempty"`
	MediaID          *int      `form:"media_id,omitempty" json:"media_id,omitempty" yaml:"media_id,omitempty" xml:"media_id,omitempty"`
	MediaURL         *string   `form:"media_url,omitempty" json:"media_url,omitempty" yaml:"media_url,omitempty" xml:"media_url,omitempty"`
	Note             *string   `form:"note,omitempty" json:"note,omitempty" yaml:"note,omitempty" xml:"note,omitempty"`
	UserID           int       `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the FieldNoteQueryResult media type instance.
func (mt *FieldNoteQueryResult) Validate() (err error) {

	if mt.CategoryKey == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "category_key"))
	}
	if mt.Creator == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "creator"))
	}
	return
}

// FieldNoteQueryResultCollection is the media type for an array of FieldNoteQueryResult (default view)
//
// Identifier: application/vnd.app.field_note_result+json; type=collection; view=default
type FieldNoteQueryResultCollection []*FieldNoteQueryResult

// Validate validates the FieldNoteQueryResultCollection media type instance.
func (mt FieldNoteQueryResultCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// FieldNotes media type (default view)
//
// Identifier: application/vnd.app.field_notes+json; view=default
type FieldNotes struct {
	Notes FieldNoteQueryResultCollection `form:"notes" json:"notes" yaml:"notes" xml:"notes"`
}

// Validate validates the FieldNotes media type instance.
func (mt *FieldNotes) Validate() (err error) {
	if mt.Notes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "notes"))
	}
	if err2 := mt.Notes.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// FirmwareSummary media type (default view)
//
// Identifier: application/vnd.app.firmware+json; view=default
type FirmwareSummary struct {
	BuildNumber int                    `form:"buildNumber" json:"buildNumber" yaml:"buildNumber" xml:"buildNumber"`
	BuildTime   int                    `form:"buildTime" json:"buildTime" yaml:"buildTime" xml:"buildTime"`
	Etag        string                 `form:"etag" json:"etag" yaml:"etag" xml:"etag"`
	ID          int                    `form:"id" json:"id" yaml:"id" xml:"id"`
	Meta        map[string]interface{} `form:"meta" json:"meta" yaml:"meta" xml:"meta"`
	Module      string                 `form:"module" json:"module" yaml:"module" xml:"module"`
	Profile     string                 `form:"profile" json:"profile" yaml:"profile" xml:"profile"`
	Time        time.Time              `form:"time" json:"time" yaml:"time" xml:"time"`
	URL         string                 `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the FirmwareSummary media type instance.
func (mt *FirmwareSummary) Validate() (err error) {

	if mt.Etag == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "etag"))
	}
	if mt.Module == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "module"))
	}
	if mt.Profile == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "profile"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "meta"))
	}

	return
}

// FirmwareSummaryCollection is the media type for an array of FirmwareSummary (default view)
//
// Identifier: application/vnd.app.firmware+json; type=collection; view=default
type FirmwareSummaryCollection []*FirmwareSummary

// Validate validates the FirmwareSummaryCollection media type instance.
func (mt FirmwareSummaryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Firmwares media type (default view)
//
// Identifier: application/vnd.app.firmwares+json; view=default
type Firmwares struct {
	Firmwares FirmwareSummaryCollection `form:"firmwares" json:"firmwares" yaml:"firmwares" xml:"firmwares"`
}

// Validate validates the Firmwares media type instance.
func (mt *Firmwares) Validate() (err error) {
	if mt.Firmwares == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firmwares"))
	}
	if err2 := mt.Firmwares.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// Location media type (default view)
//
// Identifier: application/vnd.app.location+json; view=default
type Location struct {
	Location string `form:"location" json:"location" yaml:"location" xml:"location"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Location); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.location`, mt.Location, goa.FormatURI, err2))
	}
	return
}

// MediaReferenceResponse media type (default view)
//
// Identifier: application/vnd.app.media+json; view=default
type MediaReferenceResponse struct {
	ID  int    `form:"id" json:"id" yaml:"id" xml:"id"`
	URL string `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the MediaReferenceResponse media type instance.
func (mt *MediaReferenceResponse) Validate() (err error) {

	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	return
}

// Project media type (default view)
//
// Identifier: application/vnd.app.project+json; view=default
type Project struct {
	Description       string  `form:"description" json:"description" yaml:"description" xml:"description"`
	EndTime           *string `form:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty" xml:"endTime,omitempty"`
	Goal              string  `form:"goal" json:"goal" yaml:"goal" xml:"goal"`
	ID                int     `form:"id" json:"id" yaml:"id" xml:"id"`
	Location          string  `form:"location" json:"location" yaml:"location" xml:"location"`
	Name              string  `form:"name" json:"name" yaml:"name" xml:"name"`
	NumberOfFollowers int     `form:"numberOfFollowers" json:"numberOfFollowers" yaml:"numberOfFollowers" xml:"numberOfFollowers"`
	Photo             *string `form:"photo,omitempty" json:"photo,omitempty" yaml:"photo,omitempty" xml:"photo,omitempty"`
	Private           bool    `form:"private" json:"private" yaml:"private" xml:"private"`
	ReadOnly          bool    `form:"readOnly" json:"readOnly" yaml:"readOnly" xml:"readOnly"`
	Slug              string  `form:"slug" json:"slug" yaml:"slug" xml:"slug"`
	StartTime         *string `form:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty" xml:"startTime,omitempty"`
	Tags              string  `form:"tags" json:"tags" yaml:"tags" xml:"tags"`
}

// Validate validates the Project media type instance.
func (mt *Project) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.Goal == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "goal"))
	}
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if mt.Tags == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tags"))
	}

	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// ProjectRole media type (default view)
//
// Identifier: application/vnd.app.project.role+json; view=default
type ProjectRole struct {
	ID   int    `form:"id" json:"id" yaml:"id" xml:"id"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the ProjectRole media type instance.
func (mt *ProjectRole) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// ProjectRoleCollection is the media type for an array of ProjectRole (default view)
//
// Identifier: application/vnd.app.project.role+json; type=collection; view=default
type ProjectRoleCollection []*ProjectRole

// Validate validates the ProjectRoleCollection media type instance.
func (mt ProjectRoleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ProjectUser media type (default view)
//
// Identifier: application/vnd.app.project.user+json; view=default
type ProjectUser struct {
	Membership string `form:"membership" json:"membership" yaml:"membership" xml:"membership"`
	Role       string `form:"role" json:"role" yaml:"role" xml:"role"`
	User       *User  `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the ProjectUser media type instance.
func (mt *ProjectUser) Validate() (err error) {
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	if mt.Membership == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "membership"))
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ProjectUserCollection is the media type for an array of ProjectUser (default view)
//
// Identifier: application/vnd.app.project.user+json; type=collection; view=default
type ProjectUserCollection []*ProjectUser

// Validate validates the ProjectUserCollection media type instance.
func (mt ProjectUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ProjectCollection is the media type for an array of Project (default view)
//
// Identifier: application/vnd.app.project+json; type=collection; view=default
type ProjectCollection []*Project

// Validate validates the ProjectCollection media type instance.
func (mt ProjectCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Projects media type (default view)
//
// Identifier: application/vnd.app.projects+json; view=default
type Projects struct {
	Projects ProjectCollection `form:"projects" json:"projects" yaml:"projects" xml:"projects"`
}

// Validate validates the Projects media type instance.
func (mt *Projects) Validate() (err error) {
	if mt.Projects == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "projects"))
	}
	if err2 := mt.Projects.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// User media type (default view)
//
// Identifier: application/vnd.app.user+json; view=default
type User struct {
	Admin bool       `form:"admin" json:"admin" yaml:"admin" xml:"admin"`
	Bio   string     `form:"bio" json:"bio" yaml:"bio" xml:"bio"`
	Email string     `form:"email" json:"email" yaml:"email" xml:"email"`
	ID    int        `form:"id" json:"id" yaml:"id" xml:"id"`
	Name  string     `form:"name" json:"name" yaml:"name" xml:"name"`
	Photo *UserPhoto `form:"photo,omitempty" json:"photo,omitempty" yaml:"photo,omitempty" xml:"photo,omitempty"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, mt.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, mt.Name, `\S`))
	}
	if utf8.RuneCountInString(mt.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 256, false))
	}
	return
}

// TransmissionToken media type (default view)
//
// Identifier: application/vnd.app.user.transmission.token+json; view=default
type TransmissionToken struct {
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the TransmissionToken media type instance.
func (mt *TransmissionToken) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// ProjectUsers media type (default view)
//
// Identifier: application/vnd.app.users+json; view=default
type ProjectUsers struct {
	Users ProjectUserCollection `form:"users" json:"users" yaml:"users" xml:"users"`
}

// Validate validates the ProjectUsers media type instance.
func (mt *ProjectUsers) Validate() (err error) {
	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}
