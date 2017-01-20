package controller

import (
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/db"
	"github.com/flosch/pongo2"
	"time"
	"github.com/gin-gonic/gin"
	"strconv"
	valid "github.com/asaskevich/govalidator"
	"strings"
	"code.isstream.com/stream/lib"
	"code.isstream.com/stream/idg"
	"code.isstream.com/stream/setting"
	"code.isstream.com/stream/s-user/model/bizerror"
	"code.isstream.com/stream/s-user/userservice"
	"code.isstream.com/stream/domain"
	"code.isstream.com/stream/s-user/usermodel"
)

type SignupForm struct {
	UserId   string `form:"username" json:"username"`
	NickName string `form:"nickname" json:"nickname"`
	Mobile   int64 `form:"mobile" json:"mobile"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Source   int    `form:"source" json:"source"`
}

type SignupController struct {
	lib.BaseController
}

func (ctr *SignupController) SendMobileVerifyCode(c *gin.Context) {
	mobileStr := c.Query("mobile")
	isValid := mobileStr != ""
	if !isValid {
		log.Info("bad request")
		ctr.BadRequest(c)
		return
	}

	log.Debug("send verify code to mobile %d", mobileStr)

	mobile, err := lib.Str2int64(mobileStr); if err != nil {
		log.Error(3, "failde to convert string to int64, %v", err)
		panic(err)
	}

	has, _, err := userservice.GetUserByMobile(mobile); if err != nil {
		log.Error(0, "failed to find user by mobile, %v", err)
		panic(err)
	}

	if has {
		ctr.Fail(c, bizerror.SIGNUP_MOBILE_CLAIMED)
		return
	}

	verifycodeKey := GetVerifycodeKey(c)
	//	val := utils.RandomSpecStr(6, utils.INT_CHARSET)
	//	content := getVeriCodeText(val)
	//	utils.SendSms(form.Mobile, content)

	//add tracking to stop client keeping requesting verify code
	db.Redis.HMSet(verifycodeKey, "mb", mobileStr, "vc", MOCKUP_MOBILE_VERIFY_CODE)
	db.Redis.Expire(verifycodeKey, 600 * time.Second)

	ctr.Success(c)
}

func GetVerifycodeKey(c *gin.Context) string {
	return ""
	//return c.SessionStore.SessionID() + "_mvc"
}


func (ctr *SignupController) VerifyMobileVerifyCode(c *gin.Context) {
	var form SmsForm
	err := c.Bind(&form); if err != nil {
		log.Error(3, "bad request ", err)
		log.Info("bad request")
		ctr.BadRequest(c)
		return
	}

	mobile, code := GetMobileAndCode(c)

	log.Debug("verify mobile %s with code %s", mobile, code)
	log.Debug("validate: %v", form.Mobile != mobile)

	if strconv.FormatInt(form.Mobile, 10) != mobile || form.VerifyCode != code {
		ctr.Fail(c, bizerror.MOBILE_VERIFY_INVALIDCODE)
		return
	}

	ctr.Success(c)
}

func (ctr *SignupController) SignupWithMobile(c *gin.Context) {
	var form SignupForm

	err := c.Bind(&form); if err != nil {
		log.Info("bad request")
		ctr.BadRequest(c)
		return
	}

	mobile, _ := GetMobileAndCode(c)

	if mobile == nil {
		ctr.Fail(c, bizerror.SIGNUP_MOBILE_INCONSISTENT)
		return
	}

	user := usermodel.User{}
	log.Debug("nickname: ", form.NickName)
	user.Nickname = form.NickName
	user.Password = form.Password

	user.Mobile, err = lib.Str2int64(mobile.(string))

	if err != nil {
		log.Info("bad request")
		ctr.BadRequest(c)
		return
	}

	ctr.doSignup(&user, c)
}


func getVeriCodeText(code string) string {
	tpl, err := pongo2.FromString(setting.Config.Section("sms.tpl").Key("VERI_CODE").String())
	if err != nil {
		panic(err)
	}
	content, err := tpl.Execute(pongo2.Context{"veriCode": code})
	if err != nil {
		panic(err)
	}
	return content
}

//type EmailSignupForm struct {
//	UserId   string `json:"userid"`
//	NickName string `form:"nickname" json:"nickname"`
//	Email    string `form:"email" json:"email"`
//	Password string `form:"password" json:"password"`
//	Source   int    `form:"source" json:"source"`
//}

func (ctr *SignupController) SignupWithEmail(c *gin.Context) {
	var form SignupForm

	err := c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	if form.Email == "" || form.Password == "" {
		log.Info("bad request")
		ctr.BadRequest(c)
		return
	}

	email := strings.ToLower(form.Email)
	if !valid.IsEmail(email) {
		ctr.Fail(c, bizerror.EMAIL_INVALID)
		return
	}

	has, _, err := userservice.GetUserByEmail(email); if err != nil {
		log.Error("failed to find user by email", err)
		panic(err)
	}

	if has {
		ctr.Fail(c, bizerror.SIGNUP_EMAIL_CLAIMED)
		return
	}

	user := usermodel.User{}
	log.Debug("nickname: ", form.NickName)
	user.Nickname = form.NickName
	user.Email = email
	user.Password = form.Password

	ctr.doSignup(&user, c)
}


func (ctr *SignupController) doSignup(user *usermodel.User, c *gin.Context) {
	userId, err := idg.Id(); if err != nil {
		return
	}

	user.Id = userId
	user.Status = domain.DB_COMMON_STATUS_OK

	encryptedPassword, err := lib.EncryptPassword(user.Password); if err != nil {
		log.Error("failed to encrypt password", err)
		panic(err)
	}

	user.Password = encryptedPassword
	err = userservice.CreateUser(user); if err != nil {
		log.WithField("user", user).Error("failed to create user", err)
		panic(err)
	}

	log.Debug("user :", user)

	userForm, err := userservice.GetUser(userId); if err != nil {
		log.WithField("id", userId).Error("failed to get user", err)
		panic(err)
	}

	//if err = streamservice.InitUserStreamsByTemplate(user.Id, "base_v1"); err != nil {
	//	log.Debug("failed to do initialization after user signup ", err)
	//	c.Failed(bizerror.SIGNUP_INIT_FAILED)
	//	return
	//}

	ctr.SuccessPair(c, "user", userForm)
}

