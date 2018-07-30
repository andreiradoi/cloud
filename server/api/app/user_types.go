// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "fieldkit": Application User Types
//
// Command:
// $ main

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// addAdministratorPayload user type.
type addAdministratorPayload struct {
	UserID *int `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the addAdministratorPayload type instance.
func (ut *addAdministratorPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	return
}

// Publicize creates AddAdministratorPayload from addAdministratorPayload
func (ut *addAdministratorPayload) Publicize() *AddAdministratorPayload {
	var pub AddAdministratorPayload
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// AddAdministratorPayload user type.
type AddAdministratorPayload struct {
	UserID int `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// addDeviceSourcePayload user type.
type addDeviceSourcePayload struct {
	Key  *string `form:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty" xml:"key,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the addDeviceSourcePayload type instance.
func (ut *addDeviceSourcePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Key == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "key"))
	}
	return
}

// Publicize creates AddDeviceSourcePayload from addDeviceSourcePayload
func (ut *addDeviceSourcePayload) Publicize() *AddDeviceSourcePayload {
	var pub AddDeviceSourcePayload
	if ut.Key != nil {
		pub.Key = *ut.Key
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// AddDeviceSourcePayload user type.
type AddDeviceSourcePayload struct {
	Key  string `form:"key" json:"key" yaml:"key" xml:"key"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the AddDeviceSourcePayload type instance.
func (ut *AddDeviceSourcePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "key"))
	}
	return
}

// addExpeditionPayload user type.
type addExpeditionPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" yaml:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addExpeditionPayload type instance.
func (ut *addExpeditionPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "description"))
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddExpeditionPayload from addExpeditionPayload
func (ut *addExpeditionPayload) Publicize() *AddExpeditionPayload {
	var pub AddExpeditionPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddExpeditionPayload user type.
type AddExpeditionPayload struct {
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	Name        string `form:"name" json:"name" yaml:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" yaml:"slug" xml:"slug"`
}

// Validate validates the AddExpeditionPayload type instance.
func (ut *AddExpeditionPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addFirmwarePayload user type.
type addFirmwarePayload struct {
	Etag *string `form:"etag,omitempty" json:"etag,omitempty" yaml:"etag,omitempty" xml:"etag,omitempty"`
	URL  *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the addFirmwarePayload type instance.
func (ut *addFirmwarePayload) Validate() (err error) {
	if ut.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "etag"))
	}
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url"))
	}
	return
}

// Publicize creates AddFirmwarePayload from addFirmwarePayload
func (ut *addFirmwarePayload) Publicize() *AddFirmwarePayload {
	var pub AddFirmwarePayload
	if ut.Etag != nil {
		pub.Etag = *ut.Etag
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// AddFirmwarePayload user type.
type AddFirmwarePayload struct {
	Etag string `form:"etag" json:"etag" yaml:"etag" xml:"etag"`
	URL  string `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the AddFirmwarePayload type instance.
func (ut *AddFirmwarePayload) Validate() (err error) {
	if ut.Etag == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "etag"))
	}
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url"))
	}
	return
}

// addMemberPayload user type.
type addMemberPayload struct {
	Role   *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
	UserID *int    `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the addMemberPayload type instance.
func (ut *addMemberPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "userId"))
	}
	if ut.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "role"))
	}
	return
}

