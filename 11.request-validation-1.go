package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`

	// CheckIn time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	// CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`

	// gtfield -> check out harus lebih dari pada check in
}

func main(){
	router := gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate",func(fl validator.FieldLevel) bool {
			fmt.Println(fl.Field())
				// string

			fmt.Println(fl.Field().Interface())
				// interface

		 	var date,ok = fl.Field().Interface().(time.Time)

		 	// jika tidak ada error
		 	if ok {
		 		// jika date lebih dari hari ini maka false
				if time.Now().After(date) {
					// false maka akan lolos
					return false
				}
			}

			// true maka akan di kembalikan
			return true
		})
	}

	router.GET("/bookable",func(c *gin.Context){
		var b Booking

		if err := c.ShouldBindQuery(&b); err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run()
}