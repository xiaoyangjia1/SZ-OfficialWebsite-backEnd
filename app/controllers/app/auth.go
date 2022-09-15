package app

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/common/response"
    "SZ-OfficialWebsite-backEnd/app/services"
)

func Login(c *gin.Context) {
    var form request.Login
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }

    if err, user := services.UserService.Login(form); err != nil {
        response.BusinessFail(c, err.Error())
    } else {
        tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
        if err != nil {
            response.BusinessFail(c, err.Error())
            return
        }
        response.Success(c, tokenData)
    }
}

func Logout(c *gin.Context) {
    err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
    if err != nil {
        response.BusinessFail(c, "登出失败")
        return
    }
    response.Success(c, nil)
}

func Info(c *gin.Context) {
    err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    }
    response.Success(c, user)
}
