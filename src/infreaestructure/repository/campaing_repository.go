package repository

import (
	"GiftWize/src/entity/models"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	gorm *gorm.DB
}

func NewCampaignRepository() *CampaignRepository {
	return &CampaignRepository{}
}

func (c *CampaignRepository) CreateCampaign(ctx context.Context) error {
	log := logrus.WithContext(ctx)
	log.Info("CreateCampaign repository")

	err := c.gorm.Create(&models.Campaign{})
	if err.Error != nil {
		log.Errorf("Error creating campaign: %v", err.Error)
		return err.Error
	}
	log.Info("Campaign created successfully")
	return nil
}

func (c *CampaignRepository) GetCampaign(ctx context.Context, data interface{}) error {
	log := logrus.WithContext(ctx)
	log.Info("GetCampaign repository")

	err := c.gorm.Find(&models.Campaign{})
	if err.Error != nil {
		log.Errorf("Error getting campaign: %v", err.Error)
		return err.Error
	}
	log.Info("Campaign retrieved successfully")
	return nil
}

func (c *CampaignRepository) UpdateCampaign(ctx context.Context, id int) error {
	log := logrus.WithContext(ctx)
	log.Info("UpdateCampaign repository")

	err := c.gorm.Model(&models.Campaign{}).Where("id = ?", id).Updates(&models.Campaign{})
	if err.Error != nil {
		log.Errorf("Error updating campaign: %v", err.Error)
		return err.Error
	}
	log.Info("Campaign updated successfully")
	return nil
}
