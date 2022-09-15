package services

import (
    "SZ-OfficialWebsite-backEnd/app/common/request"
    "SZ-OfficialWebsite-backEnd/app/models"
    "SZ-OfficialWebsite-backEnd/global"
)

type resumeService struct {
}

var ResumeService = new(resumeService)

func (resumeService *resumeService) SubmitResume(params request.Resume) (err error, resume models.Resume) {
    resume = models.Resume{	
        Sid: params.Sid,
        Phone: params.Phone, 
        Email: params.Email,
        Wechat: params.Wechat, 
        Sex: params.Sex, 
        Name: params.Name, 
        College: params.College, 
        Photo: params.Photo, 
        Learning: params.Learning, 
        Slife: params.Slife, 
        Link: params.Link, 
        Files: params.Files, 
    }
    err = global.App.DB.Create(&resume).Error
    return
}

func (resumeService *resumeService) GetResumeById(params request.Sid) (err error, resume models.Resume) {
    err = global.App.DB.Where("sid = ?", params.Sid).Find(&resume).Error
    return
}

