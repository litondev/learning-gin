// CUSTOME MIDDLEWARE
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// EXAMPLE USING MIDDLEWARE
// func main() {
// 	// Creates a router without any middleware by default
// 	r := gin.New()

// 	// Global middleware
// 	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
// 	// By default gin.DefaultWriter = os.Stdout
// 	r.Use(gin.Logger())

// 	// Recovery middleware recovers from any panics and writes a 500 if there was one.
// 	r.Use(gin.Recovery())

// 	// Per route middleware, you can add as many as you desire.
// 	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

// 	// Authorization group
// 	// authorized := r.Group("/", AuthRequired())
// 	// exactly the same as:
// 	authorized := r.Group("/")
// 	// per group middleware! in this case we use the custom created
// 	// AuthRequired() middleware just in the "authorized" group.
// 	authorized.Use(AuthRequired())
// 	{
// 		authorized.POST("/login", loginEndpoint)
// 		authorized.POST("/submit", submitEndpoint)
// 		authorized.POST("/read", readEndpoint)

// 		// nested group
// 		testing := authorized.Group("testing")
// 		testing.GET("/analytics", analyticsEndpoint)
// 	}

// 	// Listen and serve on 0.0.0.0:8080
// 	r.Run(":8080")
// }


// AUTH MIDLEWARE
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}


// Goroutines inside a middleware
func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}