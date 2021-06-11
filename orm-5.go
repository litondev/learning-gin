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

  // UPDATE
  user := Latihan{ID:4}
  db.First(&user)
  user.Name = "p"
  db.Save(&user)
  fmt.Println(user)

// HARUS ADA KONDISI Where
// Update with conditions
// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// User's ID is `111`:
// user := Latihan{ID:111}
// db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// Update with conditions and model value
// user := Latihan{ID:111}
// db.Model(&user).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

// Update attributes with `struct`, will only update non-zero fields
// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// Update attributes with `map`
// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// Get updated records count with `RowsAffected`
// result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin;
// result.RowsAffected // returns updated records count
// result.Error        // returns updating error
}