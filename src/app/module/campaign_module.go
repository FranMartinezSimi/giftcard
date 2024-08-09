package module

import (
	"GiftWize/src/infreaestructure/repository"
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
)

func CampaignModule(app *fiber.App) {
	db := shared.Init()
	repository := repository.NewCampaignRepository(db)

}
