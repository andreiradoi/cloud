// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package station

import (
	"context"

	stationviews "github.com/fieldkit/cloud/server/api/gen/station/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the station service interface.
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res *StationFull, err error)
	// Get implements get.
	Get(context.Context, *GetPayload) (res *StationFull, err error)
	// Update implements update.
	Update(context.Context, *UpdatePayload) (res *StationFull, err error)
	// ListMine implements list mine.
	ListMine(context.Context, *ListMinePayload) (res *StationsFull, err error)
	// ListProject implements list project.
	ListProject(context.Context, *ListProjectPayload) (res *StationsFull, err error)
	// DownloadPhoto implements download photo.
	DownloadPhoto(context.Context, *DownloadPhotoPayload) (res *DownloadedPhoto, err error)
	// ListAll implements list all.
	ListAll(context.Context, *ListAllPayload) (res *PageOfStations, err error)
	// Delete implements delete.
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "station"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [8]string{"add", "get", "update", "list mine", "list project", "download photo", "list all", "delete"}

// AddPayload is the payload type of the station service add method.
type AddPayload struct {
	Auth         string
	Name         string
	DeviceID     string
	LocationName *string
	StatusPb     *string
}

// StationFull is the result type of the station service add method.
type StationFull struct {
	ID                 int32
	Name               string
	Owner              *StationOwner
	DeviceID           string
	Uploads            []*StationUpload
	Photos             *StationPhotos
	ReadOnly           bool
	Battery            *float32
	RecordingStartedAt *int64
	MemoryUsed         *int32
	MemoryAvailable    *int32
	FirmwareNumber     *int32
	FirmwareTime       *int64
	Configurations     *StationConfigurations
	UpdatedAt          int64
	LocationName       *string
	PlaceNameOther     *string
	PlaceNameNative    *string
	Location           *StationLocation
	Data               *StationDataSummary
}

// GetPayload is the payload type of the station service get method.
type GetPayload struct {
	Auth string
	ID   int32
}

// UpdatePayload is the payload type of the station service update method.
type UpdatePayload struct {
	Auth         string
	ID           int32
	Name         string
	LocationName *string
	StatusPb     *string
}

// ListMinePayload is the payload type of the station service list mine method.
type ListMinePayload struct {
	Auth string
}

// StationsFull is the result type of the station service list mine method.
type StationsFull struct {
	Stations StationFullCollection
}

// ListProjectPayload is the payload type of the station service list project
// method.
type ListProjectPayload struct {
	Auth *string
	ID   int32
}

// DownloadPhotoPayload is the payload type of the station service download
// photo method.
type DownloadPhotoPayload struct {
	Auth        string
	StationID   int32
	Size        *int32
	IfNoneMatch *string
}

// DownloadedPhoto is the result type of the station service download photo
// method.
type DownloadedPhoto struct {
	Length      int64
	ContentType string
	Etag        string
	Body        []byte
}

// ListAllPayload is the payload type of the station service list all method.
type ListAllPayload struct {
	Auth     string
	Page     *int32
	PageSize *int32
	OwnerID  *int32
	Query    *string
	SortBy   *string
}

// PageOfStations is the result type of the station service list all method.
type PageOfStations struct {
	Stations []*EssentialStation
	Total    int32
}

// DeletePayload is the payload type of the station service delete method.
type DeletePayload struct {
	Auth      string
	StationID int32
}

type StationOwner struct {
	ID   int32
	Name string
}

type StationUpload struct {
	ID       int64
	Time     int64
	UploadID string
	Size     int64
	URL      string
	Type     string
	Blocks   []int64
}

type StationPhotos struct {
	Small string
}

type StationConfigurations struct {
	All []*StationConfiguration
}

type StationConfiguration struct {
	ID           int64
	Time         int64
	ProvisionID  int64
	MetaRecordID *int64
	SourceID     *int32
	Modules      []*StationModule
}

type StationModule struct {
	ID           int64
	HardwareID   *string
	MetaRecordID *int64
	Name         string
	Position     int32
	Flags        int32
	Internal     bool
	FullKey      string
	Sensors      []*StationSensor
}

