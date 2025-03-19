package routes

import (
	"ananasmoe/types"
	"ananasmoe/utils"

	"github.com/gofiber/fiber/v3"
)

func Clouds(app *fiber.App) {
	app.Get("/clouds", func(c fiber.Ctx) error {
		clouds, err := utils.GetConfig[[]types.Cloud]("cloud")
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"status": 500,
				"error":  "Failed to read config",
			})
		}

		return c.JSON(clouds)
	})
}
