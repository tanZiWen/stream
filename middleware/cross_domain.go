package middleware

import (
	"strings"
	"github.com/gin-gonic/gin"
	"code.isstream.com/stream/setting"
	log "github.com/Sirupsen/logrus"
)

func accept(c *gin.Context) bool {
	origin := c.Request.Header.Get("origin")
	from := c.Request.Header.Get("From")

	return strings.HasSuffix(origin, setting.App.AllowDomain) || from == setting.App.AllowFrom
}

func AccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if !accept(c) {
		//	log.Warn("request from unsupport source: ", c.Request.Header.Get("origin"), c.Request.Header.Get("from"))
		//	c.AbortWithStatus(400)
		//	return
		//}
		method := c.Request.Method
		origin := c.Request.Header.Get("origin")

		if strings.ToUpper(method) == "OPTIONS" {
			log.Debug("options request from ", origin)
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With")
			c.Writer.Header().Set("Access-Control-Max-Age", "1728000")
			c.AbortWithStatus(204)
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
		return
	}
}
