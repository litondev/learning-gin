package main

import "github.com/gin-gonic/gin"
import "fmt"

func main(){
	router := gin.Default()

	// /post?ids[a]=1234&ids[b]=hello
	// names[first]=thinkerou&names[second]=tianou

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
		// ids: map[b:hello a:1234]; 
		// names: map[second:tianou first:thinkerou]
	})

	router.Run()
}