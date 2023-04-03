package handler

import (
	"strings"

	v1 "github.com/amagimedia/nordend/internal/handler/v1"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

var Lggr *zap.SugaredLogger

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

type APIs []struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}

func Handler(port string) error {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodHead,
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
		}, ","),
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	prometheus := fiberprometheus.New("nordend")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	apiV1 := app.Group("/api/v1")
	initialiseV1Vars()
	AddToHandler(describeV1APIs(), apiV1)

	return app.Listen(":" + port)
}

func AddToHandler(APIs APIs, app fiber.Router) {
	for _, api := range APIs {
		switch api.Method {
		case GET:
			app.Get(api.Path, api.Handler)
		case POST:
			app.Post(api.Path, api.Handler)
		case PUT:
			app.Put(api.Path, api.Handler)
		case PATCH:
			app.Patch(api.Path, api.Handler)
		case DELETE:
			app.Delete(api.Path, api.Handler)
		default:
			app.Add(api.Method, api.Path, api.Handler)
		}
	}
}

func initialiseV1Vars() {
	v1.Lggr = Lggr
}

func describeV1APIs() APIs {
	return APIs{
		{
			Method:  GET,
			Path:    "/health",
			Handler: v1.Health,
		},
		{
			Method:  POST,
			Path:    "/accounts",
			Handler: v1.AddAccount,
		},
		{
			Method:  GET,
			Path:    "/accounts",
			Handler: v1.ListAccounts,
		},
		{
			Method:  GET,
			Path:    "/accounts/:amgid",
			Handler: v1.GetAccount,
		},
		{
			Method:  DELETE,
			Path:    "/accounts/:amgid",
			Handler: v1.DeleteAccount,
		},
		{
			Method:  POST,
			Path:    "/channels",
			Handler: v1.AddChannel,
		},
		{
			Method:  GET,
			Path:    "/channels",
			Handler: v1.ListChannels,
		},
		{
			Method:  GET,
			Path:    "/channels/:channelid",
			Handler: v1.FetchChannel,
		},
		{
			Method:  DELETE,
			Path:    "/channels/:channelid",
			Handler: v1.DeleteChannel,
		},
		{
			Method:  PUT,
			Path:    "/channels/:channelid/captions",
			Handler: v1.UpdateCaptions,
		},
		{
			Method:  PUT,
			Path:    "/channels/:channelid/audios",
			Handler: v1.UpdateAudios,
		},
		{
			Method:  PUT,
			Path:    "/channels/:channelid/artworks",
			Handler: v1.UpdateArtworks,
		},
		{
			Method:  PUT,
			Path:    "/channels/:channelid/ratings",
			Handler: v1.UpdateRatings,
		},
	}
}
