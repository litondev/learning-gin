package main

import "github.com/gin-gonic/gin"

type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main(){
	router := gin.Default()

	router.POST("loginJSON",func(c *gin.Context){
		var json Login

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(422,gin.H{"error" : err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "12345" {
			c.JSON(422,gin.H{"error" : "unauto"})
			return 
		}

		c.JSON(200,gin.H{"login" : "login"})
	})

	router.POST("loginForm",func(c *gin.Context){
		var form Login

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(422,gin.H{"error" : err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "12345" {
			c.JSON(422,gin.H{"error" : "unauto"})
			return 
		}

		c.JSON(200,gin.H{"login" : "login"})
	})

	router.Run()
}
