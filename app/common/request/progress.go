package request

type Delivery_item struct {
	Pid string `form:"pid" json:"pid" binding:"required"`
	Sid string `form:"sid" json:"sid" binding:"required"`
}
func (delivery_item Delivery_item) GetMessages() ValidatorMessages {
    return ValidatorMessages{
		"Pid.required": "岗位ID不能为空",
		"Sid.required": "用户ID不能为空",
    }
}

type Sid struct {
	Sid string `form:"sid" json:"sid" binding:"required"`
}
func (sid Sid) GetMessages() ValidatorMessages {
    return ValidatorMessages{
		"Sid.required": "用户ID不能为空",
    }
}