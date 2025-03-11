package mentorsRoute

import (
	"github.com/xLeSHka/mentorLink/internal/app/Validators"
	"github.com/xLeSHka/mentorLink/internal/repository"
	"github.com/xLeSHka/mentorLink/internal/service"
	"github.com/xLeSHka/mentorLink/internal/transport/http/handler/ApiRouters"
	"github.com/xLeSHka/mentorLink/internal/transport/http/handler/ws"

	"go.uber.org/fx"
)

type Route struct {
	routers         *ApiRouters.ApiRouters
	validator       *Validators.Validator
	mentorService   service.MentorService
	userService     service.UserService
	minioRepository repository.MinioRepository
	wsconn          *ws.WebSocket
}

type FxOpts struct {
	fx.In
	ApiRouter       *ApiRouters.ApiRouters
	Validator       *Validators.Validator
	UsersService    service.UserService
	MentorService   service.MentorService
	MinioRepository repository.MinioRepository
	Ws              *ws.WebSocket
}

func MentorsRoute(opts FxOpts) *Route {
	router := &Route{
		routers:         opts.ApiRouter,
		validator:       opts.Validator,
		userService:     opts.UsersService,
		mentorService:   opts.MentorService,
		minioRepository: opts.MinioRepository,
		wsconn:          opts.Ws,
	}

	opts.ApiRouter.MentorRoute.GET("/mentors/students", router.students)
	opts.ApiRouter.MentorRoute.GET("/mentors/requests", router.getRequests)
	opts.ApiRouter.MentorRoute.POST("/mentors/requests", router.updateRequest)

	return router
}
