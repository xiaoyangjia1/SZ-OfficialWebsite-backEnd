package services

import (
	"fmt"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/models"
    "SZ-OfficialWebsite-backEnd/global"
)

type progressService struct {
}

var ProgressService = new(progressService)

func (progressService *progressService) DeliveryJob(params request.Delivery_item) (err error, progress models.Progress) {
	var position models.Position
  global.App.DB.First(&position, "pid = ?", params.Pid)
  progress = models.Progress{	
    Pid: params.Pid,
		Sid: params.Sid,
    Test: position.Test,
    Interview: position.Interview, 
    Check1: position.Check1, 
    Check2: position.Check2, 
    Offer: 0,
  }
  fmt.Println(progress)
  err = global.App.DB.Create(&progress).Error
  return 
}

func (progressService *progressService) GetDeliveredJob(params request.Sid) (err error, progress []models.Progress) {
  err = global.App.DB.Where("sid = ?", params.Sid).Find(&progress).Error
  return 
}

func (progressService *progressService) SubmitEpicycleResult(params request.Submit_epicycle_result) (err error, progress models.Progress) {
  global.App.DB.Model(&progress).Where("pid = ?",params.Pid).Where("sid = ?",params.Sid).Update(params.Epicycle,params.Result)
  return 
}
  
  