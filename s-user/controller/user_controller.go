package controller
import (
	"github.com/gin-gonic/gin"
	"strconv"
	log "github.com/Sirupsen/logrus"
	"strings"
	"code.isstream.com/stream/db"
	"time"
	"code.isstream.com/stream/lib"
	"code.isstream.com/stream/s-user/userservice"
	"code.isstream.com/stream/s-user/model/bizerror"
)

type userForm struct {
	UserName    string     `form:"username" json:"username"`
	Nickname    string     `form:"nickname" json:"nickname"`
	Email       string     `form:"email" json:"email"`
	Password    string     `form:"password" json:"password"`
	NewPassword string       `form:"newpassword" json:"newpassword"`
	Mobile      int64       `form:"mobile" json:"mobile"`
	VerifyCode  string       `form:"verifycode" json:"verifycode"`
}

type SmsForm struct {
	Mobile     int64 `form:"mobile" json:"mobile,string"`
	VerifyCode string `form:"verifycode" json:"verifycode"`
}

const (
	MOCKUP_MOBILE_VERIFY_CODE = "123456"
)

type UserController struct {
	lib.BaseController
}

func (ctr *UserController) GetUserBrief(c *gin.Context) {
	idstr := c.Param("id")
	log.Debug("get user brief with id ", idstr)

	//TODO security warning!!! validate privilege here

	id, err := strconv.ParseInt(idstr, 10, 64); if err != nil {
		ctr.Fail(c, bizerror.ID_INVALID)
		return
	}

	user, err := userservice.GetUser(id); if err != nil {
		log.WithField("id", id).Error("failed to get user", err)
		panic(err)
	}

	if user == nil {
		ctr.Fail(c, bizerror.USER_NOTEXIST)
		return
	}

	ctr.SuccessPair(c, "user", user)
}

func (ctr *UserController) MyBrief(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	log.Debug("get user brief info with id", uid);

	user, err := userservice.GetUserBriefEntityById(uid)

	if err != nil {
		panic(err)
	}

	if user == nil {
		ctr.Fail(c, bizerror.USER_NOTEXIST)
		return
	}

	ctr.SuccessPair(c, "user", user)
}

func (ctr *UserController) UpdateNickname(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form userForm
	err = c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	nickname := strings.TrimSpace(form.Nickname)
	if len(nickname) <= 0 {
		ctr.Fail(c, bizerror.USER_NICKNAME_INVALID)
		return
	}

	log.Debug("get user by id", uid)

	user, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user", err)
		panic(err)
	}

	if user == nil {
		ctr.Fail(c, bizerror.USER_NOTEXIST)
		return
	}

	err = userservice.UpdateNickname(nickname, uid); if err != nil {
		log.WithField("nickname", nickname).WithField("id", uid).Error("failed to update nickname", err)
		panic(err)
	}

	savedUser, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user", err)
		panic(err)
	}

	ctr.SuccessPair(c, "user", savedUser)
}

func (ctr *UserController) UpdateUsername(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form userForm
	err = c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	username := strings.TrimSpace(form.UserName)
	if len(username) <= 0 {
		ctr.Fail(c, bizerror.USER_USERNAME_INVALID)
		return
	}

	log.Debug("get user by id", uid)
	user, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user", err)
		panic(err)
	}

	if user == nil {
		ctr.Fail(c, bizerror.USER_NOTEXIST)
		return
	}


	if len(user.UserName) > 0 {
		ctr.Fail(c, bizerror.USER_USERNAME_IS_EXIST)
		return
	}

	err = userservice.UpdateUsername(username, uid); if err != nil {
		log.WithField("username", username).WithField("id", uid).Error("failed to update username", err)
		panic(err)
	}

	savedUser, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user", err)
		panic(err)
	}

	ctr.SuccessPair(c, "user", savedUser)
}

type PasswordForm struct {
	OldPassword string `form:"old" json:"old"`
	NewPassword string `form:"new" json:"new"`
}

