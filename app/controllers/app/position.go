package app

import (
    "github.com/gin-gonic/gin"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/common/response"
    "SZ-OfficialWebsite-backEnd/app/services"
)

func PostJob(c *gin.Context) {
    var form request.Position
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }

    err, jobData := services.PositionService.PostJob(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, jobData)
}