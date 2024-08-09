package usecase

import (
	"GiftWize/src/entity/request"
	"GiftWize/src/entity/response"
	"GiftWize/src/infreaestructure/repository"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type CampaignUseCase struct {
	campaingRepo repository.CampaignRepository
}

func NewCampaignUseCase(campaingRepo repository.CampaignRepository) *CampaignUseCase {
	return &CampaignUseCase{
		campaingRepo: campaingRepo,
	}
}

func (c *CampaignUseCase) CreateCampaign(ctx context.Context, data request.CreateCampaignRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("CreateCampaign usecase")

	err := c.campaingRepo.CreateCampaign(ctx, data)
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

	err, campaign := c.campaingRepo.GetCampaign(ctx, id)
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
		Name:               campaign.Name,
		Description:        campaign.Description,
		StartDate:          campaign.StartDate.Format("2006-01-02"), // Formato de fecha
		EndDate:            campaign.EndDate.Format("2006-01-02"),   // Formato de fecha
		DiscountPercentage: fmt.Sprintf("%.2f%%", campaign.DiscountPercentage),
	}, nil
}

func (c *CampaignUseCase) UpdateCampaign(ctx context.Context, id int, data any) error {
	log := logrus.WithContext(ctx)
	log.Info("UpdateCampaign usecase")

	err := c.campaingRepo.UpdateCampaign(ctx, id, data)
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

	campaigns, err := c.campaingRepo.SearchCampaign(ctx, param)
	if err != nil {
		log.Errorf("Error searching campaign: %v", err)
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
			DiscountPercentage: fmt.Sprintf("%.2f%%", campaign.DiscountPercentage),
		})
	}

	log.Info("Campaign retrieved successfully")
	return campaignResponses, nil
}
