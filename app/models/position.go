package models

import "strconv"

type Position struct {
	ID int
	Pid string 
	Title string 
	Batch string 
	Deadline string
	Description string 
	Requirements string 
	Apply_number int 
	Head_count int 
	Test int 
	Interview int 
    Check1 int 
    Check2 int
}
func (Position) TableName() string {
	return "position"
}
func (position Position) GetUid() string {
    return strconv.Itoa(int(position.ID))
}