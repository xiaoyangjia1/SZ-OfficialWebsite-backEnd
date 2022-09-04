package services

import (
    "fmt"
    "time"
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/models"
    "SZ-OfficialWebsite-backEnd/global"
)

type positionService struct {
}

var PositionService = new(positionService)

func (positionService *positionService) PostJob(params request.Position) (err error, position models.Position) {
  fmt.Println(params)
  //返回time对象
  t := time.Unix(int64(params.Deadline), 0)
  //返回string
  dateStr := t.Format("2006-01-02 15:04:05")

    position = models.Position{	
        Pid: "SZ2023FE",
        Title: params.Title, 
        Batch: params.Batch,
        Deadline: dateStr, 
        Description: params.Description, 
        Requirements: params.Requirements, 
        Apply_number: 0,
        Head_count: params.Head_count, 
        Test: params.Test,
        Interview: params.Interview, 
        Check1: params.Check1, 
        Check2: params.Check2, 
    }
    fmt.Println(position)
    err = global.App.DB.Create(&position).Error
    return
}
func (positionService *positionService) GetJobs() (err error, position []models.Position) {
    err = global.App.DB.Find(&position).Error
    return 
  }
  
