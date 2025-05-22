package handler

import (
	"GiftWize/src/app/usecase"
	"GiftWize/src/entity/request"
	"GiftWize/src/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CampaignHandler struct {
	useCase usecase.ICampaignUseCase // Depends on the interface
}

func NewCampaignHandler(useCase usecase.ICampaignUseCase) *CampaignHandler { // Accepts the interface
	return &CampaignHandler{
		useCase: useCase,
	}
}

func (h *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("CreateCampaign usecase")

	var body request.CreateCampaignRequest
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Validate the request body
	if validationErrors := shared.ValidateStruct(body); len(validationErrors) > 0 {
		log.Error("Validation errors:", validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(shared.FormatValidationErrors(validationErrors))
	}

	err := h.useCase.CreateCampaign(ctx.Context(), body)
	if err != nil {
		log.Errorf("Error creating campaign: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Campaign created successfully")
	return ctx.SendStatus(fiber.StatusCreated)
}

func (h *CampaignHandler) GetCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("GetCampaign handler")

	id, err := ctx.ParamsInt("id")

	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	campaign, err := h.useCase.GetCampaign(ctx.Context(), id)
	if err != nil {
		log.Errorf("Error getting campaign: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if campaign == nil {
		log.Error("Campaign not found")
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(campaign)
}

func (h *CampaignHandler) UpdateCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("UpdateCampaign handler")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
	}

	var body request.UpdateCampaignRequest
	if err := ctx.BodyParser(&body); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Validate the request body
	if validationErrors := shared.ValidateStruct(body); len(validationErrors) > 0 {
		log.Error("Validation errors:", validationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(shared.FormatValidationErrors(validationErrors))
	}

	err = h.useCase.UpdateCampaign(ctx.Context(), id, &body)
	if err != nil {
		log.Errorf("Error updating campaign: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Campaign updated successfully")
	return ctx.SendStatus(fiber.StatusAccepted)
}

func (h *CampaignHandler) DeleteCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("DeleteCampaign handler")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Errorf("Error parsing id: %v",
			err)
	}

	err = h.useCase.DeleteCampaign(ctx.Context(), id)
	if err != nil {
		log.Errorf("Error deleting campaign: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	log.Info("Campaign deleted successfully")
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *CampaignHandler) ListCampaigns(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("ListCampaigns handler")

	campaigns, err := h.useCase.ListCampaigns(ctx.Context())
	if err != nil {
		log.Errorf("Error listing campaigns: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(campaigns)
}

func (h *CampaignHandler) SearchCampaign(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("SearchCampaign handler")

	param := ctx.Query("param")
	campaigns, err := h.useCase.SearchCampaign(ctx.Context(), param)
	if err != nil {
		log.Errorf("Error searching campaigns: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(campaigns)
}
