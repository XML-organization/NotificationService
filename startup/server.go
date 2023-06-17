package startup

import (
	"fmt"
	"log"
	"net"
	"notification_service/handler"
	"notification_service/repository"
	"notification_service/service"
	"notification_service/startup/config"

	notification "github.com/XML-organization/common/proto/notification_service"

	saga "github.com/XML-organization/common/saga/messaging"
	"github.com/XML-organization/common/saga/messaging/nats"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "notification_service"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	notificationRepo := server.initNotificationRepository(postgresClient)
	notificationService := server.initNotificationService(notificationRepo)
	notificationHandler := server.initNotificationHandler(notificationService)

	server.startGrpcServer(notificationHandler)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := repository.GetClient(
		server.config.NotificationDBHost, server.config.NotificationDBUser,
		server.config.NotificationDBPass, server.config.NotificationDBName,
		server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationRepository(client *gorm.DB) *repository.NotificationRepository {
	return repository.NewNotificationRepository(client)
}

func (server *Server) initNotificationService(repo *repository.NotificationRepository) *service.NotificationService {
	return service.NewNotificationService(repo)
}

func (server *Server) initNotificationHandler(service *service.NotificationService) *handler.NotificationHandler {
	return handler.NewNotificationHandler(service)
}

func (server *Server) startGrpcServer(notificationHandler *handler.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	notification.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	reflection.Register(grpcServer)
	println("GRPC SERVER USPJESNO NAPRAVLJEN")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		println("GRPC SERVER NIJE USPJESNO NAPRAVLJEN")
	}
}
