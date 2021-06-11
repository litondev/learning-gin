package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(200, "Hello %s", name)
	// })

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(200, message)
	// })

	// For each matched request Context will hold the route definition
	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	c.String(200,c.FullPath())
	// })

	router.Run()
}