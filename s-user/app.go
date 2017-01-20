package main

import (
	"code.isstream.com/stream/s-user/controller"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/setting"
	"code.isstream.com/stream/db"
	"github.com/gin-gonic/gin"

	"code.isstream.com/stream/middleware"
	"code.isstream.com/stream/idg"
	"code.isstream.com/stream/auth"
	"os"
)

func initializeModules() {
	setting.Initialize()
	log.Debug("global setting initialized")
	db.Initialize()
	log.Debug("db setting initialized")
	idg.Initialize()
	log.Debug("idg setting initialized")
	auth.Initialize()
	log.Debug("auth setting initialized")
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	log.SetLevel(log.DebugLevel)
}

func main() {

	initializeModules()

	app := gin.New()
	app.Use(middleware.AccessControl())
	app.Use(
		gin.Recovery(),
		middleware.ErrorHandler(),
	)
	controller.RegisterHandlers(app)

	app.Run(setting.App.ListenPort)
}
