package router

import (
	"github.com/gin-gonic/gin"
	. "easy-gin/controllers"
	)

var R *gin.Engine

func init()  {
	R = gin.Default()
	R.GET("/", IndexApi)
	R.POST("/person", AddPersonApi)

}
