package models

import "strconv"

type User struct {
	ID    int
	Email       string
	Password         string
}
func (User) TableName() string {
	return "user"
}
func (user User) GetUid() string {
    return strconv.Itoa(int(user.ID))
}