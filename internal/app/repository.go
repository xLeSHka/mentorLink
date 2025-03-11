package app

import (
	repositoryGroup "github.com/xLeSHka/mentorLink/internal/repository/group"
	repositoryMentor "github.com/xLeSHka/mentorLink/internal/repository/mentor"
	repositoryMinio "github.com/xLeSHka/mentorLink/internal/repository/minio"
	repositoryUser "github.com/xLeSHka/mentorLink/internal/repository/user"

	"go.uber.org/fx"
)

var Repositories = fx.Provide(
	repositoryUser.New,
	repositoryMentor.New,
	repositoryGroup.New,
	repositoryMinio.New,
)
