package controllers

import (
	"github.com/gin-gonic/gin"
	//"code.isstream.com/stream/m-user"
	"code.isstream.com/stream/auth"
)

func RegisterHandlers(group *gin.RouterGroup) {

	authHandler := auth.Auth.MiddlewareFunc()
	//userController := &m_user.UserController{}
	customerController := &CustomerController{}
	itemController := &ItemController{}
	orderController := &OrderController{}

	loginGroup := group.Group("/login")
	loginGroup.POST("", auth.Auth.LoginHandler)

	//userGroup := group.Group("/users")
	//userGroup.PUT("/nickname", authHandler, userController.UpdateNickname)
	//userGroup.GET("/:id", userController.GetUserBrief)
	//userGroup.POST("/feedback", feedbackController.Feedback)

	myGroup := group.Group("/my")
	myGroup.GET("/customers", authHandler, customerController.GetMyCustomers)
	myGroup.POST("/customers/mobile", authHandler, customerController.CheckMobile)
	myGroup.POST("/customers", authHandler, customerController.CreateMyCustomers)
	myGroup.PUT("/customers/:id", authHandler, customerController.UpdateMyCustomers)
	myGroup.GET("/customers/orders", authHandler, orderController.GetMyCustomerOrders)

	customerGroup := group.Group("/customers")
	customerGroup.POST("/:id/orders", authHandler, orderController.CreateOrderForCustomer)
	customerGroup.GET("/:id", authHandler, customerController.GetCustomer)
	customerGroup.GET("/:id/order", authHandler,orderController.GetOrderDetail)
	customerGroup.GET("/:id/orders", authHandler,orderController.GetOrdersByCustomerId)

	//orderGroup := group.Group("/orders")
	//orderGroup.GET("/:id", authHandler, orderController.GetOrderDetail)

	itemGroup := group.Group("/items")
	itemGroup.GET("", itemController.GetItems)
}
