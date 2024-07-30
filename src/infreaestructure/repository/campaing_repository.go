package repository

import (
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CampaignRepositoryInterface interface {
	CreateCampaign(ctx context.Context) error
	GetCampaign(ctx context.Context, data interface{}) (error, *models.Campaign)
	UpdateCampaign(ctx context.Context, id int) error
	SearchFullCampaign(ctx context.Context, data interface{}) error
	DeleteCampaign(ctx context.Context, id int) error
}

type CampaignRepository struct {
	gorm *gorm.DB
}

func NewCampaignRepository() *CampaignRepository {
	return &CampaignRepository{}
}

func (c *CampaignRepository) CreateCampaign(ctx context.Context, data request.CreateCampaignRequest) error {
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

func (c *CampaignRepository) GetCampaign(ctx context.Context, id int) (error, *models.Campaign) {
	log := logrus.WithContext(ctx)
	log.Info("GetCampaign repository")

	data := c.gorm.Find(&models.Campaign{}, id).First(&models.Campaign{}, id)

	if data.Error != nil {
		log.Errorf("Error getting campaign: %v", data.Error)
		return data.Error, nil
	}

	log.Info("Campaign retrieved successfully")
	return nil, &models.Campaign{}
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

func (c *CampaignRepository) SearchFullCampaign(ctx context.Context, data interface{}) error {
	log := logrus.WithContext(ctx)
	log.Info("SearchFullCampaign repository")

	err := c.gorm.Find(&models.Campaign{}).Where("id = ? OR name = ? OR description = ? OR start_date = ? OR end_date = ? OR discount_percentage = ?", data)
	if err.Error != nil {
		log.Errorf("Error searching full campaign: %v", err.Error)
		return err.Error
	}
	log.Info("Full campaign retrieved successfully")
	return nil

}

func (c *CampaignRepository) DeleteCampaign(ctx context.Context, id int) error {
	log := logrus.WithContext(ctx)
	log.Info("DeleteCampaign repository")

	err := c.gorm.Delete(&models.Campaign{}, id)
	if err.Error != nil {
		log.Errorf("Error deleting campaign: %v", err.Error)
		return err.Error
	}
	log.Info("Campaign deleted successfully")
	return nil
}

func (c *CampaignRepository) SearchCampaign(ctx context.Context, query string) ([]*models.Campaign, error) {
	log := logrus.WithContext(ctx)
	log.Info("SearchCampaign repository")

	var campaigns []*models.Campaign

	if err := c.gorm.Where("MATCH(name, description, start_date, end_date, discount_percentage) AGAINST (?)", query).Find(&campaigns).Error; err != nil {
		log.Errorf("Error searching campaign: %v", err)
		return nil, err
	}

	if len(campaigns) == 0 {
		log.Error("Campaign not found")
		return nil, nil
	}

	log.Info("Campaign retrieved successfully")
	return campaigns, nil
}