// Publicize creates AddMemberPayload from addMemberPayload
func (ut *addMemberPayload) Publicize() *AddMemberPayload {
	var pub AddMemberPayload
	if ut.Role != nil {
		pub.Role = *ut.Role
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// AddMemberPayload user type.
type AddMemberPayload struct {
	Role   string `form:"role" json:"role" yaml:"role" xml:"role"`
	UserID int    `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the AddMemberPayload type instance.
func (ut *AddMemberPayload) Validate() (err error) {

	if ut.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "role"))
	}
	return
}

// addProjectPayload user type.
type addProjectPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" yaml:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addProjectPayload type instance.
func (ut *addProjectPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "description"))
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddProjectPayload from addProjectPayload
func (ut *addProjectPayload) Publicize() *AddProjectPayload {
	var pub AddProjectPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddProjectPayload user type.
type AddProjectPayload struct {
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	Name        string `form:"name" json:"name" yaml:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" yaml:"slug" xml:"slug"`
}

// Validate validates the AddProjectPayload type instance.
func (ut *AddProjectPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addTeamPayload user type.
type addTeamPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" yaml:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addTeamPayload type instance.
func (ut *addTeamPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "description"))
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddTeamPayload from addTeamPayload
func (ut *addTeamPayload) Publicize() *AddTeamPayload {
	var pub AddTeamPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddTeamPayload user type.
type AddTeamPayload struct {
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	Name        string `form:"name" json:"name" yaml:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" yaml:"slug" xml:"slug"`
}

// Validate validates the AddTeamPayload type instance.
func (ut *AddTeamPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "description"))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addTwitterAccountSourcePayload user type.
type addTwitterAccountSourcePayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the addTwitterAccountSourcePayload type instance.
func (ut *addTwitterAccountSourcePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	return
}

// Publicize creates AddTwitterAccountSourcePayload from addTwitterAccountSourcePayload
func (ut *addTwitterAccountSourcePayload) Publicize() *AddTwitterAccountSourcePayload {
	var pub AddTwitterAccountSourcePayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// AddTwitterAccountSourcePayload user type.
type AddTwitterAccountSourcePayload struct {
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the AddTwitterAccountSourcePayload type instance.
func (ut *AddTwitterAccountSourcePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	return
}

// addUserPayload user type.
type addUserPayload struct {
	Bio         *string `form:"bio,omitempty" json:"bio,omitempty" yaml:"bio,omitempty" xml:"bio,omitempty"`
	Email       *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	InviteToken *string `form:"invite_token,omitempty" json:"invite_token,omitempty" yaml:"invite_token,omitempty" xml:"invite_token,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Password    *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the addUserPayload type instance.
func (ut *addUserPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "username"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "bio"))
	}
	if ut.InviteToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "invite_token"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 10, true))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates AddUserPayload from addUserPayload
func (ut *addUserPayload) Publicize() *AddUserPayload {
	var pub AddUserPayload
	if ut.Bio != nil {
		pub.Bio = *ut.Bio
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.InviteToken != nil {
		pub.InviteToken = *ut.InviteToken
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// AddUserPayload user type.
type AddUserPayload struct {
	Bio         string `form:"bio" json:"bio" yaml:"bio" xml:"bio"`
	Email       string `form:"email" json:"email" yaml:"email" xml:"email"`
	InviteToken string `form:"invite_token" json:"invite_token" yaml:"invite_token" xml:"invite_token"`
	Name        string `form:"name" json:"name" yaml:"name" xml:"name"`
	Password    string `form:"password" json:"password" yaml:"password" xml:"password"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the AddUserPayload type instance.
func (ut *AddUserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "username"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if ut.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "bio"))
	}
	if ut.InviteToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "invite_token"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if utf8.RuneCountInString(ut.Password) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 10, true))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}

// loginPayload user type.
type loginPayload struct {
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "username"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 10, true))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "username"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if utf8.RuneCountInString(ut.Password) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 10, true))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}

// updateDeviceFirmwarePayload user type.
type updateDeviceFirmwarePayload struct {
	DeviceID *int    `form:"deviceId,omitempty" json:"deviceId,omitempty" yaml:"deviceId,omitempty" xml:"deviceId,omitempty"`
	Etag     *string `form:"etag,omitempty" json:"etag,omitempty" yaml:"etag,omitempty" xml:"etag,omitempty"`
	URL      *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the updateDeviceFirmwarePayload type instance.
func (ut *updateDeviceFirmwarePayload) Validate() (err error) {
	if ut.DeviceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "deviceId"))
	}
	if ut.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "etag"))
	}
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url"))
	}
	return
}

// Publicize creates UpdateDeviceFirmwarePayload from updateDeviceFirmwarePayload
func (ut *updateDeviceFirmwarePayload) Publicize() *UpdateDeviceFirmwarePayload {
	var pub UpdateDeviceFirmwarePayload
	if ut.DeviceID != nil {
		pub.DeviceID = *ut.DeviceID
	}
	if ut.Etag != nil {
		pub.Etag = *ut.Etag
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// UpdateDeviceFirmwarePayload user type.
type UpdateDeviceFirmwarePayload struct {
	DeviceID int    `form:"deviceId" json:"deviceId" yaml:"deviceId" xml:"deviceId"`
	Etag     string `form:"etag" json:"etag" yaml:"etag" xml:"etag"`
	URL      string `form:"url" json:"url" yaml:"url" xml:"url"`
}

// Validate validates the UpdateDeviceFirmwarePayload type instance.
func (ut *UpdateDeviceFirmwarePayload) Validate() (err error) {

	if ut.Etag == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "etag"))
	}
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url"))
	}
	return
}

