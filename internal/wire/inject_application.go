package wire

import (
	"github.com/boreq/velo/application"
	"github.com/boreq/velo/application/auth"
	"github.com/boreq/velo/application/tracker"
	"github.com/google/wire"
)

//lint:ignore U1000 because
var appSet = wire.NewSet(
	wire.Struct(new(application.Application), "*"),
	authSet,
	trackerSet,
)

//lint:ignore U1000 because
var authSet = wire.NewSet(
	wire.Struct(new(auth.Auth), "*"),
	auth.NewRegisterInitialHandler,
	auth.NewLoginHandler,
	auth.NewLogoutHandler,
	auth.NewCheckAccessTokenHandler,
	auth.NewListHandler,
	auth.NewCreateInvitationHandler,
	auth.NewRegisterHandler,
	auth.NewRemoveHandler,
	auth.NewSetPasswordHandler,
	auth.NewGetUserHandler,
	auth.NewUpdateProfileHandler,
	auth.NewChangePasswordHandler,

	authRepositoriesSet,
	authTransactableRepositoriesSet,
)

//lint:ignore U1000 because
var trackerSet = wire.NewSet(
	wire.Struct(new(tracker.Tracker), "*"),
	tracker.NewAddActivityHandler,
	tracker.NewGetActivityHandler,
	tracker.NewEditActivityHandler,
	tracker.NewDeleteActivityHandler,
	tracker.NewListUserActivitiesHandler,
	tracker.NewAddPrivacyZoneHandler,
	tracker.NewGetPrivacyZoneHandler,
	tracker.NewListUserPrivacyZonesHandler,
	tracker.NewDeletePrivacyZoneHandler,
	tracker.NewImportStravaHandler,
)
