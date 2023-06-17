package handler

import (
	"context"
	"log"
	"notification_service/service"

	pb "github.com/XML-organization/common/proto/notification_service"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	Service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		Service: service,
	}
}

func (handler *NotificationHandler) GetAllForUser(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	println("Usao u NotificationHandler.GetAllByUserID-----")
	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println("UserID u Notification Servicu poslije parsiranja: ", userID.String())

	notifications, err := handler.Service.GetAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	print("Lista koja je ucitana iz baze")
	println("Dugacka je: ", len(notifications))
	for j, tmp := range notifications {
		println(j, ". Notification: ", tmp.Text)
	}

	response := &pb.GetAllResponse{
		Notifications: []*pb.SaveRequest{},
	}

	for _, notification := range notifications {
		current := mapSaveNotificationFromNotification(&notification)
		response.Notifications = append(response.Notifications, current)
	}

	println("Notifications koje vraca Notifications Servis:")
	println("Dugacka je: ", len(response.Notifications))

	for j, tmp := range response.Notifications {
		println(j, ". Notifications: ", tmp.Text)
	}

	return response, nil
}

func (handler *NotificationHandler) Save(ctx context.Context, request *pb.SaveRequest) (*pb.SaveResponse, error) {
	notification := mapNotificationFromSaveNotification(request)
	message, err := handler.Service.Save(&notification)
	if err != nil {
		log.Println(err)
	}
	response := pb.SaveResponse{
		Message: message.Message,
	}

	return &response, err
}
