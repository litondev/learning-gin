package main

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

  dsn := "root:@tcp(127.0.0.1:3306)/laravel?charset=utf8mb4&parseTime=True&loc=Local"

  _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  fmt.Println("Success connect")
}