package module

import (
	"GiftWize/src/app/usecase"
	handler2 "GiftWize/src/infreaestructure/handler"
	"GiftWize/src/infreaestructure/repository"
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
)

func GiftCardModule(app *fiber.App) {
	db := shared.Init()
	giftCardRepo := repository.NewGiftCardRepository(db)    // Returns IGiftCardRepository
	giftCardUseCase := usecase.NewGiftCardUseCase(giftCardRepo) // Expects IGiftCardRepository, returns IGiftCardUseCase
	handler := handler2.NewGiftCardHandler(giftCardUseCase)  // Expects IGiftCardUseCase

	app.Post("/giftcard", handler.CreateGiftCard)
	app.Get("/giftcard/:id", handler.GetGiftCardByID)
	app.Put("/giftcard/:id", handler.UpdateGiftCard)
	app.Delete("/giftcard/:id", handler.DeleteGiftCard)
	app.Get("/giftcards", handler.GetAllGiftCards)
	app.Get("/giftcards/search", handler.FullTextSearchGiftCard)
}
