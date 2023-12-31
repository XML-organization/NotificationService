package repository

import (
	"log"
	"notification_service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	DatabaseConnection *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	err := db.AutoMigrate(&model.Notification{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return &NotificationRepository{
		DatabaseConnection: db,
	}
}

func (repo *NotificationRepository) Save(notification *model.Notification) model.RequestMessage {
	dbResult := repo.DatabaseConnection.Save(notification)

	if dbResult.Error != nil {
		log.Println(dbResult.Error)
		return model.RequestMessage{
			Message: "An error occured, please try again!",
		}
	}

	return model.RequestMessage{
		Message: "Success!",
	}
}

func (repo *NotificationRepository) GetAll() ([]model.Notification, model.RequestMessage) {
	var notification []model.Notification
	dbResult := repo.DatabaseConnection.Find(&notification)

	if dbResult.Error != nil {
		log.Println(dbResult.Error)
		return nil, model.RequestMessage{
			Message: "An error occurred, please try again!",
		}
	}

	return notification, model.RequestMessage{
		Message: "Success!",
	}
}

func (repo *NotificationRepository) GetAllByUserID(userID uuid.UUID) ([]model.Notification, error) {
	notifications := []model.Notification{}
	result := repo.DatabaseConnection.Where("userID = ?", userID).Find(&notifications)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return notifications, nil
}
