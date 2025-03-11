package app

import (
	"github.com/xLeSHka/mentorLink/internal/connetions/db"
	"github.com/xLeSHka/mentorLink/internal/connetions/minio"
	"github.com/xLeSHka/mentorLink/internal/transport/http"
	httpHandlers "github.com/xLeSHka/mentorLink/internal/transport/http/handler"
	"github.com/xLeSHka/mentorLink/internal/transport/http/pkg/jwt"

	"go.uber.org/fx"
)

var App = fx.Options(
	fx.Provide(
		db.New,
		http.New,
		jwt.New,
		//redis.New,
		minio.New,
	),
	Repositories,
	Services,
	httpHandlers.HttpHandlers,
)
