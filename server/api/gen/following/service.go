// Code generated by goa v3.2.4, DO NOT EDIT.
//
// following service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package following

import (
	"context"

	followingviews "github.com/fieldkit/cloud/server/api/gen/following/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the following service interface.
type Service interface {
	// Follow implements follow.
	Follow(context.Context, *FollowPayload) (err error)
	// Unfollow implements unfollow.
	Unfollow(context.Context, *UnfollowPayload) (err error)
	// Followers implements followers.
	Followers(context.Context, *FollowersPayload) (res *FollowersPage, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "following"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"follow", "unfollow", "followers"}

// FollowPayload is the payload type of the following service follow method.
type FollowPayload struct {
	Auth *string
	ID   *int64
}

// UnfollowPayload is the payload type of the following service unfollow method.
type UnfollowPayload struct {
	Auth *string
	ID   *int64
}

// FollowersPayload is the payload type of the following service followers
// method.
type FollowersPayload struct {
	ID   *int64
	Page *int64
}

// FollowersPage is the result type of the following service followers method.
type FollowersPage struct {
	Followers FollowerCollection
	Total     int32
	Page      int32
}

type FollowerCollection []*Follower

type Follower struct {
	ID     int64
	Name   string
	Avatar *Avatar
}

type Avatar struct {
	URL string
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

// NewFollowersPage initializes result type FollowersPage from viewed result
// type FollowersPage.
func NewFollowersPage(vres *followingviews.FollowersPage) *FollowersPage {
	return newFollowersPage(vres.Projected)
}

// NewViewedFollowersPage initializes viewed result type FollowersPage from
// result type FollowersPage using the given view.
func NewViewedFollowersPage(res *FollowersPage, view string) *followingviews.FollowersPage {
	p := newFollowersPageView(res)
	return &followingviews.FollowersPage{Projected: p, View: "default"}
}

// newFollowersPage converts projected type FollowersPage to service type
// FollowersPage.
func newFollowersPage(vres *followingviews.FollowersPageView) *FollowersPage {
	res := &FollowersPage{}
	if vres.Total != nil {
		res.Total = *vres.Total
	}
	if vres.Page != nil {
		res.Page = *vres.Page
	}
	if vres.Followers != nil {
		res.Followers = newFollowerCollection(vres.Followers)
	}
	return res
}

// newFollowersPageView projects result type FollowersPage to projected type
// FollowersPageView using the "default" view.
func newFollowersPageView(res *FollowersPage) *followingviews.FollowersPageView {
	vres := &followingviews.FollowersPageView{
		Total: &res.Total,
		Page:  &res.Page,
	}
	if res.Followers != nil {
		vres.Followers = newFollowerCollectionView(res.Followers)
	}
	return vres
}

// newFollowerCollection converts projected type FollowerCollection to service
// type FollowerCollection.
func newFollowerCollection(vres followingviews.FollowerCollectionView) FollowerCollection {
	res := make(FollowerCollection, len(vres))
	for i, n := range vres {
		res[i] = newFollower(n)
	}
	return res
}

// newFollowerCollectionView projects result type FollowerCollection to
// projected type FollowerCollectionView using the "default" view.
func newFollowerCollectionView(res FollowerCollection) followingviews.FollowerCollectionView {
	vres := make(followingviews.FollowerCollectionView, len(res))
	for i, n := range res {
		vres[i] = newFollowerView(n)
	}
	return vres
}

// newFollower converts projected type Follower to service type Follower.
func newFollower(vres *followingviews.FollowerView) *Follower {
	res := &Follower{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Avatar != nil {
		res.Avatar = transformFollowingviewsAvatarViewToAvatar(vres.Avatar)
	}
	return res
}

// newFollowerView projects result type Follower to projected type FollowerView
// using the "default" view.
func newFollowerView(res *Follower) *followingviews.FollowerView {
	vres := &followingviews.FollowerView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	if res.Avatar != nil {
		vres.Avatar = transformAvatarToFollowingviewsAvatarView(res.Avatar)
	}
	return vres
}

// transformFollowingviewsAvatarViewToAvatar builds a value of type *Avatar
// from a value of type *followingviews.AvatarView.
func transformFollowingviewsAvatarViewToAvatar(v *followingviews.AvatarView) *Avatar {
	if v == nil {
		return nil
	}
	res := &Avatar{
		URL: *v.URL,
	}

	return res
}

// transformAvatarToFollowingviewsAvatarView builds a value of type
// *followingviews.AvatarView from a value of type *Avatar.
func transformAvatarToFollowingviewsAvatarView(v *Avatar) *followingviews.AvatarView {
	if v == nil {
		return nil
	}
	res := &followingviews.AvatarView{
		URL: &v.URL,
	}

	return res
}
