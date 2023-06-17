package config

import "os"

type Config struct {
	Port                     string
	NotificationDBHost       string
	NotificationDBPort       string
	NotificationDBName       string
	NotificationDBUser       string
	NotificationDBPass       string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	CreateUserCommandSubject string
	CreateUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("NOTIFICATION_SERVICE_PORT"),
		NotificationDBHost:       os.Getenv("NOTIFICATION_DB_HOST"),
		NotificationDBPort:       os.Getenv("NOTIFICATION_DB_PORT"),
		NotificationDBName:       os.Getenv("NOTIFICATION_DB_NAME"),
		NotificationDBUser:       os.Getenv("NOTIFICATION_DB_USER"),
		NotificationDBPass:       os.Getenv("NOTIFICATION_DB_PASS"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		CreateUserCommandSubject: os.Getenv("CREATE_USER_COMMAND_SUBJECT"),
		CreateUserReplySubject:   os.Getenv("CREATE_USER_REPLY_SUBJECT"),
	}
}
