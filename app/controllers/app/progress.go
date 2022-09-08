package app

import (
    "github.com/gin-gonic/gin"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/common/response"
    "SZ-OfficialWebsite-backEnd/app/services"
)

func DeliveryJob(c *gin.Context) {
    var form request.Delivery_item
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }
    err, res := services.ProgressService.DeliveryJob(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, res)
}

func GetDeliveredJob(c *gin.Context) {
    var form request.Sid
    if err := c.ShouldBindJSON(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }
    err, res := services.ProgressService.GetDeliveredJob(form); 
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    } 
    response.Success(c, res)
}