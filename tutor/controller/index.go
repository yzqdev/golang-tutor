package controller

import (
	"github.com/gin-gonic/gin"
	"gotutor/util"
)

func Index(c *gin.Context) {
	util.JSON(c, 200, "success", "hhhhh")
}
