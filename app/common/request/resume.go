package request

type Resume struct {
	Sid string `form:"sid" json:"sid" binding:"required"`
	Phone string `form:"phone" json:"phone" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Wechat string `form:"wechat" json:"wechat" binding:"required"`
	Sex int `form:"sex" json:"sex" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	College string `form:"college" json:"college" binding:"required"`
	Photo string `form:"photo" json:"photo" binding:"required"`
	Learning string `form:"learning" json:"learning" binding:"required"`
	Slife string `form:"slife" json:"slife" binding:"required"`
    Link string `form:"Link" json:"Link"`
    Files string `form:"Files" json:"Files"`
}

func (resume Resume) GetMessages() ValidatorMessages {
    return ValidatorMessages{
		"Sid.required": "学号不能为空",
		"Phone.required": "手机号不能为空",
		"Email.required": "邮件不能为空",
		"Wechat.required": "微信号不能为空",
		"Sex.required": "性别不能为空",
		"Name.required": "名字不能为空",
		"College.required": "学院不能为空",
		"Photo.required": "照片不能为空",
		"Learning.required": "学习情况不能为空",
		"Slife.required": "校园生活不能为空",
    }
}