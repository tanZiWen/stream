package controller

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"code.isstream.com/stream/lib"
)

type LoginController struct {
	lib.BaseController
}

func (ctr *LoginController) Logout(c *gin.Context) {
	ctr.Success(c)
}


