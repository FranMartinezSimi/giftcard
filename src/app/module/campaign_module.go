package module

import (
	"GiftWize/src/app/usecase"
	handler2 "GiftWize/src/infreaestructure/handler"
	"GiftWize/src/infreaestructure/repository"
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
)

func CampaignModule(app *fiber.App) {
	db := shared.Init()
	campaignRepo := repository.NewCampaignRepository(db)
	campaignUseCase := usecase.NewCampaignUseCase(*campaignRepo)
	handler := handler2.NewCampaignHandler(*campaignUseCase)

	app.Post("/campaign", handler.CreateCampaign)
	app.Get("/campaign/:id", handler.GetCampaign)
	app.Put("/campaign/:id", handler.UpdateCampaign)
	app.Delete("/campaign/:id", handler.DeleteCampaign)
	app.Get("/campaigns", handler.ListCampaigns)
}
