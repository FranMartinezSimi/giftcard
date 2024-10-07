package usecase

import (
	"GiftWize/src/entity/request"
	"GiftWize/src/entity/response"
	"GiftWize/src/infreaestructure/repository"
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type CampaignUseCase struct {
	campaignRepo repository.CampaignRepository
}

func NewCampaignUseCase(campaignRepo repository.CampaignRepository) *CampaignUseCase {
	return &CampaignUseCase{
		campaignRepo: campaignRepo,
	}
}

func (c *CampaignUseCase) CreateCampaign(ctx context.Context, data request.CreateCampaignRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("CreateCampaign use case")

	UUID, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("Error generating UUID: %v", err)
		return err
	}
	campaignUUID := fmt.Sprintf("CAMP-%d", UUID)

	err = c.campaignRepo.CreateCampaign(ctx, data, campaignUUID)
	if err != nil {
		log.Errorf("Error creating campaign: %v", err)
		return err
	}

	log.Info("Campaign created successfully")
	return nil
}

func (c *CampaignUseCase) GetCampaign(ctx context.Context, id int) (*response.CampaignResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("GetCampaign usecase")

	campaign, err := c.campaignRepo.GetCampaign(ctx, id)
	if err != nil {
		log.Errorf("Error getting campaign: %v", err)
		return nil, err
	}
	if campaign == nil {
		log.Error("Campaign not found")
		return nil, nil
	}

	return &response.CampaignResponse{
		ID:                 campaign.ID,
		Uuid:               campaign.CampaignUUID,
		Name:               campaign.Name,
		Description:        campaign.Description,
		StartDate:          campaign.StartDate.Format("2006-01-02"), // Formato de fecha
		EndDate:            campaign.EndDate.Format("2006-01-02"),   // Formato de fecha
		DiscountPercentage: fmt.Sprintf("%.2f%%", campaign.DiscountPercentage),
	}, nil
}

func (c *CampaignUseCase) UpdateCampaign(ctx context.Context, id int, data *request.UpdateCampaignRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("UpdateCampaign usecase")

	campaignFound, err := c.campaignRepo.GetCampaign(ctx, id)

	if err != nil {
		log.Errorf("Errror finding campaign: %v", err)
		return err
	}
	if campaignFound == nil {
		log.Error("Campaign not found")
		return nil
	}

	err = c.campaignRepo.UpdateCampaign(ctx, id, data)
	if err != nil {
		log.Errorf("Error updating campaign: %v", err)
		return err
	}

	log.Info("Campaign updated successfully")
	return nil
}

func (c *CampaignUseCase) SearchCampaign(ctx context.Context, param string) ([]response.CampaignResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("SearchCampaign usecase")

	campaigns, err := c.campaignRepo.SearchCampaign(ctx, param)
	if err != nil {
		log.Errorf("Error searching campaigns: %v", err)
		return nil, err
	}

	var campaignResponses []response.CampaignResponse
	for _, campaign := range campaigns {
		campaignResponses = append(campaignResponses, response.CampaignResponse{
			ID:                 campaign.ID,
			Name:               campaign.Name,
			Description:        campaign.Description,
			StartDate:          campaign.StartDate.Format("2006-01-02"), // Formato de fecha
			EndDate:            campaign.EndDate.Format("2006-01-02"),   // Formato de fecha
			IsEnabled:          campaign.IsEnabled,
			DiscountPercentage: fmt.Sprintf("%.2f%%", campaign.DiscountPercentage),
		})
	}

	log.Info("Campaigns retrieved successfully")
	return campaignResponses, nil
}

func (c *CampaignUseCase) DeleteCampaign(ctx context.Context, id int) error {
	log := logrus.WithContext(ctx)
	log.Info("DeleteCampaign usecase")

	campaignFound, err := c.campaignRepo.GetCampaign(ctx, id)
	if err != nil {
		log.Errorf("Error finding campaign: %v", err)
		return err
	}
	if campaignFound == nil {
		log.Error("Campaign not found")
		return nil
	}

	err = c.campaignRepo.DeleteCampaign(ctx, id)
	if err != nil {
		log.Errorf("Error deleting campaign: %v", err)
		return err
	}

	log.Info("Campaign deleted successfully")
	return nil
}

func (c *CampaignUseCase) ListCampaigns(ctx context.Context) ([]response.CampaignResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("ListCampaigns usecase")

	campaigns, err := c.campaignRepo.ListCampaigns(ctx)
	if err != nil {
		log.Errorf("Error listing campaigns: %v", err)
		return nil, err
	}

	var campaignResponses []response.CampaignResponse
	for _, campaign := range campaigns {
		campaignResponses = append(campaignResponses, response.CampaignResponse{
			ID:                 campaign.ID,
			Name:               campaign.Name,
			Description:        campaign.Description,
			StartDate:          campaign.StartDate.Format("2006-01-02"), // Formato de fecha
			EndDate:            campaign.EndDate.Format("2006-01-02"),   // Formato de fecha
			IsEnabled:          campaign.IsEnabled,
			DiscountPercentage: fmt.Sprintf("%.2f%%", campaign.DiscountPercentage),
		})
	}

	log.Info("Campaigns retrieved successfully")
	return campaignResponses, nil
}
