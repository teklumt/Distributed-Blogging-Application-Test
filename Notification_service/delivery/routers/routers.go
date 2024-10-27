package routers

import (
	"teklumt/Distributed-Blogging-Application-Test-notification-service/delivery/controllers"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/repository"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	notiRepo := repository.NewNotificationRepository()
	notiUsecase := usecase.NewNotificationUsecase(notiRepo)
	notiController := controllers.NewNotificationController(notiUsecase)

	// r.POST("/notifications", notiController.CreateNotification)
	r.GET("/notifications", notiController.GetAllNotifications)
	// r.GET("/notifications", notiController.GetAllNotifications) // New route
}
