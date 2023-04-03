package main

import (
	"os"
	"strconv"
	"time"

	"github.com/amagimedia/nordend/internal/database"
	"github.com/amagimedia/nordend/internal/database/schema"
	"github.com/amagimedia/nordend/internal/handler"
	"github.com/amagimedia/nordend/internal/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	app             *cli.App
	port            int
	lggr            *zap.SugaredLogger
	numberOfWorkers int
)

func main() {
	lggr = logger.NewLogger("NORDEND", nil, "webserver")
	app = cli.NewApp()
	_ = database.Connect()

	lggr.Infow("Starting migration of DB", "activity", "DB_MIGRATION", "event", "MIGRATION_BEGINS")
	migrate()
	lggr.Infow("Completed migration of DB", "activity", "DB_MIGRATION", "event", "MIGRATION_COMPLETES")

	app.Name = "Nordend"
	app.Usage = "Webserver"
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " Amagi Media Labs"
	app.Compiled = time.Now()
	app.HideVersion = false
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Required:    false,
			Destination: &port,
			EnvVars:     []string{"WEBPORT"},
			Usage:       "port at which the API server needs to be deployed",
		},
		&cli.IntFlag{
			Name:        "workers",
			Aliases:     []string{"w"},
			Required:    false,
			Destination: &numberOfWorkers,
			EnvVars:     []string{"WORKERS"},
			Usage:       "number of workers",
		},
	}

	app.EnableBashCompletion = true

	app.Action = func(c *cli.Context) error {
		handler.Lggr = lggr
		return handler.Handler(strconv.Itoa(port))
	}

	err := app.Run(os.Args)
	if err != nil {
		lggr.Fatalw("Failed to launch the application : "+err.Error(), "activity", "LAUNCH", "event", "ERROR_LAUNCHING_APPLICATION")
	}

}

func migrate() {
	_ = database.Database.AutoMigrate(&schema.Account{})
	_ = database.Database.AutoMigrate(&schema.Channel{})
	_ = database.Database.AutoMigrate(&schema.Audio{})
	_ = database.Database.AutoMigrate(&schema.Artwork{})
	_ = database.Database.AutoMigrate(&schema.Caption{})
	_ = database.Database.AutoMigrate(&schema.Rating{})
}
