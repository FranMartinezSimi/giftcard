package main

import (
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	envs := shared.GetEnvs()
	shared.Init()

	err := app.Listen(":" + envs["PORT"])
	if err != nil {
		panic(err)
	}
}