func (ctr *UserController) UpdatePassword(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form PasswordForm
	err = c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	password := strings.TrimSpace(form.OldPassword)
	newPassword := strings.TrimSpace(form.NewPassword)

	if len(password) <= 0 || len(newPassword) <= 0 {
		ctr.Fail(c, bizerror.AUTH_PASSWORD_INVALID)
		return
	}

	log.Debug("get user by id", uid)
	user, err := userservice.GetUser(uid); if err != nil {
		log.WithField("id", uid).Error("failed to get user", err)
		panic(err)
	}

	if user == nil {
		ctr.Fail(c, bizerror.USER_NOTEXIST)
		return
	}

	pwdParts := strings.Split(user.Password, "_")
	log.Debug("saved password ", user.Password)

	encryptedPassword, err := lib.EncryptPassword(password, pwdParts[2]); if err != nil {
		log.Error("failed to encrypt password", err)
		panic(err)
	}
	log.Debug("submit password ", encryptedPassword)
	if encryptedPassword != user.Password {
		ctr.Fail(c, bizerror.AUTH_PASSWORD_INVALID)
		return
	}

	newEncyrptedPassword, err := lib.EncryptPassword(form.NewPassword); if err != nil {
		log.Error("failed to encrypt password", err)
		panic(err)
	}
	log.Debug("new password ", newEncyrptedPassword)

	err = userservice.UpdatePassword(newEncyrptedPassword, uid); if err != nil {
		log.Error("failed to update password", err)
		panic(err)
	}

	ctr.Success(c)
}

func (ctr *UserController) ResetPassword(c *gin.Context) {
	_, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form userForm
	err = c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	log.Debug("get user by mobile", form.Mobile)
	has, user, err := userservice.GetUserByMobile(form.Mobile); if err != nil || !has {
		log.WithField("mobile", form.Mobile).Error("failed to get user by mobile", err)
		panic(err)
	}

	newPassword := strings.TrimSpace(form.NewPassword)
	if len(newPassword) <= 0 {
		ctr.Fail(c, bizerror.AUTH_PASSWORD_INVALID)
		return
	}

	newEncyrptedPassword, err := lib.EncryptPassword(form.NewPassword); if err != nil {
		log.Error("failed to encrypt password", err)
		panic(err)
	}

	log.Debug("new password ", newEncyrptedPassword)
	err = userservice.UpdatePassword(newEncyrptedPassword, user.Id); if err != nil {
		log.Error("failed to update password", err)
		panic(err)
	}

	ctr.Success(c)
}

//TODO move to verify controller
func (ctr *UserController) SendMobileVerifyCodeForResetPassword(c *gin.Context) {
	var form userForm
	err := c.Bind(&form); if err != nil {
		log.Info("bad request", err)
		ctr.BadRequest(c)
		return
	}

	mobile := strconv.FormatInt(form.Mobile, 10)
	isValid := mobile != ""
	if !isValid {
		ctr.Fail(c, bizerror.MOBILE_INVALID)
		return
	}

	log.Debug("send verify code to mobile %s", mobile)

	has, _, err := userservice.GetUserByMobile(form.Mobile); if err != nil {
		log.WithField("mobile", mobile).Error("failed to find user by mobile", err)
		panic(err)
	}

	if !has {
		ctr.Fail(c, bizerror.MOBILE_NOTEXIST)
		return
	}
	verifycodeKey := GetVerifycodeKey(c)
	verifycode := lib.RandomSpecStr(6, lib.INT_CHARSET)

	//add tracking to stop client keeping requesting verify code
	db.Redis.HMSet(verifycodeKey, "mb", mobile, "vc", verifycode)
	db.Redis.Expire(verifycodeKey, 600 * time.Second)

	//yunpian.SendVerifyCode()
	//_, err = sms.SendVerifyCode(verifycode, mobile); if err != nil {
	//	log.WithField("mobile", mobile).Error("failed to send verify code", err)
	//	panic(err)
	//}

	ctr.Success(c)
}

func (ctr *UserController) SaveAvatar(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	//c.Request.ParseMultipartForm(32 << 20)
	//file, fileHeader, _ := c.Request.FormFile("image")
	//defer file.Close()
	//if file == nil {
	//	ctr.Fail(c, bizerror.USER_IMAGE_UPLOAD_EMPTY)
	//	return
	//}

	//err = userservice.SaveImage(uid, file, fileHeader)
	//
	//if err != nil {
	//	ctr.Fail(c, bizerror.USER_IMAGE_SAVE_FAILED)
	//	return
	//}
	savedUser, err := userservice.GetUser(uid)

	if err != nil {
		log.WithField("id", uid).Error("failed to get user:", err)
		panic(err)
	}

	ctr.SuccessPair(c, "user", savedUser)
}