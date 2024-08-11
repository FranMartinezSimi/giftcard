package handler

import (
	"GiftWize/src/app/usecase"
	"GiftWize/src/entity/request"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CampaignHandler struct {
	useCase usecase.CampaignUseCase
}

func NewCampaignHandler(useCase usecase.CampaignUseCase) *CampaignHandler {
	return &CampaignHandler{
		useCase: useCase,
	}
}

// TODO: Implementar los metodos de los handlers

func (h *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("CreateCampaign usecase")

	var body request.CreateCampaignRequest
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return err
	}

	err := h.useCase.CreateCampaign(ctx.Context(), body)
	if err != nil {
		log.Errorf("Error creating campaign: %v", err)
		return err
	}

	log.Info("Campaign created successfully")
	return ctx.SendStatus(fiber.StatusCreated)
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
