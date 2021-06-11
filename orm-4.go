package main

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"  
)

type Latihan struct {
	ID		uint 	`gorm:"primaryKey;autoIncrement"`
	Name	string	`gorm:"type=string;size:50"`
	CreatedAt     string `gorm:"type=timestamp;default=null"`
	UpdatedAt	  string `gorm:"type=timestamp"`
}


// hook
// func (l *Latihan) BeforeCreate(tx *gorm.DB) (err error) {
//   	l.Name = l.Name + "testing"
//   	return
// }

func main(){
  dsn := "root:@tcp(127.0.0.1:3306)/learning_gin?charset=utf8mb4&parseTime=True&loc=Local"

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  fmt.Println("Success connect")


  // db.AutoMigrate(&Latihan{})

  /*
  // insert
  latihan := Latihan{Name: "litoninot"}

  result := db.Create(&latihan)

  fmt.Println(latihan.ID)
  fmt.Println(result.Error)
  fmt.Println(result.RowsAffected)
  */	

  /*
  // insert batch
  var users = []Latihan{{Name: "jinzhu_1"},{Name: "jinzhu_10000"}}

  // batch size 100
  db.CreateInBatches(users, 100)
  */

  /*
  // hook example
  // latihan := Latihan{Name: "litoninot"}
  // db.Create(&latihan)
  */

  // 
  // latihan := Latihan{Name: "jkj"}
  // db.Create(&latihan)


	// type CreditCard struct {
	//   gorm.Model
	//   Number   string
	//   UserID   uint
	// }

	// type User struct {
	//   gorm.Model
	//   Name       string
	//   CreditCard CreditCard
	// }

	// db.Create(&User{
	//   Name: "jinzhu",
	//   CreditCard: CreditCard{Number: "411111111111"}
	// })
}