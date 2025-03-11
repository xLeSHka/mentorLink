package app

import (
	groupService "github.com/xLeSHka/mentorLink/internal/service/group"
	mentorService "github.com/xLeSHka/mentorLink/internal/service/mentor"
	userService "github.com/xLeSHka/mentorLink/internal/service/user"

	"go.uber.org/fx"
)

var Services = fx.Provide(
	userService.New,
	mentorService.New,
	groupService.New,
)
