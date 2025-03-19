package routes

import (
	"ananasmoe/types"
	"ananasmoe/utils"

	"github.com/gofiber/fiber/v3"
)

func Redirects(app *fiber.App) {
	app.Get("/r/:id", func(c fiber.Ctx) error {
		redirects, err := utils.GetConfig[[]types.Redirect]("redirect")
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"status": 500,
				"error":  "Failed to read config",
			})
		}

		id := c.Params("id")
		for _, v := range redirects {
			if v.ID == id {
				return c.JSON(v)
			}
		}

		return c.Status(404).JSON(&fiber.Map{
			"status": 404,
			"error":  "No redirect with that id",
		})
	})
}
