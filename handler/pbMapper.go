package handler

import (
	"log"
	"notification_service/model"
	"strconv"
	"time"

	pb "github.com/XML-organization/common/proto/notification_service"
	"github.com/google/uuid"
)

func mapNotificationFromSaveNotification(notification *pb.SaveRequest) model.Notification {

	userID, err := uuid.Parse(notification.UserID)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	id, err := uuid.Parse(notification.Id)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	statusInt, err := strconv.Atoi(notification.Status)
	if err != nil {
		log.Println(err)
	}

	layout := "2006-01-02"

	notificationTime, err := time.Parse(layout, notification.NotificationTime)
	if err != nil {
		log.Println(err)
	}

	return model.Notification{
		ID:               id,
		Text:             notification.Text,
		NotificationTime: notificationTime,
		UserID:           userID,
		Status:           model.NotificationStatus(statusInt),
	}
}

func mapSaveNotificationFromNotification(notification *model.Notification) *pb.SaveRequest {

	id := notification.ID.String()
	notificationTime := notification.NotificationTime.Format("2006-01-02")
	usedID := notification.UserID.String()

	var statusString string
	switch notification.Status {
	case model.NOT_SEEN:
		statusString = "NOT_SEEN"
	case model.SEEN:
		statusString = "SEEN"
	}

	return &pb.SaveRequest{
		Id:               id,
		Text:             notification.Text,
		NotificationTime: notificationTime,
		UserID:           usedID,
		Status:           statusString,
	}
}
