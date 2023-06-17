package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID               uuid.UUID          `json:"id"`
	Text             string             `json:"text"`
	NotificationTime time.Time          `json:"notificationTime" gorm:"not null;type:string"`
	UserID           uuid.UUID          `gorm:"column:userID" json:"userID"`
	Status           NotificationStatus `json:"status"`
}

type RequestMessage struct {
	Message string `json:"message"`
}

/*func (not *Notification) BeforeCreate(scope *gorm.DB) error {
	not.ID = uuid.New()
	return nil
}*/
