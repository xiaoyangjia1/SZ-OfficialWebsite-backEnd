package main

import (
  "fmt"
  "connectDB"
)
type User struct {
	ID    int
	Sid       string
	Password         string
}
func (User) TableName() string {
	return "user"
}
func main()  {
	db :=connectDB.Connect()
	user := User{Sid: "3119005073", Password: "200101"}
	result := db.Create(&user) // 通过数据的指针来创建
	// result :=db.Select("sid", "password").Create(&user)

	// data := db.Find(&user, "sid = ?", "3119005073")
	fmt.Println(result)
	// fmt.Println(data)
	// fmt.Println(user)
}