package routes

import (
    "github.com/gin-gonic/gin"
    "SZ-OfficialWebsite-backEnd/app/controllers/app"
    "SZ-OfficialWebsite-backEnd/app/middleware"
    "SZ-OfficialWebsite-backEnd/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
    router.POST("/auth/register", app.Register)
    router.POST("/auth/login", app.Login)

    authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
    {
        authRouter.POST("/auth/info", app.Info)
        authRouter.POST("/auth/logout", app.Logout)
        authRouter.POST("/auth/submitResume", app.SubmitResume)
        authRouter.POST("/auth/postJob", app.PostJob)
        authRouter.GET("/auth/getJobs", app.GetJobs)
        authRouter.POST("/auth/changeJobStatus", app.ChangeJobStatus)
    }
}
