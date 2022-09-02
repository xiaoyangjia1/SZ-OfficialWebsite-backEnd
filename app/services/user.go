package services

import (
    "strconv"
    "errors"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/models"
    "SZ-OfficialWebsite-backEnd/global"
    "SZ-OfficialWebsite-backEnd/utils"
)

type userService struct {
}

var UserService = new(userService)


func (userService *userService) Register(params request.Register) (err error, user models.User) {
    var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
    if result.RowsAffected != 0 {
        err = errors.New("用户已存在")
        return
    }
    user = models.User{Email: params.Email, Password: utils.BcryptMake([]byte(params.Password))}
    err = global.App.DB.Create(&user).Error
    return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
    err = global.App.DB.Where("email = ?", params.Email).First(&user).Error
    if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
        err = errors.New("用户名不存在或密码错误")
    }
    return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
    intId, err := strconv.Atoi(id)
    err = global.App.DB.First(&user, intId).Error
    if err != nil {
        err = errors.New("数据不存在")
    }
    return
}