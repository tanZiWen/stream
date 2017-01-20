package controller

import (
	"strconv"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"code.isstream.com/stream/db"
	"code.isstream.com/stream/lib"
	"code.isstream.com/stream/s-user/userservice"
	"code.isstream.com/stream/s-user/model/bizerror"
)

type VerifyController struct {
	lib.BaseController
}


func (ctr *VerifyController) SendMobileVerifyCode(c *gin.Context) {
	mobile := c.Query("mobile")
	isValid := mobile != ""
	if !isValid {
		ctr.Fail(c, bizerror.MOBILE_INVALID)
		return
	}

	verifyCodeKey := GetVerifyCodeKey(c)
	log.Debug("verifyCodeKey key:%s", verifyCodeKey)

	db.Redis.Set(verifyCodeKey, MOCKUP_MOBILE_VERIFY_CODE)
	db.Redis.Expire(verifyCodeKey, 600)
	ctr.Success(c)
}

func GetVerifyCodeKey(c *gin.Context) string {
	//TODO use new key strategy
	//uid, _ := c.Get(auth.KEY_USERID)
	//return uid + "_vmc"

	return ""
}

func (ctr *VerifyController) VerifyCode(c *gin.Context) {
	verifyCode := c.PostForm("verifyCode")

	isValid := verifyCode != ""
	if !isValid {
		ctr.Fail(c, bizerror.VERIFYCODE_INVALID)
		return
	}

	verifyCodeKey := GetVerifyCodeKey(c)

	verifyData := db.Redis.Get(verifyCodeKey).Val()

	log.Debug("verifyCode: %s", verifyCode, "verifyCodeKey: %s", verifyCodeKey)

	if verifyData != verifyCode {
		ctr.Fail(c, bizerror.VERIFYCODE_INVALID)
		return
	}
	ctr.Success(c)
}

func (ctr *VerifyController) VerifyMobileVerifyCode(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form SmsForm

	err = c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	mobile, verifyCode := GetMobileAndCode(c)

	log.Debug("verify mobile %s with code %s", mobile, verifyCode)

	if strconv.FormatInt(form.Mobile, 10) != mobile || form.VerifyCode != verifyCode {
		ctr.Fail(c, bizerror.MOBILE_VERIFY_INVALIDCODE)
		return
	}

	err = userservice.UpdateMobile(form.Mobile, uid); if err != nil {
		log.Error("update user mobile error: %s", err)
		panic(err)
	}

	savedUser, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user:", err)
		panic(err)
	}

	ctr.SuccessPair(c, "user", savedUser)
}

func GetMobileAndCode(c *gin.Context) (interface{}, interface{}) {
	verifycodeKey := GetVerifycodeKey(c)

	log.Debug("mobile verify data key in redis is %s", verifycodeKey)
	verifyData, err := db.Redis.HMGet(verifycodeKey, "mb", "vc").Result(); if err != nil {
		log.Error(3, "get redis verifycodekey error: %v", err)
		panic(err)
	}

	log.Debug("mobile and verify code, %v", verifyData)

	return verifyData[0], verifyData[1]
}