package controllers

import (
	"net/http"

	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/domain"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationUsecase domain.NotificationUsecase
}

func NewNotificationController(notificationUsecase domain.NotificationUsecase) *NotificationController {
	return &NotificationController{NotificationUsecase: notificationUsecase}
}

// Retrieve all notifications for a specific user
func (c *NotificationController) GetAllNotifications(ctx *gin.Context) {
	notifications, err := c.NotificationUsecase.GetAllNotifications()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