// updateDeviceSourceLocationPayload user type.
type updateDeviceSourceLocationPayload struct {
	Key       *string  `form:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty" xml:"key,omitempty"`
	Latitude  *float64 `form:"latitude,omitempty" json:"latitude,omitempty" yaml:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *float64 `form:"longitude,omitempty" json:"longitude,omitempty" yaml:"longitude,omitempty" xml:"longitude,omitempty"`
}

// Validate validates the updateDeviceSourceLocationPayload type instance.
func (ut *updateDeviceSourceLocationPayload) Validate() (err error) {
	if ut.Key == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "key"))
	}
	if ut.Longitude == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "longitude"))
	}
	if ut.Latitude == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "latitude"))
	}
	return
}

// Publicize creates UpdateDeviceSourceLocationPayload from updateDeviceSourceLocationPayload
func (ut *updateDeviceSourceLocationPayload) Publicize() *UpdateDeviceSourceLocationPayload {
	var pub UpdateDeviceSourceLocationPayload
	if ut.Key != nil {
		pub.Key = *ut.Key
	}
	if ut.Latitude != nil {
		pub.Latitude = *ut.Latitude
	}
	if ut.Longitude != nil {
		pub.Longitude = *ut.Longitude
	}
	return &pub
}

// UpdateDeviceSourceLocationPayload user type.
type UpdateDeviceSourceLocationPayload struct {
	Key       string  `form:"key" json:"key" yaml:"key" xml:"key"`
	Latitude  float64 `form:"latitude" json:"latitude" yaml:"latitude" xml:"latitude"`
	Longitude float64 `form:"longitude" json:"longitude" yaml:"longitude" xml:"longitude"`
}

// Validate validates the UpdateDeviceSourceLocationPayload type instance.
func (ut *UpdateDeviceSourceLocationPayload) Validate() (err error) {
	if ut.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "key"))
	}

	return
}

// updateDeviceSourcePayload user type.
type updateDeviceSourcePayload struct {
	Key  *string `form:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty" xml:"key,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the updateDeviceSourcePayload type instance.
func (ut *updateDeviceSourcePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Key == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "key"))
	}
	return
}

// Publicize creates UpdateDeviceSourcePayload from updateDeviceSourcePayload
func (ut *updateDeviceSourcePayload) Publicize() *UpdateDeviceSourcePayload {
	var pub UpdateDeviceSourcePayload
	if ut.Key != nil {
		pub.Key = *ut.Key
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// UpdateDeviceSourcePayload user type.
type UpdateDeviceSourcePayload struct {
	Key  string `form:"key" json:"key" yaml:"key" xml:"key"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the UpdateDeviceSourcePayload type instance.
func (ut *UpdateDeviceSourcePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "key"))
	}
	return
}

// updateDeviceSourceSchemaPayload user type.
type updateDeviceSourceSchemaPayload struct {
	Active     *bool   `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	JSONSchema *string `form:"jsonSchema,omitempty" json:"jsonSchema,omitempty" yaml:"jsonSchema,omitempty" xml:"jsonSchema,omitempty"`
	Key        *string `form:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty" xml:"key,omitempty"`
}

// Validate validates the updateDeviceSourceSchemaPayload type instance.
func (ut *updateDeviceSourceSchemaPayload) Validate() (err error) {
	if ut.Key == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "key"))
	}
	if ut.Active == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "active"))
	}
	if ut.JSONSchema == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "jsonSchema"))
	}
	return
}

// Publicize creates UpdateDeviceSourceSchemaPayload from updateDeviceSourceSchemaPayload
func (ut *updateDeviceSourceSchemaPayload) Publicize() *UpdateDeviceSourceSchemaPayload {
	var pub UpdateDeviceSourceSchemaPayload
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.JSONSchema != nil {
		pub.JSONSchema = *ut.JSONSchema
	}
	if ut.Key != nil {
		pub.Key = *ut.Key
	}
	return &pub
}

