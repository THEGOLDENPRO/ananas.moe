package routes

import (
	"ananasmoe/types"
	"ananasmoe/utils"

	"github.com/gofiber/fiber/v3"
)

func Projects(app *fiber.App) {
	app.Get("/projects", func(c fiber.Ctx) error {
		projects, err := utils.GetConfig[[]types.Project]("project")
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"status": 500,
				"error":  "Failed to read config",
			})
		}

		return c.JSON(projects)
	})
}
