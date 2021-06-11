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

func main(){
  dsn := "root:@tcp(127.0.0.1:3306)/learning_gin?charset=utf8mb4&parseTime=True&loc=Local"

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  fmt.Println("Success connect")

  // Email's ID is `10`
  latihan := Latihan{ID:3}
  db.Delete(&latihan)
  // DELETE from emails where id = 10;

  // Delete with additional conditions
  // db.Where("name = ?", "jinzhu").Delete(&email)
  // DELETE from emails where id = 10 AND name = "jinzhu";

  // db.Delete(&User{}, 10)
  // // DELETE FROM users WHERE id = 10;

  // db.Delete(&User{}, "10")
  // // DELETE FROM users WHERE id = 10;

  // db.Delete(&users, []int{1,2,3})
  // // DELETE FROM users WHERE id IN (1,2,3);

  // func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
  //   if u.Role == "admin" {
  //     return errors.New("admin user not allowed to delete")
  //   }
  //   return
  // }

 // db.Delete(&User{}).Error // gorm.ErrMissingWhereClause

  // db.Where("1 = 1").Delete(&User{})
  // // DELETE FROM `users` WHERE 1=1

  // db.Exec("DELETE FROM users")
  // // DELETE FROM users

  // db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
  // // DELETE FROM users
}