// UpdateDeviceSourceSchemaPayload user type.
type UpdateDeviceSourceSchemaPayload struct {
	Active     bool   `form:"active" json:"active" yaml:"active" xml:"active"`
	JSONSchema string `form:"jsonSchema" json:"jsonSchema" yaml:"jsonSchema" xml:"jsonSchema"`
	Key        string `form:"key" json:"key" yaml:"key" xml:"key"`
}

// Validate validates the UpdateDeviceSourceSchemaPayload type instance.
func (ut *UpdateDeviceSourceSchemaPayload) Validate() (err error) {
	if ut.Key == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "key"))
	}

	if ut.JSONSchema == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "jsonSchema"))
	}
	return
}

// updateMemberPayload user type.
type updateMemberPayload struct {
	Role *string `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}

// Validate validates the updateMemberPayload type instance.
func (ut *updateMemberPayload) Validate() (err error) {
	if ut.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "role"))
	}
	return
}

// Publicize creates UpdateMemberPayload from updateMemberPayload
func (ut *updateMemberPayload) Publicize() *UpdateMemberPayload {
	var pub UpdateMemberPayload
	if ut.Role != nil {
		pub.Role = *ut.Role
	}
	return &pub
}

// UpdateMemberPayload user type.
type UpdateMemberPayload struct {
	Role string `form:"role" json:"role" yaml:"role" xml:"role"`
}

// Validate validates the UpdateMemberPayload type instance.
func (ut *UpdateMemberPayload) Validate() (err error) {
	if ut.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "role"))
	}
	return
}

// updateSourcePayload user type.
type updateSourcePayload struct {
	Active *bool   `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	Name   *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"teamId,omitempty" json:"teamId,omitempty" yaml:"teamId,omitempty" xml:"teamId,omitempty"`
	UserID *int    `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the updateSourcePayload type instance.
func (ut *updateSourcePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	return
}

// Publicize creates UpdateSourcePayload from updateSourcePayload
func (ut *updateSourcePayload) Publicize() *UpdateSourcePayload {
	var pub UpdateSourcePayload
	if ut.Active != nil {
		pub.Active = ut.Active
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.TeamID != nil {
		pub.TeamID = ut.TeamID
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// UpdateSourcePayload user type.
type UpdateSourcePayload struct {
	Active *bool  `form:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty" xml:"active,omitempty"`
	Name   string `form:"name" json:"name" yaml:"name" xml:"name"`
	TeamID *int   `form:"teamId,omitempty" json:"teamId,omitempty" yaml:"teamId,omitempty" xml:"teamId,omitempty"`
	UserID *int   `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the UpdateSourcePayload type instance.
func (ut *UpdateSourcePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	return
}

// updateTwitterAccountSourcePayload user type.
type updateTwitterAccountSourcePayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"teamId,omitempty" json:"teamId,omitempty" yaml:"teamId,omitempty" xml:"teamId,omitempty"`
	UserID *int    `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// Publicize creates UpdateTwitterAccountSourcePayload from updateTwitterAccountSourcePayload
func (ut *updateTwitterAccountSourcePayload) Publicize() *UpdateTwitterAccountSourcePayload {
	var pub UpdateTwitterAccountSourcePayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.TeamID != nil {
		pub.TeamID = ut.TeamID
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// UpdateTwitterAccountSourcePayload user type.
type UpdateTwitterAccountSourcePayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"teamId,omitempty" json:"teamId,omitempty" yaml:"teamId,omitempty" xml:"teamId,omitempty"`
	UserID *int    `form:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty" xml:"userId,omitempty"`
}

// updateUserPayload user type.
type updateUserPayload struct {
	Bio   *string `form:"bio,omitempty" json:"bio,omitempty" yaml:"bio,omitempty" xml:"bio,omitempty"`
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the updateUserPayload type instance.
func (ut *updateUserPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "username"))
	}
	if ut.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "bio"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.Bio != nil {
		pub.Bio = *ut.Bio
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// UpdateUserPayload user type.
type UpdateUserPayload struct {
	Bio   string `form:"bio" json:"bio" yaml:"bio" xml:"bio"`
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	Name  string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Username
	Username string `form:"username" json:"username" yaml:"username" xml:"username"`
}

// Validate validates the UpdateUserPayload type instance.
func (ut *UpdateUserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "username"))
	}
	if ut.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "bio"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}
