package controller
import (
	"github.com/gin-gonic/gin"
	"code.isstream.com/stream/auth"
)

func RegisterHandlers(app *gin.Engine) {

	authHandler := auth.Auth.MiddlewareFunc()

	signupController := &SignupController{}
	loginController := &LoginController{}
	userController := &UserController{}
	verifyController := &VerifyController{}

	v1 := app.Group("/v1")
	signupGroup := v1.Group("/signup")
	signupGroup.POST("/mobile", signupController.SignupWithMobile)
	signupGroup.POST("/email", signupController.SignupWithEmail)

	loginGroup := v1.Group("/login")
	loginGroup.POST("", auth.Auth.LoginHandler)
	//loginGroup.POST("/oauth", loginController.OauthLogin)

	v1.DELETE("/logout", loginController.Logout)

	v1.PUT("/nickname", authHandler, userController.UpdateNickname)
	v1.PUT("/avatar", userController.SaveAvatar)
	v1.PUT("/username", userController.UpdateUsername)
	v1.PUT("/password", userController.UpdatePassword)
	v1.POST("/password/reset/code", userController.SendMobileVerifyCodeForResetPassword)
	v1.PUT("/password/reset", userController.ResetPassword)

	v1.POST("/password/reset/verifycode", verifyController.VerifyMobileVerifyCode)
	v1.GET("/mobile/verifycode", verifyController.SendMobileVerifyCode)
	v1.POST("/mobile/verifycode", verifyController.VerifyCode)
	v1.POST("/mobile/new_verifycode", verifyController.VerifyMobileVerifyCode)
}
