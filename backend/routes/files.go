package routes

import (
	"ananasmoe/types"
	"ananasmoe/utils"
	"fmt"
	"path"

	"github.com/gofiber/fiber/v3"
)

func Files(app *fiber.App, filePath string) {
	app.Get("/file/:id", func(c fiber.Ctx) error {
		files, err := utils.GetConfig[[]types.File]("file")
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"status": 500,
				"error":  fmt.Sprintf("Failed to read config: %s", err),
			})
		}

		id := c.Params("id")
		for _, v := range files {
			if v.ID == id {
				return c.JSON(v)
			}
		}

		return c.Status(404).JSON(&fiber.Map{
			"status": 404,
			"error":  "No file with that id",
		})
	})

	app.Get("/file/:id/download", func(c fiber.Ctx) error {
		files, err := utils.GetConfig[[]types.File]("file")
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"status": 500,
				"error":  "Failed to read config",
			})
		}

		id := c.Params("id")
		for _, v := range files {
			if v.ID == id {
				file := path.Join(filePath, v.File)
				fmt.Print(file, filePath)

				return c.SendFile(file)
			}
		}

		return c.Status(404).JSON(&fiber.Map{
			"status": 404,
			"error":  "No file with that id",
		})
	})
}
