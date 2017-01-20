package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/setting"
	"code.isstream.com/stream/db"
	"code.isstream.com/stream/partner-service/controllers"
	"github.com/gin-gonic/gin"

	//log "github.com/Sirupsen/logrus"
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

func main() {
	initializeModules()

	log.Debug("finish initializing modules")

	app := gin.New()
	app.Use(middleware.AccessControl())
	version := fmt.Sprintf("/%s", setting.App.Version)
	verGroup := app.Group(version,
		gin.Recovery(),
		middleware.ErrorHandler(),
	)
	controllers.RegisterHandlers(verGroup)

	app.Run(setting.App.ListenPort)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}