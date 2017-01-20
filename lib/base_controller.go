package lib

import (
	"github.com/gin-gonic/gin"
	_domain "code.isstream.com/stream/domain"
	"strconv"
	"net/http"
	"errors"
)

type BaseController struct {
}

func (ctr *BaseController) GetInt64(c *gin.Context, key string) (v int64, has bool, err error) {
	str, has := c.Get(key)

	if !has {
		return 0, false, nil
	}

	v, err = strconv.ParseInt(str.(string), 10, 64)

	return v, true, err
}

func (ctr *BaseController) GetInt64Value(c *gin.Context, key string) (v int64, err error) {
	v, has, err := ctr.GetInt64(c, key)

	if has || err != nil {
		return v, err
	}

	return 0, nil
}

func (ctr *BaseController) GetUid(c *gin.Context) (v int64, err error) {
	str, has := c.Get(_domain.KEY_USERID)

	if !has {
		return 0, errors.New("user id not present")
	}

	v, err = strconv.ParseInt(str.(string), 10, 64)

	return v, err
}

func (ctr *BaseController) Success(c *gin.Context) {
	ctr.SuccessData(c, nil)
}

func (ctr *BaseController) SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, _domain.Response{Code: "ok", Data: data})
}

func (ctr *BaseController) SuccessPair(c *gin.Context, key string, value interface{}) {
	c.JSON(http.StatusOK, _domain.Response{Code: "ok", Data: map[string]interface{}{key: value }})
}

func (ctr *BaseController) Fail(c *gin.Context, code string) {
	c.JSON(http.StatusOK, _domain.Response{Code: code})
	c.Abort()
}

func (ctr *BaseController) FailMessage(c *gin.Context, code string, message string) {
	c.JSON(http.StatusOK, _domain.Response{Code: code, Message: message})
	c.Abort()
}

func (ctr *BaseController) FailError(c *gin.Context, code string, err error) {
	c.JSON(http.StatusOK, _domain.Response{Code: code, Message: err.Error()})
	c.Abort()
}

func (ctr *BaseController) BadRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func (ctr *BaseController) Unauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (ctr *BaseController) NotFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}