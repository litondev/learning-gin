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

  // Get the first record ordered by primary key
  db.First(&user)
  // SELECT * FROM users ORDER BY id LIMIT 1;

  // Get one record, no specified order
  db.Take(&user)
  // SELECT * FROM users LIMIT 1;

  // Get last record, order by primary key desc
  db.Last(&user)
  // SELECT * FROM users ORDER BY id DESC LIMIT 1;

  result := db.First(&user)
  result.RowsAffected // returns found records count
  result.Error        // returns error

  // check error ErrRecordNotFound
  errors.Is(result.Error, gorm.ErrRecordNotFound)


  var user User
  var users []User  

  // works
  db.First(&user)
  // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  // works
  result := map[string]interface{}{}
  db.Model(&User{}).First(&result)
  // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

  // doesn't work
  result := map[string]interface{}{}
  db.Table("users").First(&result)

  // works with Take
  result := map[string]interface{}{}
  db.Table("users").Take(&result)

  // order by first field
  type Language struct {
    Code string
    Name string
  }
  db.First(&Language{})
  // SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1

  db.First(&user, 10)
  // SELECT * FROM users WHERE id = 10;

  db.First(&user, "10")
  // SELECT * FROM users WHERE id = 10;

  db.Find(&users, []int{1,2,3})
  // SELECT * FROM users WHERE id IN (1,2,3);

  db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
  // SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

  // Get all records
  result := db.Find(&users)
  // SELECT * FROM users;

  result.RowsAffected // returns found records count, equals `len(users)`
  result.Error        // returns error

  // Get first matched record
  db.Where("name = ?", "jinzhu").First(&user)
  // SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

  // Get all matched records
  db.Where("name <> ?", "jinzhu").Find(&users)
  // SELECT * FROM users WHERE name <> 'jinzhu';

  // IN
  db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
  // SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

  // LIKE
  db.Where("name LIKE ?", "%jin%").Find(&users)
  // SELECT * FROM users WHERE name LIKE '%jin%';

  // AND
  db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
  // SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

  // Time
  db.Where("updated_at > ?", lastWeek).Find(&users)
  // SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

  // BETWEEN
  db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
  // SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

  // Struct
  db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
  // SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

  // Map
  db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
  // SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

  // Slice of primary keys
  db.Where([]int64{20, 21, 22}).Find(&users)
  // SELECT * FROM users WHERE id IN (20, 21, 22);


  db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
  // SELECT * FROM users WHERE name = "jinzhu";

  db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
  // SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

  db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
  // SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

  db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
  // SELECT * FROM users WHERE age = 0;


  // SELECT * FROM users WHERE id = 23;
// Get by primary key if it were a non-integer type
db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

// Plain SQL
db.Find(&user, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

// Map
db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;

db.Not("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

// Not In
db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

// Struct
db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

// Not In slice of primary keys
db.Not([]int64{1,2,3}).First(&user)
// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1

db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

db.Select("name", "age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;

db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

// Multiple orders
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

db.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&User{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

type result struct {
  Date  time.Time
  Total int
}

db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name`


db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
for rows.Next() {
  ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
for rows.Next() {
  ...
}

type Result struct {
  Date  time.Time
  Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)


db.Distinct("name", "age").Order("name, age desc").Find(&results)

type result struct {
  Name  string
  Email string
}
db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
}