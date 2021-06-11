package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()

	router.POST("/",func(c *gin.Context){
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick","nickname")

		c.JSON(200,gin.H{
			"status" : "posted",
			"message" : message,
			"nick" : nick,
		})
	})

	router.Run()
}