package models

import "strconv"

type Progress struct {
	ID int
	Pid string 
	Sid string
	Test int 
	Interview int 
    Check1 int 
    Check2 int
	Offer int
	Files1 string
	Files2 string
	Link1 string
	Link2 string
	Timestamps
}

type Delivery struct {
	ID int
	Pid       string
	Sid         string
	Timestamps
}


func (Progress) TableName() string {
	return "progress"
}
func (progress Progress) GetUid() string {
    return strconv.Itoa(int(progress.ID))
}
func (delivery Delivery) GetUid() string {
    return strconv.Itoa(int(delivery.ID))
}