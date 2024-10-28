package routers

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/delivery/controllers"
	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/repository"
	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	notiRepo := repository.NewNotificationRepository()
	notiUsecase := usecase.NewNotificationUsecase(notiRepo)
	notiController := controllers.NewNotificationController(notiUsecase)

	r.GET("/notifications", notiController.GetAllNotifications)
}
