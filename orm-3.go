package main

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
  "time"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

  dsn := "root:@tcp(127.0.0.1:3306)/learning_gin?charset=utf8mb4&parseTime=True&loc=Local"

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  fmt.Println("Success connect")

  // tags-tags

  // column
    // `gorm:"column=kelas_kita;not null;primaryKey"`

  // type
    // `gorm:"type=string"`

  // size
    // `gorm:"size=250"`

  // unique
   // `gorm:"unique"`

  // default
   // `gorm:"default=test"`

  // not null

  // autoIncrement

  // index

  // uniqueIndex

  // check

  type User struct {
    ID           uint `gorm:"primaryKey;autoIncrement"`
    Name         string `gorm:"size:100"`
    Email        string `gorm:"unique;size:100"`
    Age          int `gorm:"type:integer;size:11"`
    Married      bool
    Role         string `gorm:"type:enum('testa','testb')"`
    Birthday     string `gorm:"type:timestamp"`
    MemberNumber string `gorm:"not null"`
    ActivatedAt time.Time `gorm:"default=null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
  }

  // Migrate the schema
  db.AutoMigrate(&User{})
}