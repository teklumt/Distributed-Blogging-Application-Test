package domain

import "time"

type Notification struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `json:"user_id"`
	Message    string    `json:"message"`
	ReadStatus bool      `json:"read_status"`
	CreatedAt  time.Time `json:"created_at"`
}
func (Notification) TableName() string {
	return "Notifications"
}

type NotificationUsecase interface {
	CreateNotification(notification *Notification) error
	// GetUserNotifications(userID uint) ([]Notification, error)
	GetAllNotifications() ([]Notification, error)
}

type NotificationRepository interface {
	CreateNotification(notification *Notification) error
	// GetUserNotifications(userID uint) ([]Notification, error)
	GetAllNotifications() ([]Notification, error)
}
