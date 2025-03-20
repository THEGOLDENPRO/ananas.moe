package main

import (
	"ananasmoe/routes"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	filePath := os.Getenv("FILES_PATH")
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	log.Print(allowedOrigins)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	routes.Redirects(app)
	routes.Clouds(app)
	routes.Projects(app)
	routes.Files(app, filePath)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hewwo")
	})

	log.Fatal(app.Listen(":3000"))
}

// UwU
