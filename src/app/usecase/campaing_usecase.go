package usecase

import (
	"GiftWize/src/infreaestructure/repository"
	"context"

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

func (c *CampaignUseCase) CreateCampaign(ctx context.Context) error {
	log := logrus.WithContext(ctx)
	log.Info("CreateCampaign usecase")

	err := c.campaingRepo.CreateCampaign(ctx)
	if err != nil {
		log.Errorf("Error creating campaign: %v", err)
		return err
	}
	log.Info("Campaign created successfully")
	return nil
}


