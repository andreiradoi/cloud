// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Expedition media type (default view)
//
// Identifier: application/vnd.app.expedition+json; view=default
type Expedition struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Expedition media type instance.
func (mt *Expedition) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// ExpeditionCollection is the media type for an array of Expedition (default view)
//
// Identifier: application/vnd.app.expedition+json; type=collection; view=default
type ExpeditionCollection []*Expedition

// Validate validates the ExpeditionCollection media type instance.
func (mt ExpeditionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Expeditions media type (default view)
//
// Identifier: application/vnd.app.expeditions+json; view=default
type Expeditions struct {
	Expeditions ExpeditionCollection `form:"expeditions" json:"expeditions" xml:"expeditions"`
}

// Validate validates the Expeditions media type instance.
func (mt *Expeditions) Validate() (err error) {
	if mt.Expeditions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "expeditions"))
	}
	if err2 := mt.Expeditions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// Project media type (default view)
//
// Identifier: application/vnd.app.project+json; view=default
type Project struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
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
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
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
	Projects ProjectCollection `form:"projects" json:"projects" xml:"projects"`
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
	ID       int    `form:"id" json:"id" xml:"id"`
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, mt.Username, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, mt.Username, utf8.RuneCountInString(mt.Username), 40, false))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.app.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Users media type (default view)
//
// Identifier: application/vnd.app.users+json; view=default
type Users struct {
	Users UserCollection `form:"users" json:"users" xml:"users"`
}

// Validate validates the Users media type instance.
func (mt *Users) Validate() (err error) {
	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}
