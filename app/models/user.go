package models

// type User struct {
//     ID
//     Name string `json:"name" gorm:"not null;comment:用户名称"`
//     Mobile string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
//     Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
//     Timestamps
//     SoftDeletes
// }
type User struct {
	ID    int
	Sid       string
	Password         string
}
func (User) TableName() string {
	return "user"
}