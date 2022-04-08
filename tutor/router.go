package main

import (
	"github.com/gin-gonic/gin"
	"gotutor/controller"
	"gotutor/libs"
	"gotutor/tips/mail"
)

func InitRouter(e *gin.Engine) {
	baseRouter := e.Group("/base")
	{
		baseRouter.POST("/login", mail.SendEmail)
		baseRouter.POST("/process", libs.ProcessBar)

	}
	adminRouter := e.Group("/admin")
	{

		adminRouter.GET("/index", controller.Index)

	}
	homeRouter := e.Group("/default")
	{
		homeRouter.GET("/index")

	}

}
