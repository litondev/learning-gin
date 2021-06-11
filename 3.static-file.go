package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

    r.Static("/assets", "./assets")
    
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "static.html", gin.H{
			"title": "Main website",
		})
	})

	r.Run()
}