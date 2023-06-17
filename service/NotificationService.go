package service

import (
	"log"
	"notification_service/model"
	"notification_service/repository"

	"github.com/google/uuid"
)

type NotificationService struct {
	Repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		Repo: repo,
	}
}

func (service *NotificationService) Save(notification *model.Notification) (model.RequestMessage, error) {
	response := model.RequestMessage{
		Message: service.Repo.Save(notification).Message,
	}
	return response, nil
}

func (service *NotificationService) GetAllByUserID(userID uuid.UUID) ([]model.Notification, error) {
	accomodations, err := service.Repo.GetAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return accomodations, nil
}
