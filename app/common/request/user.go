package request

type Register struct {
    Email string `form:"email" json:"email" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
    return ValidatorMessages{
        "Email.required": "学号不能为空",
        "Password.required": "密码不能为空",
    }
}

type Login struct {
    Email string `form:"email" json:"email" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
    return ValidatorMessages{
        "Email.required": "学号不能为空",
        "Password.required": "密码不能为空",
    }
}