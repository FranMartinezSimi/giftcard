package rest

import (
	"GiftWize/src/app/usecase"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CampaignHandler struct {
	useCase usecase.CampaignUseCase
}

func NewCampaign(useCase usecase.CampaignUseCase) *CampaignHandler {
	return &CampaignHandler{useCase: useCase}
}

func (c *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	log.WithContext(ctx.Context()).Info("CreateCampaign hanlder")
	return ctx.SendString("CreateCampaign")
}

func (c *CampaignHandler) GetCampaign(ctx *fiber.Ctx) error {
	log.WithContext(ctx.Context()).Info("GetCampaign hanlder")
	return ctx.SendString("GetCampaign")
}

func (c *CampaignHandler) UpdateCampaign(ctx *fiber.Ctx) error {
	log.WithContext(ctx.Context()).Info("UpdateCampaign hanlder")
	return ctx.SendString("UpdateCampaign")
}
