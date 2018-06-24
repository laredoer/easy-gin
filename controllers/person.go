package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	. "easy-gin/database"
)

func IndexApi(c *gin.Context) {
	res,err := DB.Table("person").Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data":res,
	})
}

