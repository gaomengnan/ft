package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gaomengnan/ft/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	routes.RegisterRoute(app)
	go func() {
		log.Fatal(app.Listen(port))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	app.Shutdown()
	os.Exit(0)
}
