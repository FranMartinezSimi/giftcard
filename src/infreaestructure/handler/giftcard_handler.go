package handler

import (
	"GiftWize/src/app/usecase"
	"GiftWize/src/entity/request"
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type GiftCardHandler struct {
	giftCardUseCase usecase.IGiftCardUseCase // Depends on the interface
}

func NewGiftCardHandler(giftCardUseCase usecase.IGiftCardUseCase) *GiftCardHandler { // Accepts the interface
	return &GiftCardHandler{
		giftCardUseCase: giftCardUseCase,
	}
}

func (g *GiftCardHandler) CreateGiftCard(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("CreateGiftCard usecase")

	var body request.CreateGiftCardRequest
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Validate the request body
	if validationErrors := shared.ValidateStruct(body); len(validationErrors) > 0 {
		log.Error("Validation errors:", validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(shared.FormatValidationErrors(validationErrors))
	}

	err := g.giftCardUseCase.CreateGiftCard(ctx.Context(), body)
	if err != nil {
		log.Errorf("Error creating gift card: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Gift card created successfully")
	return ctx.SendStatus(fiber.StatusCreated)
}

func (g *GiftCardHandler) DeleteGiftCard(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("DeleteGiftCard usecase")

	id := ctx.Params("id")
	if id == "" {
		log.Error("Gift card ID is required")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := g.giftCardUseCase.DeleteGiftCard(ctx.Context(), id)
	if err != nil {
		log.Errorf("Error deleting gift card: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Gift card deleted successfully")
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (g *GiftCardHandler) UpdateGiftCard(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("UpdateGiftCard usecase")

	id := ctx.Params("id")
	if id == "" {
		log.Error("Gift card ID is required")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var body request.UpdateGiftCardRequest
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Validate the request body
	if validationErrors := shared.ValidateStruct(body); len(validationErrors) > 0 {
		log.Error("Validation errors:", validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(shared.FormatValidationErrors(validationErrors))
	}

	err := g.giftCardUseCase.UpdateGiftCard(ctx.Context(), id, body)
	if err != nil {
		log.Errorf("Error updating gift card: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Gift card updated successfully")
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (g *GiftCardHandler) GetAllGiftCards(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("GetAllGiftCards usecase")

	list, err := g.giftCardUseCase.GetAllGiftCardList(ctx.Context())
	if err != nil {
		log.Errorf("Error getting gift card list: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(list)
}

func (g *GiftCardHandler) GetGiftCardByID(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("GetGiftCardByID usecase")

	id := ctx.Params("id")
	if id == "" {
		log.Error("Gift card ID is required")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	card, err := g.giftCardUseCase.GetGiftCardByID(ctx.Context(), id)
	if err != nil {
		log.Errorf("Error getting gift card by ID: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(card)
}

func (g *GiftCardHandler) FullTextSearchGiftCard(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("FullTextSearchGiftCard usecase")

	query := ctx.Query("query")
	results, err := g.giftCardUseCase.FullTextSearchGiftCard(ctx.Context(), query)
	if err != nil {
		log.Errorf("Error full text searching gift card: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(results)
}
