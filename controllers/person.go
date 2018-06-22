package controllers

import (
	"github.com/gin-gonic/gin"
	. "easy-gin/database"
	"fmt"
)

func IndexApi(c *gin.Context) {
	res,err := DB.Table("person").First()
	if err!=nil{
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data":res,
	})
}

