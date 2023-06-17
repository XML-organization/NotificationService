package main

import (
	"log"
	"notification_service/startup"
	cfg "notification_service/startup/config"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	config := cfg.NewConfig()
	log.Println("Starting server Notification Service...")
	server := startup.NewServer(config)
	server.Start()
}
