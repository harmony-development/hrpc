package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	chat "github.com/harmony-development/hrpc/examples/chat/gen"
	"github.com/harmony-development/hrpc/server/fiberadapter"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	service := chat.NewChatServiceHandler(&ChatService{
		msgChan: make(chan *chat.Message),
	})
	fiberadapter.RegisterFiber(service, app)
	log.Fatal(app.Listen(":6969"))
}
