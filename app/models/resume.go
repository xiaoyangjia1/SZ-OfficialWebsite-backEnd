package models

import "strconv"

type Resume struct {
	ID int
	Sid string 
	Phone string 
	Email string 
	Wechat string 
	Sex int 
	Name string 
	College string 
	Photo string 
	Learning string 
	Slife string 
    Link string 
    Files string 
}
func (Resume) TableName() string {
	return "resume"
}
func (resume Resume) GetUid() string {
    return strconv.Itoa(int(resume.ID))
}