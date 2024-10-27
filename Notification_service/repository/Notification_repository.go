package repository

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/config"
	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/domain"
)

type NotificationRepository struct{}

func NewNotificationRepository() domain.NotificationRepository {
	return &NotificationRepository{}
}

func (r *NotificationRepository) CreateNotification(notification *domain.Notification) error {
	err := config.DB.Create(notification).Error
	return err
}


func (r *NotificationRepository) GetAllNotifications() ([]domain.Notification, error) {
	var res []domain.Notification
	err := config.DB.Find(&res).Error	
	if err != nil {
		return []domain.Notification{}, err
	}
	return res, nil
}
