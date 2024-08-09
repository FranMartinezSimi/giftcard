package handler

import (
	"GiftWize/src/app/usecase"
	"github.com/gofiber/fiber/v2"
)

type CampaignHandler struct {
	useCase usecase.CampaignUseCase
}

func NewCampaignHandler(useCase usecase.CampaignUseCase) *CampaignHandler {
	return &CampaignHandler{
		useCase: useCase,
	}
}

// TODO Implementar los metodos de los handlers

func (h *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	return nil
}

func (h *CampaignHandler) GetCampaign(ctx *fiber.Ctx) error {
	return nil
}

func (h *CampaignHandler) UpdateCampaign(ctx *fiber.Ctx) error {
	return nil
}

func (h *CampaignHandler) DeleteCampaign(ctx *fiber.Ctx) error {
	return nil
}

func (h *CampaignHandler) ListCampaigns(ctx *fiber.Ctx) error {
	return nil
}
