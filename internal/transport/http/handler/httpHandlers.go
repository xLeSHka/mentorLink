package httpHandlers

import (
	"github.com/xLeSHka/mentorLink/internal/app/Validators"
	"github.com/xLeSHka/mentorLink/internal/transport/http/handler/ApiRouters"
	groupsRoute "github.com/xLeSHka/mentorLink/internal/transport/http/handler/group"
	mentorsRoute "github.com/xLeSHka/mentorLink/internal/transport/http/handler/mentor"
	publicRoute "github.com/xLeSHka/mentorLink/internal/transport/http/handler/public"
	usersRoute "github.com/xLeSHka/mentorLink/internal/transport/http/handler/user"
	"github.com/xLeSHka/mentorLink/internal/transport/http/handler/ws"
	"go.uber.org/fx"
)

var HttpHandlers = fx.Module("httpHandlers",
	fx.Provide(
		ApiRouters.CreateApiRoutes,
		Validators.New,
		ws.New,
		fx.Private),
	//publicRoute.PublicRoute,
	fx.Invoke(
		publicRoute.PublicRoute,
		usersRoute.UsersRoute,
		mentorsRoute.MentorsRoute,
		groupsRoute.GroupsRoutes,
	),
)
