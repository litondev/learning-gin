package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") 

		c.JSON(200, gin.H{
			"message" : "Hai",
			"firstname" : firstname,
			"lastname" : lastname,
		})
	})

	r.Run()
}