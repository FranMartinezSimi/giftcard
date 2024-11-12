package repository

import (
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"GiftWize/src/entity/response"
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	gorm *gorm.DB
}

func NewCampaignRepository(gorm *gorm.DB) *CampaignRepository {
	return &CampaignRepository{gorm: gorm}
}

func (c *CampaignRepository) CreateCampaign(ctx context.Context, data request.CreateCampaignRequest, uuid string) error {
	log.WithContext(ctx).Info("CreateCampaign repository")

	res := c.gorm.WithContext(ctx).Create(&models.Campaign{
		CampaignUUID:       uuid,
		Name:               data.Name,
		Description:        data.Description,
		StartDate:          data.StartDate,
		EndDate:            data.EndDate,
		DiscountPercentage: data.DiscountPercentage,
	})

	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error creating campaign: %v", res.Error)
		return res.Error
	}

	log.WithContext(ctx).Info("Campaign created successfully")
	return nil
}

func (c *CampaignRepository) GetCampaign(ctx context.Context, id int) (*models.Campaign, error) {
	log.WithContext(ctx).Info("GetCampaign repository")

	var campaign models.Campaign
	res := c.gorm.WithContext(ctx).First(&campaign, id)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Errorf("Campaign not found: %v", res.Error)
			return nil, nil
		}
		log.WithContext(ctx).Errorf("Error getting campaign: %v", res.Error)
		return nil, res.Error
	}

	log.WithContext(ctx).Info("Campaign retrieved successfully")
	return &campaign, nil
}

func (c *CampaignRepository) UpdateCampaign(ctx context.Context, id int, data *request.UpdateCampaignRequest) error {
	log.WithContext(ctx).Info("UpdateCampaign repository")

	updateData := map[string]interface{}{
		"name":                data.Name,
		"description":         data.Description,
		"start_date":          data.StartDate,
		"end_date":            data.EndDate,
		"is_enabled":          data.IsEnabled,
		"discount_percentage": data.DiscountPercentage,
	}

	res := c.gorm.WithContext(ctx).Model(&models.Campaign{}).Where("id = ?", id).Updates(updateData)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error updating campaign: %v", res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.WithContext(ctx).Error("Cannot update campaign")
		return nil
	}

	log.WithContext(ctx).Info("Campaign updated successfully")
	return nil
}

func (c *CampaignRepository) FullTextSearchCampaign(ctx context.Context, data *request.FullTextSearchCampaignRequest) (*response.CampaignResponse, error) {
	log.WithContext(ctx).Info("FullTextSearchCampaign repository")

	var campaign models.Campaign
	res := c.gorm.WithContext(ctx).Where("id = ? OR name = ? OR description = ? OR start_date = ? OR end_date = ?",
		data.ID, data.Name, data.Description, data.StartDate, data.EndDate).First(&campaign)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Error("Campaign not found")
			return nil, nil
		}
		log.WithContext(ctx).Errorf("Error searching campaign: %v", res.Error)
		return nil, res.Error
	}

	campaignResponse := &response.CampaignResponse{
		ID:                 campaign.ID,
		Name:               campaign.Name,
		Description:        campaign.Description,
		StartDate:          campaign.StartDate.Format("2006-01-02"),
		EndDate:            campaign.EndDate.Format("2006-01-02"),
		DiscountPercentage: fmt.Sprintf("%.2f", campaign.DiscountPercentage),
		IsEnabled:          campaign.IsEnabled,
		CreatedAt:          campaign.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	log.WithContext(ctx).Info("Campaign retrieved successfully")
	return campaignResponse, nil
}

func (c *CampaignRepository) DeleteCampaign(ctx context.Context, id int) error {
	log.WithContext(ctx).Info("DeleteCampaign repository")

	res := c.gorm.WithContext(ctx).Delete(&models.Campaign{}, id)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error deleting campaign: %v", res.Error)
		return res.Error
	}

	log.WithContext(ctx).Info("Campaign deleted successfully")
	return nil
}

func (c *CampaignRepository) SearchCampaign(ctx context.Context, query string) ([]*models.Campaign, error) {
	log.WithContext(ctx).Info("SearchCampaign repository")

	var campaigns []*models.Campaign
	res := c.gorm.WithContext(ctx).Where("MATCH(name, description, start_date, end_date, discount_percentage) AGAINST (?)", query).Find(&campaigns)

	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error searching campaign: %v", res.Error)
		return nil, res.Error
	}

	if len(campaigns) == 0 {
		log.WithContext(ctx).Info("No campaigns found")
		return nil, nil
	}

	log.WithContext(ctx).Info("Campaigns retrieved successfully")
	return campaigns, nil
}

func (c *CampaignRepository) ListCampaigns(ctx context.Context) ([]*models.Campaign, error) {
	log.WithContext(ctx).Info("ListCampaigns repository")

	var campaigns []*models.Campaign
	res := c.gorm.WithContext(ctx).Find(&campaigns)

	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error listing campaigns: %v", res.Error)
		return nil, res.Error
	}

	if len(campaigns) == 0 {
		log.WithContext(ctx).Info("No campaigns found")
		return nil, nil
	}

	log.WithContext(ctx).Info("Campaigns retrieved successfully")
	return campaigns, nil
}
