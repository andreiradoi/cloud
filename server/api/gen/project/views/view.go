// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project views
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ProjectUpdate is the viewed result type that is projected based on a view.
type ProjectUpdate struct {
	// Type to project
	Projected *ProjectUpdateView
	// View to render
	View string
}

// PendingInvites is the viewed result type that is projected based on a view.
type PendingInvites struct {
	// Type to project
	Projected *PendingInvitesView
	// View to render
	View string
}

// Project is the viewed result type that is projected based on a view.
type Project struct {
	// Type to project
	Projected *ProjectView
	// View to render
	View string
}

// Projects is the viewed result type that is projected based on a view.
type Projects struct {
	// Type to project
	Projected *ProjectsView
	// View to render
	View string
}

// DownloadedPhoto is the viewed result type that is projected based on a view.
type DownloadedPhoto struct {
	// Type to project
	Projected *DownloadedPhotoView
	// View to render
	View string
}

// ProjectUpdateView is a type that runs validations on a projected type.
type ProjectUpdateView struct {
	ID        *int64
	Body      *string
	CreatedAt *int64
}

// PendingInvitesView is a type that runs validations on a projected type.
type PendingInvitesView struct {
	Pending  []*PendingInviteView
	Projects ProjectCollectionView
}

// PendingInviteView is a type that runs validations on a projected type.
type PendingInviteView struct {
	ID      *int64
	Project *ProjectSummaryView
	Time    *int64
	Role    *int32
}

// ProjectSummaryView is a type that runs validations on a projected type.
type ProjectSummaryView struct {
	ID   *int64
	Name *string
}

// ProjectCollectionView is a type that runs validations on a projected type.
type ProjectCollectionView []*ProjectView

// ProjectView is a type that runs validations on a projected type.
type ProjectView struct {
	ID          *int32
	Name        *string
	Description *string
	Goal        *string
	Location    *string
	Tags        *string
	Private     *bool
	StartTime   *string
	EndTime     *string
	Photo       *string
	ReadOnly    *bool
	Following   *ProjectFollowingView
}

// ProjectFollowingView is a type that runs validations on a projected type.
type ProjectFollowingView struct {
	Total     *int32
	Following *bool
}

// ProjectsView is a type that runs validations on a projected type.
type ProjectsView struct {
	Projects ProjectCollectionView
}

// DownloadedPhotoView is a type that runs validations on a projected type.
type DownloadedPhotoView struct {
	Length      *int64
	ContentType *string
	Etag        *string
	Body        []byte
}

var (
	// ProjectUpdateMap is a map of attribute names in result type ProjectUpdate
	// indexed by view name.
	ProjectUpdateMap = map[string][]string{
		"default": []string{
			"id",
			"body",
		},
	}
	// PendingInvitesMap is a map of attribute names in result type PendingInvites
	// indexed by view name.
	PendingInvitesMap = map[string][]string{
		"default": []string{
			"pending",
			"projects",
		},
	}
	// ProjectMap is a map of attribute names in result type Project indexed by
	// view name.
	ProjectMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"description",
			"goal",
			"location",
			"tags",
			"private",
			"startTime",
			"endTime",
			"photo",
			"readOnly",
			"following",
		},
	}
	// ProjectsMap is a map of attribute names in result type Projects indexed by
	// view name.
	ProjectsMap = map[string][]string{
		"default": []string{
			"projects",
		},
	}
	// DownloadedPhotoMap is a map of attribute names in result type
	// DownloadedPhoto indexed by view name.
	DownloadedPhotoMap = map[string][]string{
		"default": []string{
			"length",
			"body",
			"contentType",
			"etag",
		},
	}
	// ProjectCollectionMap is a map of attribute names in result type
	// ProjectCollection indexed by view name.
	ProjectCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"description",
			"goal",
			"location",
			"tags",
			"private",
			"startTime",
			"endTime",
			"photo",
			"readOnly",
			"following",
		},
	}
)

// ValidateProjectUpdate runs the validations defined on the viewed result type
// ProjectUpdate.
func ValidateProjectUpdate(result *ProjectUpdate) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectUpdateView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidatePendingInvites runs the validations defined on the viewed result
// type PendingInvites.
func ValidatePendingInvites(result *PendingInvites) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePendingInvitesView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProject runs the validations defined on the viewed result type
// Project.
func ValidateProject(result *Project) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProjects runs the validations defined on the viewed result type
// Projects.
func ValidateProjects(result *Projects) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateDownloadedPhoto runs the validations defined on the viewed result
// type DownloadedPhoto.
func ValidateDownloadedPhoto(result *DownloadedPhoto) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateDownloadedPhotoView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProjectUpdateView runs the validations defined on ProjectUpdateView
// using the "default" view.
func ValidateProjectUpdateView(result *ProjectUpdateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "result"))
	}
	return
}

// ValidatePendingInvitesView runs the validations defined on
// PendingInvitesView using the "default" view.
func ValidatePendingInvitesView(result *PendingInvitesView) (err error) {
	if result.Pending == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pending", "result"))
	}
	for _, e := range result.Pending {
		if e != nil {
			if err2 := ValidatePendingInviteView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Projects != nil {
		if err2 := ValidateProjectCollectionView(result.Projects); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePendingInviteView runs the validations defined on PendingInviteView.
func ValidatePendingInviteView(result *PendingInviteView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Project == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("project", "result"))
	}
	if result.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "result"))
	}
	if result.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "result"))
	}
	if result.Project != nil {
		if err2 := ValidateProjectSummaryView(result.Project); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectSummaryView runs the validations defined on
// ProjectSummaryView.
func ValidateProjectSummaryView(result *ProjectSummaryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateProjectCollectionView runs the validations defined on
// ProjectCollectionView using the "default" view.
func ValidateProjectCollectionView(result ProjectCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateProjectView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectView runs the validations defined on ProjectView using the
// "default" view.
func ValidateProjectView(result *ProjectView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.Goal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("goal", "result"))
	}
	if result.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "result"))
	}
	if result.Private == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("private", "result"))
	}
	if result.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "result"))
	}
	if result.ReadOnly == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("readOnly", "result"))
	}
	if result.Following == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("following", "result"))
	}
	if result.Following != nil {
		if err2 := ValidateProjectFollowingView(result.Following); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectFollowingView runs the validations defined on
// ProjectFollowingView.
func ValidateProjectFollowingView(result *ProjectFollowingView) (err error) {
	if result.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "result"))
	}
	if result.Following == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("following", "result"))
	}
	return
}

// ValidateProjectsView runs the validations defined on ProjectsView using the
// "default" view.
func ValidateProjectsView(result *ProjectsView) (err error) {

	if result.Projects != nil {
		if err2 := ValidateProjectCollectionView(result.Projects); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDownloadedPhotoView runs the validations defined on
// DownloadedPhotoView using the "default" view.
func ValidateDownloadedPhotoView(result *DownloadedPhotoView) (err error) {
	if result.Length == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("length", "result"))
	}
	if result.ContentType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contentType", "result"))
	}
	if result.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("etag", "result"))
	}
	return
}