type StationSensor struct {
	Name          string
	UnitOfMeasure string
	Reading       *SensorReading
	Key           string
	FullKey       string
	Ranges        []*SensorRange
}

type SensorReading struct {
	Last float32
	Time int64
}

type SensorRange struct {
	Minimum float32
	Maximum float32
}

type StationLocation struct {
	Precise []float64
	Regions []*StationRegion
	URL     *string
}

type StationRegion struct {
	Name  string
	Shape [][][]float64
}

type StationDataSummary struct {
	Start           int64
	End             int64
	NumberOfSamples int64
}

type StationFullCollection []*StationFull

type EssentialStation struct {
	ID                 int64
	DeviceID           string
	Name               string
	Owner              *StationOwner
	CreatedAt          int64
	UpdatedAt          int64
	RecordingStartedAt *int64
	MemoryUsed         *int32
	MemoryAvailable    *int32
	FirmwareNumber     *int32
	FirmwareTime       *int64
	Location           *StationLocation
	LastIngestionAt    *int64
}

// MakeStationOwnerConflict builds a goa.ServiceError from an error.
func MakeStationOwnerConflict(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "station-owner-conflict",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad-request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewStationFull initializes result type StationFull from viewed result type
// StationFull.
func NewStationFull(vres *stationviews.StationFull) *StationFull {
	return newStationFull(vres.Projected)
}

// NewViewedStationFull initializes viewed result type StationFull from result
// type StationFull using the given view.
func NewViewedStationFull(res *StationFull, view string) *stationviews.StationFull {
	p := newStationFullView(res)
	return &stationviews.StationFull{Projected: p, View: "default"}
}

// NewStationsFull initializes result type StationsFull from viewed result type
// StationsFull.
func NewStationsFull(vres *stationviews.StationsFull) *StationsFull {
	return newStationsFull(vres.Projected)
}

// NewViewedStationsFull initializes viewed result type StationsFull from
// result type StationsFull using the given view.
func NewViewedStationsFull(res *StationsFull, view string) *stationviews.StationsFull {
	p := newStationsFullView(res)
	return &stationviews.StationsFull{Projected: p, View: "default"}
}

// NewDownloadedPhoto initializes result type DownloadedPhoto from viewed
// result type DownloadedPhoto.
func NewDownloadedPhoto(vres *stationviews.DownloadedPhoto) *DownloadedPhoto {
	return newDownloadedPhoto(vres.Projected)
}

// NewViewedDownloadedPhoto initializes viewed result type DownloadedPhoto from
// result type DownloadedPhoto using the given view.
func NewViewedDownloadedPhoto(res *DownloadedPhoto, view string) *stationviews.DownloadedPhoto {
	p := newDownloadedPhotoView(res)
	return &stationviews.DownloadedPhoto{Projected: p, View: "default"}
}

// NewPageOfStations initializes result type PageOfStations from viewed result
// type PageOfStations.
func NewPageOfStations(vres *stationviews.PageOfStations) *PageOfStations {
	return newPageOfStations(vres.Projected)
}

// NewViewedPageOfStations initializes viewed result type PageOfStations from
// result type PageOfStations using the given view.
func NewViewedPageOfStations(res *PageOfStations, view string) *stationviews.PageOfStations {
	p := newPageOfStationsView(res)
	return &stationviews.PageOfStations{Projected: p, View: "default"}
}

// newStationFull converts projected type StationFull to service type
// StationFull.
func newStationFull(vres *stationviews.StationFullView) *StationFull {
	res := &StationFull{
		Battery:            vres.Battery,
		RecordingStartedAt: vres.RecordingStartedAt,
		MemoryUsed:         vres.MemoryUsed,
		MemoryAvailable:    vres.MemoryAvailable,
		FirmwareNumber:     vres.FirmwareNumber,
		FirmwareTime:       vres.FirmwareTime,
		LocationName:       vres.LocationName,
		PlaceNameOther:     vres.PlaceNameOther,
		PlaceNameNative:    vres.PlaceNameNative,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.DeviceID != nil {
		res.DeviceID = *vres.DeviceID
	}
	if vres.ReadOnly != nil {
		res.ReadOnly = *vres.ReadOnly
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Owner != nil {
		res.Owner = transformStationviewsStationOwnerViewToStationOwner(vres.Owner)
	}
	if vres.Uploads != nil {
		res.Uploads = make([]*StationUpload, len(vres.Uploads))
		for i, val := range vres.Uploads {
			res.Uploads[i] = transformStationviewsStationUploadViewToStationUpload(val)
		}
	}
	if vres.Photos != nil {
		res.Photos = transformStationviewsStationPhotosViewToStationPhotos(vres.Photos)
	}
	if vres.Configurations != nil {
		res.Configurations = transformStationviewsStationConfigurationsViewToStationConfigurations(vres.Configurations)
	}
	if vres.Data != nil {
		res.Data = transformStationviewsStationDataSummaryViewToStationDataSummary(vres.Data)
	}
	if vres.Location != nil {
		res.Location = newStationLocation(vres.Location)
	}
	return res
}

// newStationFullView projects result type StationFull to projected type
// StationFullView using the "default" view.
func newStationFullView(res *StationFull) *stationviews.StationFullView {
	vres := &stationviews.StationFullView{
		ID:                 &res.ID,
		Name:               &res.Name,
		DeviceID:           &res.DeviceID,
		ReadOnly:           &res.ReadOnly,
		Battery:            res.Battery,
		RecordingStartedAt: res.RecordingStartedAt,
		MemoryUsed:         res.MemoryUsed,
		MemoryAvailable:    res.MemoryAvailable,
		FirmwareNumber:     res.FirmwareNumber,
		FirmwareTime:       res.FirmwareTime,
		UpdatedAt:          &res.UpdatedAt,
		LocationName:       res.LocationName,
		PlaceNameOther:     res.PlaceNameOther,
		PlaceNameNative:    res.PlaceNameNative,
	}
	if res.Owner != nil {
		vres.Owner = transformStationOwnerToStationviewsStationOwnerView(res.Owner)
	}
	if res.Uploads != nil {
		vres.Uploads = make([]*stationviews.StationUploadView, len(res.Uploads))
		for i, val := range res.Uploads {
			vres.Uploads[i] = transformStationUploadToStationviewsStationUploadView(val)
		}
	}
	if res.Photos != nil {
		vres.Photos = transformStationPhotosToStationviewsStationPhotosView(res.Photos)
	}
	if res.Configurations != nil {
		vres.Configurations = transformStationConfigurationsToStationviewsStationConfigurationsView(res.Configurations)
	}
	if res.Data != nil {
		vres.Data = transformStationDataSummaryToStationviewsStationDataSummaryView(res.Data)
	}
	if res.Location != nil {
		vres.Location = newStationLocationView(res.Location)
	}
	return vres
}

// newStationLocation converts projected type StationLocation to service type
// StationLocation.
func newStationLocation(vres *stationviews.StationLocationView) *StationLocation {
	res := &StationLocation{
		URL: vres.URL,
	}
	if vres.Precise != nil {
		res.Precise = make([]float64, len(vres.Precise))
		for i, val := range vres.Precise {
			res.Precise[i] = val
		}
	}
	if vres.Regions != nil {
		res.Regions = make([]*StationRegion, len(vres.Regions))
		for i, val := range vres.Regions {
			res.Regions[i] = transformStationviewsStationRegionViewToStationRegion(val)
		}
	}
	return res
}

// newStationLocationView projects result type StationLocation to projected
// type StationLocationView using the "default" view.
func newStationLocationView(res *StationLocation) *stationviews.StationLocationView {
	vres := &stationviews.StationLocationView{
		URL: res.URL,
	}
	if res.Precise != nil {
		vres.Precise = make([]float64, len(res.Precise))
		for i, val := range res.Precise {
			vres.Precise[i] = val
		}
	}
	if res.Regions != nil {
		vres.Regions = make([]*stationviews.StationRegionView, len(res.Regions))
		for i, val := range res.Regions {
			vres.Regions[i] = transformStationRegionToStationviewsStationRegionView(val)
		}
	}
	return vres
}

// newStationsFull converts projected type StationsFull to service type
// StationsFull.
func newStationsFull(vres *stationviews.StationsFullView) *StationsFull {
	res := &StationsFull{}
	if vres.Stations != nil {
		res.Stations = newStationFullCollection(vres.Stations)
	}
	return res
}

// newStationsFullView projects result type StationsFull to projected type
// StationsFullView using the "default" view.
func newStationsFullView(res *StationsFull) *stationviews.StationsFullView {
	vres := &stationviews.StationsFullView{}
	if res.Stations != nil {
		vres.Stations = newStationFullCollectionView(res.Stations)
	}
	return vres
}

// newStationFullCollection converts projected type StationFullCollection to
// service type StationFullCollection.
func newStationFullCollection(vres stationviews.StationFullCollectionView) StationFullCollection {
	res := make(StationFullCollection, len(vres))
	for i, n := range vres {
		res[i] = newStationFull(n)
	}
	return res
}

// newStationFullCollectionView projects result type StationFullCollection to
// projected type StationFullCollectionView using the "default" view.
func newStationFullCollectionView(res StationFullCollection) stationviews.StationFullCollectionView {
	vres := make(stationviews.StationFullCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStationFullView(n)
	}
	return vres
}

// newDownloadedPhoto converts projected type DownloadedPhoto to service type
// DownloadedPhoto.
func newDownloadedPhoto(vres *stationviews.DownloadedPhotoView) *DownloadedPhoto {
	res := &DownloadedPhoto{
		Body: vres.Body,
	}
	if vres.Length != nil {
		res.Length = *vres.Length
	}
	if vres.ContentType != nil {
		res.ContentType = *vres.ContentType
	}
	if vres.Etag != nil {
		res.Etag = *vres.Etag
	}
	return res
}

// newDownloadedPhotoView projects result type DownloadedPhoto to projected
// type DownloadedPhotoView using the "default" view.
func newDownloadedPhotoView(res *DownloadedPhoto) *stationviews.DownloadedPhotoView {
	vres := &stationviews.DownloadedPhotoView{
		Length:      &res.Length,
		ContentType: &res.ContentType,
		Etag:        &res.Etag,
		Body:        res.Body,
	}
	return vres
}

// newPageOfStations converts projected type PageOfStations to service type
// PageOfStations.
func newPageOfStations(vres *stationviews.PageOfStationsView) *PageOfStations {
	res := &PageOfStations{}
	if vres.Total != nil {
		res.Total = *vres.Total
	}
	if vres.Stations != nil {
		res.Stations = make([]*EssentialStation, len(vres.Stations))
		for i, val := range vres.Stations {
			res.Stations[i] = transformStationviewsEssentialStationViewToEssentialStation(val)
		}
	}
	return res
}

// newPageOfStationsView projects result type PageOfStations to projected type
// PageOfStationsView using the "default" view.
func newPageOfStationsView(res *PageOfStations) *stationviews.PageOfStationsView {
	vres := &stationviews.PageOfStationsView{
		Total: &res.Total,
	}
	if res.Stations != nil {
		vres.Stations = make([]*stationviews.EssentialStationView, len(res.Stations))
		for i, val := range res.Stations {
			vres.Stations[i] = transformEssentialStationToStationviewsEssentialStationView(val)
		}
	}
	return vres
}

// transformStationviewsStationOwnerViewToStationOwner builds a value of type
// *StationOwner from a value of type *stationviews.StationOwnerView.
func transformStationviewsStationOwnerViewToStationOwner(v *stationviews.StationOwnerView) *StationOwner {
	if v == nil {
		return nil
	}
	res := &StationOwner{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformStationviewsStationUploadViewToStationUpload builds a value of type
// *StationUpload from a value of type *stationviews.StationUploadView.
func transformStationviewsStationUploadViewToStationUpload(v *stationviews.StationUploadView) *StationUpload {
	if v == nil {
		return nil
	}
	res := &StationUpload{
		ID:       *v.ID,
		Time:     *v.Time,
		UploadID: *v.UploadID,
		Size:     *v.Size,
		URL:      *v.URL,
		Type:     *v.Type,
	}
	if v.Blocks != nil {
		res.Blocks = make([]int64, len(v.Blocks))
		for i, val := range v.Blocks {
			res.Blocks[i] = val
		}
	}

	return res
}

// transformStationviewsStationPhotosViewToStationPhotos builds a value of type
// *StationPhotos from a value of type *stationviews.StationPhotosView.
func transformStationviewsStationPhotosViewToStationPhotos(v *stationviews.StationPhotosView) *StationPhotos {
	if v == nil {
		return nil
	}
	res := &StationPhotos{
		Small: *v.Small,
	}

	return res
}

// transformStationviewsStationConfigurationsViewToStationConfigurations builds
// a value of type *StationConfigurations from a value of type
// *stationviews.StationConfigurationsView.
func transformStationviewsStationConfigurationsViewToStationConfigurations(v *stationviews.StationConfigurationsView) *StationConfigurations {
	if v == nil {
		return nil
	}
	res := &StationConfigurations{}
	if v.All != nil {
		res.All = make([]*StationConfiguration, len(v.All))
		for i, val := range v.All {
			res.All[i] = transformStationviewsStationConfigurationViewToStationConfiguration(val)
		}
	}

	return res
}

// transformStationviewsStationConfigurationViewToStationConfiguration builds a
// value of type *StationConfiguration from a value of type
// *stationviews.StationConfigurationView.
func transformStationviewsStationConfigurationViewToStationConfiguration(v *stationviews.StationConfigurationView) *StationConfiguration {
	res := &StationConfiguration{
		ID:           *v.ID,
		Time:         *v.Time,
		ProvisionID:  *v.ProvisionID,
		MetaRecordID: v.MetaRecordID,
		SourceID:     v.SourceID,
	}
	if v.Modules != nil {
		res.Modules = make([]*StationModule, len(v.Modules))
		for i, val := range v.Modules {
			res.Modules[i] = transformStationviewsStationModuleViewToStationModule(val)
		}
	}

	return res
}

// transformStationviewsStationModuleViewToStationModule builds a value of type
// *StationModule from a value of type *stationviews.StationModuleView.
func transformStationviewsStationModuleViewToStationModule(v *stationviews.StationModuleView) *StationModule {
	res := &StationModule{
		ID:           *v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         *v.Name,
		Position:     *v.Position,
		Flags:        *v.Flags,
		Internal:     *v.Internal,
		FullKey:      *v.FullKey,
	}
	if v.Sensors != nil {
		res.Sensors = make([]*StationSensor, len(v.Sensors))
		for i, val := range v.Sensors {
			res.Sensors[i] = transformStationviewsStationSensorViewToStationSensor(val)
		}
	}

	return res
}

// transformStationviewsStationSensorViewToStationSensor builds a value of type
// *StationSensor from a value of type *stationviews.StationSensorView.
func transformStationviewsStationSensorViewToStationSensor(v *stationviews.StationSensorView) *StationSensor {
	res := &StationSensor{
		Name:          *v.Name,
		UnitOfMeasure: *v.UnitOfMeasure,
		Key:           *v.Key,
		FullKey:       *v.FullKey,
	}
	if v.Reading != nil {
		res.Reading = transformStationviewsSensorReadingViewToSensorReading(v.Reading)
	}
	if v.Ranges != nil {
		res.Ranges = make([]*SensorRange, len(v.Ranges))
		for i, val := range v.Ranges {
			res.Ranges[i] = transformStationviewsSensorRangeViewToSensorRange(val)
		}
	}

	return res
}

// transformStationviewsSensorReadingViewToSensorReading builds a value of type
// *SensorReading from a value of type *stationviews.SensorReadingView.
func transformStationviewsSensorReadingViewToSensorReading(v *stationviews.SensorReadingView) *SensorReading {
	if v == nil {
		return nil
	}
	res := &SensorReading{
		Last: *v.Last,
		Time: *v.Time,
	}

	return res
}

// transformStationviewsSensorRangeViewToSensorRange builds a value of type
// *SensorRange from a value of type *stationviews.SensorRangeView.
func transformStationviewsSensorRangeViewToSensorRange(v *stationviews.SensorRangeView) *SensorRange {
	res := &SensorRange{
		Minimum: *v.Minimum,
		Maximum: *v.Maximum,
	}

	return res
}

// transformStationviewsStationDataSummaryViewToStationDataSummary builds a
// value of type *StationDataSummary from a value of type
// *stationviews.StationDataSummaryView.
func transformStationviewsStationDataSummaryViewToStationDataSummary(v *stationviews.StationDataSummaryView) *StationDataSummary {
	if v == nil {
		return nil
	}
	res := &StationDataSummary{
		Start:           *v.Start,
		End:             *v.End,
		NumberOfSamples: *v.NumberOfSamples,
	}

	return res
}

// transformStationOwnerToStationviewsStationOwnerView builds a value of type
// *stationviews.StationOwnerView from a value of type *StationOwner.
func transformStationOwnerToStationviewsStationOwnerView(v *StationOwner) *stationviews.StationOwnerView {
	res := &stationviews.StationOwnerView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformStationUploadToStationviewsStationUploadView builds a value of type
// *stationviews.StationUploadView from a value of type *StationUpload.
func transformStationUploadToStationviewsStationUploadView(v *StationUpload) *stationviews.StationUploadView {
	res := &stationviews.StationUploadView{
		ID:       &v.ID,
		Time:     &v.Time,
		UploadID: &v.UploadID,
		Size:     &v.Size,
		URL:      &v.URL,
		Type:     &v.Type,
	}
	if v.Blocks != nil {
		res.Blocks = make([]int64, len(v.Blocks))
		for i, val := range v.Blocks {
			res.Blocks[i] = val
		}
	}

	return res
}

// transformStationPhotosToStationviewsStationPhotosView builds a value of type
// *stationviews.StationPhotosView from a value of type *StationPhotos.
func transformStationPhotosToStationviewsStationPhotosView(v *StationPhotos) *stationviews.StationPhotosView {
	res := &stationviews.StationPhotosView{
		Small: &v.Small,
	}

	return res
}

// transformStationConfigurationsToStationviewsStationConfigurationsView builds
// a value of type *stationviews.StationConfigurationsView from a value of type
// *StationConfigurations.
func transformStationConfigurationsToStationviewsStationConfigurationsView(v *StationConfigurations) *stationviews.StationConfigurationsView {
	res := &stationviews.StationConfigurationsView{}
	if v.All != nil {
		res.All = make([]*stationviews.StationConfigurationView, len(v.All))
		for i, val := range v.All {
			res.All[i] = transformStationConfigurationToStationviewsStationConfigurationView(val)
		}
	}

	return res
}

// transformStationConfigurationToStationviewsStationConfigurationView builds a
// value of type *stationviews.StationConfigurationView from a value of type
// *StationConfiguration.
func transformStationConfigurationToStationviewsStationConfigurationView(v *StationConfiguration) *stationviews.StationConfigurationView {
	res := &stationviews.StationConfigurationView{
		ID:           &v.ID,
		Time:         &v.Time,
		ProvisionID:  &v.ProvisionID,
		MetaRecordID: v.MetaRecordID,
		SourceID:     v.SourceID,
	}
	if v.Modules != nil {
		res.Modules = make([]*stationviews.StationModuleView, len(v.Modules))
		for i, val := range v.Modules {
			res.Modules[i] = transformStationModuleToStationviewsStationModuleView(val)
		}
	}

	return res
}

// transformStationModuleToStationviewsStationModuleView builds a value of type
// *stationviews.StationModuleView from a value of type *StationModule.
func transformStationModuleToStationviewsStationModuleView(v *StationModule) *stationviews.StationModuleView {
	res := &stationviews.StationModuleView{
		ID:           &v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         &v.Name,
		Position:     &v.Position,
		Flags:        &v.Flags,
		Internal:     &v.Internal,
		FullKey:      &v.FullKey,
	}
	if v.Sensors != nil {
		res.Sensors = make([]*stationviews.StationSensorView, len(v.Sensors))
		for i, val := range v.Sensors {
			res.Sensors[i] = transformStationSensorToStationviewsStationSensorView(val)
		}
	}

	return res
}

// transformStationSensorToStationviewsStationSensorView builds a value of type
// *stationviews.StationSensorView from a value of type *StationSensor.
func transformStationSensorToStationviewsStationSensorView(v *StationSensor) *stationviews.StationSensorView {
	res := &stationviews.StationSensorView{
		Name:          &v.Name,
		UnitOfMeasure: &v.UnitOfMeasure,
		Key:           &v.Key,
		FullKey:       &v.FullKey,
	}
	if v.Reading != nil {
		res.Reading = transformSensorReadingToStationviewsSensorReadingView(v.Reading)
	}
	if v.Ranges != nil {
		res.Ranges = make([]*stationviews.SensorRangeView, len(v.Ranges))
		for i, val := range v.Ranges {
			res.Ranges[i] = transformSensorRangeToStationviewsSensorRangeView(val)
		}
	}

	return res
}

// transformSensorReadingToStationviewsSensorReadingView builds a value of type
// *stationviews.SensorReadingView from a value of type *SensorReading.
func transformSensorReadingToStationviewsSensorReadingView(v *SensorReading) *stationviews.SensorReadingView {
	if v == nil {
		return nil
	}
	res := &stationviews.SensorReadingView{
		Last: &v.Last,
		Time: &v.Time,
	}

	return res
}

// transformSensorRangeToStationviewsSensorRangeView builds a value of type
// *stationviews.SensorRangeView from a value of type *SensorRange.
func transformSensorRangeToStationviewsSensorRangeView(v *SensorRange) *stationviews.SensorRangeView {
	res := &stationviews.SensorRangeView{
		Minimum: &v.Minimum,
		Maximum: &v.Maximum,
	}

	return res
}

// transformStationDataSummaryToStationviewsStationDataSummaryView builds a
// value of type *stationviews.StationDataSummaryView from a value of type
// *StationDataSummary.
func transformStationDataSummaryToStationviewsStationDataSummaryView(v *StationDataSummary) *stationviews.StationDataSummaryView {
	if v == nil {
		return nil
	}
	res := &stationviews.StationDataSummaryView{
		Start:           &v.Start,
		End:             &v.End,
		NumberOfSamples: &v.NumberOfSamples,
	}

	return res
}

// transformStationviewsStationRegionViewToStationRegion builds a value of type
// *StationRegion from a value of type *stationviews.StationRegionView.
func transformStationviewsStationRegionViewToStationRegion(v *stationviews.StationRegionView) *StationRegion {
	if v == nil {
		return nil
	}
	res := &StationRegion{
		Name: *v.Name,
	}
	if v.Shape != nil {
		res.Shape = make([][][]float64, len(v.Shape))
		for i, val := range v.Shape {
			res.Shape[i] = make([][]float64, len(val))
			for j, val := range val {
				res.Shape[i][j] = make([]float64, len(val))
				for k, val := range val {
					res.Shape[i][j][k] = val
				}
			}
		}
	}

	return res
}

// transformStationRegionToStationviewsStationRegionView builds a value of type
// *stationviews.StationRegionView from a value of type *StationRegion.
func transformStationRegionToStationviewsStationRegionView(v *StationRegion) *stationviews.StationRegionView {
	if v == nil {
		return nil
	}
	res := &stationviews.StationRegionView{
		Name: &v.Name,
	}
	if v.Shape != nil {
		res.Shape = make([][][]float64, len(v.Shape))
		for i, val := range v.Shape {
			res.Shape[i] = make([][]float64, len(val))
			for j, val := range val {
				res.Shape[i][j] = make([]float64, len(val))
				for k, val := range val {
					res.Shape[i][j][k] = val
				}
			}
		}
	}

	return res
}

// transformStationviewsEssentialStationViewToEssentialStation builds a value
// of type *EssentialStation from a value of type
// *stationviews.EssentialStationView.
func transformStationviewsEssentialStationViewToEssentialStation(v *stationviews.EssentialStationView) *EssentialStation {
	if v == nil {
		return nil
	}
	res := &EssentialStation{
		ID:                 *v.ID,
		DeviceID:           *v.DeviceID,
		Name:               *v.Name,
		CreatedAt:          *v.CreatedAt,
		UpdatedAt:          *v.UpdatedAt,
		RecordingStartedAt: v.RecordingStartedAt,
		MemoryUsed:         v.MemoryUsed,
		MemoryAvailable:    v.MemoryAvailable,
		FirmwareNumber:     v.FirmwareNumber,
		FirmwareTime:       v.FirmwareTime,
		LastIngestionAt:    v.LastIngestionAt,
	}
	if v.Owner != nil {
		res.Owner = transformStationviewsStationOwnerViewToStationOwner(v.Owner)
	}
	if v.Location != nil {
		res.Location = transformStationviewsStationLocationViewToStationLocation(v.Location)
	}

	return res
}

// transformStationviewsStationLocationViewToStationLocation builds a value of
// type *StationLocation from a value of type *stationviews.StationLocationView.
func transformStationviewsStationLocationViewToStationLocation(v *stationviews.StationLocationView) *StationLocation {
	if v == nil {
		return nil
	}
	res := &StationLocation{
		URL: v.URL,
	}
	if v.Precise != nil {
		res.Precise = make([]float64, len(v.Precise))
		for i, val := range v.Precise {
			res.Precise[i] = val
		}
	}
	if v.Regions != nil {
		res.Regions = make([]*StationRegion, len(v.Regions))
		for i, val := range v.Regions {
			res.Regions[i] = transformStationviewsStationRegionViewToStationRegion(val)
		}
	}

	return res
}

// transformEssentialStationToStationviewsEssentialStationView builds a value
// of type *stationviews.EssentialStationView from a value of type
// *EssentialStation.
func transformEssentialStationToStationviewsEssentialStationView(v *EssentialStation) *stationviews.EssentialStationView {
	res := &stationviews.EssentialStationView{
		ID:                 &v.ID,
		DeviceID:           &v.DeviceID,
		Name:               &v.Name,
		CreatedAt:          &v.CreatedAt,
		UpdatedAt:          &v.UpdatedAt,
		RecordingStartedAt: v.RecordingStartedAt,
		MemoryUsed:         v.MemoryUsed,
		MemoryAvailable:    v.MemoryAvailable,
		FirmwareNumber:     v.FirmwareNumber,
		FirmwareTime:       v.FirmwareTime,
		LastIngestionAt:    v.LastIngestionAt,
	}
	if v.Owner != nil {
		res.Owner = transformStationOwnerToStationviewsStationOwnerView(v.Owner)
	}
	if v.Location != nil {
		res.Location = transformStationLocationToStationviewsStationLocationView(v.Location)
	}

	return res
}

// transformStationLocationToStationviewsStationLocationView builds a value of
// type *stationviews.StationLocationView from a value of type *StationLocation.
func transformStationLocationToStationviewsStationLocationView(v *StationLocation) *stationviews.StationLocationView {
	if v == nil {
		return nil
	}
	res := &stationviews.StationLocationView{
		URL: v.URL,
	}
	if v.Precise != nil {
		res.Precise = make([]float64, len(v.Precise))
		for i, val := range v.Precise {
			res.Precise[i] = val
		}
	}
	if v.Regions != nil {
		res.Regions = make([]*stationviews.StationRegionView, len(v.Regions))
		for i, val := range v.Regions {
			res.Regions[i] = transformStationRegionToStationviewsStationRegionView(val)
		}
	}

	return res
}
