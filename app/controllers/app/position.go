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

func GetJobs(c *gin.Context) {
    err, jobs := services.PositionService.GetJobs(); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, jobs)
}

func ChangeJobStatus(c *gin.Context) {
    var form request.Position_id
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }
    err, newStatus := services.PositionService.ChangeJobStatus(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, newStatus)
}
func DeleteJob(c *gin.Context) {
    var form request.Position_id
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }
    err, res := services.PositionService.DeleteJob(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, res)
}