package app

import (
    "github.com/gin-gonic/gin"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/common/response"
    "SZ-OfficialWebsite-backEnd/app/services"
)

func SubmitResume(c *gin.Context) {
    var form request.Resume
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }

    err, resume := services.ResumeService.SubmitResume(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, resume)
}

func GetResumeById(c *gin.Context) {
    var form request.Sid
    if err := c.ShouldBindQuery(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }

    err, resume := services.ResumeService.GetResumeById(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, resume)
}

