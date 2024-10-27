package usecase

import "teklumt/Distributed-Blogging-Application-Test-notification-service/domain"

type NotificationUsecase struct {
	NotificationRepository domain.NotificationRepository
}

func NewNotificationUsecase(notificationRepository domain.NotificationRepository) domain.NotificationUsecase {
	return &NotificationUsecase{NotificationRepository: notificationRepository}
}

func (u *NotificationUsecase) CreateNotification(notification *domain.Notification) error {
	return u.NotificationRepository.CreateNotification(notification)
}

func (u *NotificationUsecase) GetAllNotifications() ([]domain.Notification, error) {
	return u.NotificationRepository.GetAllNotifications()
}
