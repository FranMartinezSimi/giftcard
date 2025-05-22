package repository

import (
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"context"
	"errors" // Import the errors package
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// IGiftCardRepository defines the interface for gift card repository operations.
type IGiftCardRepository interface {
	GiftCardNumberExists(ctx context.Context, giftCardNumber string) (bool, error)
	CreateGiftCard(ctx context.Context, data request.CreateGiftCardRequest, uuid string, giftCardNumber string) error
	GetGiftCardByCode(ctx context.Context, code string) (*models.GiftCard, error)
	GetByGiftCardNumber(ctx context.Context, giftCardNumber string) (*models.GiftCard, error)
	GetAllGiftCardList(ctx context.Context) ([]models.GiftCard, error)
	UpdateGiftCard(ctx context.Context, code string, data request.UpdateGiftCardRequest) error
	UpdateGiftCardBalanceAndStatus(ctx context.Context, code string, balance float64, status string) error
	FullTextSearchGiftCard(ctx context.Context, query string) ([]models.GiftCard, error)
	DeleteGiftCard(ctx context.Context, id string) error
}

type GiftCardRepository struct {
	gorm *gorm.DB
}

// NewGiftCardRepository creates a new instance of GiftCardRepository.
// It now returns IGiftCardRepository.
func NewGiftCardRepository(gorm *gorm.DB) IGiftCardRepository {
	return &GiftCardRepository{gorm: gorm}
}

// Ensure GiftCardRepository implements IGiftCardRepository
var _ IGiftCardRepository = (*GiftCardRepository)(nil)

func (c *GiftCardRepository) GiftCardNumberExists(ctx context.Context, giftCardNumber string) (bool, error) {
	log.WithContext(ctx).Info("GiftCardNumberExists repository")

	var count int64
	res := c.gorm.WithContext(ctx).Model(&models.GiftCard{}).Where("gift_card_number = ?", giftCardNumber).Count(&count)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error checking gift card number existence: %v", res.Error)
		return false, res.Error
	}

	log.WithContext(ctx).Info("Gift card number existence checked successfully")
	return count > 0, nil
}

func (c *GiftCardRepository) CreateGiftCard(ctx context.Context, data request.CreateGiftCardRequest, uuid string, giftCardNumber string) error {
	log.WithContext(ctx).Info("CreateGiftCard repository")

	expirationDate, expirationErr := time.Parse("2006-01-02", data.ExpirationDate)
	if expirationErr != nil {
		log.WithContext(ctx).Error("create-giftcard-repository Error parsing ExpirationDate:", expirationErr)
		return expirationErr
	}

	res := c.gorm.WithContext(ctx).Create(&models.GiftCard{
		Code:           uuid, // Assign the incoming uuid to the Code field
		Type:           data.Type,
		GiftCardNumber: giftCardNumber,
		Balance:        data.Balance,
		ExpirationDate: expirationDate,
		Status:         data.Status,
		IsPromotional:  data.IsPromotional,
		CampaignID:     data.CampaignID,
	})

	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error creating gift card: %v", res.Error)
		return res.Error
	}

	log.WithContext(ctx).Info("Gift card created successfully")
	return nil
}

func (c *GiftCardRepository) GetGiftCardByCode(ctx context.Context, code string) (*models.GiftCard, error) {
	log.WithContext(ctx).Infof("GetGiftCardByCode repository for code: %s", code)

	var giftCard models.GiftCard
	res := c.gorm.WithContext(ctx).Where("code = ?", code).First(&giftCard)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Warnf("Gift card with code %s not found: %v", code, res.Error)
			return nil, gorm.ErrRecordNotFound // Return specific error
		}
		log.WithContext(ctx).Errorf("Error getting gift card by code %s: %v", code, res.Error)
		return nil, res.Error
	}

	log.WithContext(ctx).Info("Gift card retrieved successfully")
	return &giftCard, nil
}

// GetByGiftCardNumber retrieves a gift card by its number.
// Returns gorm.ErrRecordNotFound if not found.
func (c *GiftCardRepository) GetByGiftCardNumber(ctx context.Context, giftCardNumber string) (*models.GiftCard, error) {
	log.WithContext(ctx).Infof("GetByGiftCardNumber repository for number: %s", giftCardNumber)
	var giftCard models.GiftCard
	res := c.gorm.WithContext(ctx).Where("gift_card_number = ?", giftCardNumber).First(&giftCard)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Warnf("Gift card with number %s not found: %v", giftCardNumber, res.Error)
			return nil, gorm.ErrRecordNotFound
		}
		log.WithContext(ctx).Errorf("Error getting gift card by number %s: %v", giftCardNumber, res.Error)
		return nil, res.Error
	}
	log.WithContext(ctx).Infof("Gift card with number %s retrieved successfully", giftCardNumber)
	return &giftCard, nil
}

func (c *GiftCardRepository) GetAllGiftCardList(ctx context.Context) ([]models.GiftCard, error) {
	log.WithContext(ctx).Info("GetAllGiftCardList repository")

	var giftCards []models.GiftCard
	res := c.gorm.WithContext(ctx).Find(&giftCards)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error getting gift card list: %v", res.Error)
		return []models.GiftCard{}, res.Error
	}

	log.WithContext(ctx).Info("Gift card list retrieved successfully")
	return giftCards, nil
}

func (c *GiftCardRepository) UpdateGiftCard(ctx context.Context, code string, data request.UpdateGiftCardRequest) error {
	log.WithContext(ctx).Infof("UpdateGiftCard repository for code: %s", code)

	// Fields to update
	updateFields := map[string]interface{}{}

	if data.Type != "" {
		updateFields["type"] = data.Type
	}
	// Assuming Balance can be 0, so we update it if it's part of the request.
	// If Balance is a pointer in UpdateGiftCardRequest, we can check for nil.
	// For now, always including it if the request means to set it.
	updateFields["balance"] = data.Balance

	if data.ExpirationDate != "" {
		expirationDate, err := time.Parse("2006-01-02", data.ExpirationDate)
		if err != nil {
			log.WithContext(ctx).Errorf("Error parsing ExpirationDate for code %s: %v", code, err) // Ensure 'code' is used here
			return err
		}
		updateFields["expiration_date"] = expirationDate
	}
	if data.Status != "" {
		updateFields["status"] = data.Status
	}
	// IsPromotional is a bool, so it will always have a value.
	// This means it will always be included in the update if present in the struct.
	// If partial update is needed for bools, use a pointer or specific logic.
	updateFields["is_promotional"] = data.IsPromotional
	
	// Check if there's anything to update
    if len(updateFields) == 0 {
        log.WithContext(ctx).Infof("No fields to update for gift card code %s", code) // Changed UUID to code and id to code
        return nil
    }

	res := c.gorm.WithContext(ctx).Model(&models.GiftCard{}).Where("code = ?", code).Updates(updateFields)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error updating gift card code %s: %v", code, res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.WithContext(ctx).Warnf("No gift card found with code %s to update (or no data changed)", code)
		// To differentiate between "not found" and "no data changed", a prior select might be needed.
		// For now, this is treated as not an error, but could return gorm.ErrRecordNotFound if strict "must exist" is needed.
		// However, the use case already checks for existence.
	}

	log.WithContext(ctx).Infof("Gift card code %s updated successfully or no changes needed", code)
	return nil
}

// UpdateGiftCardBalanceAndStatus updates the balance and status of a gift card.
func (c *GiftCardRepository) UpdateGiftCardBalanceAndStatus(ctx context.Context, code string, balance float64, status string) error {
	log.WithContext(ctx).Infof("UpdateGiftCardBalanceAndStatus repository for code: %s", code)

	updateFields := map[string]interface{}{
		"balance": balance,
		"status":  status,
	}

	res := c.gorm.WithContext(ctx).Model(&models.GiftCard{}).Where("code = ?", code).Updates(updateFields)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error updating balance/status for gift card code %s: %v", code, res.Error) // Ensure 'code' is used here
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.WithContext(ctx).Warnf("No gift card found with code %s to update balance/status", code) // Ensure 'code' is used here
		return gorm.ErrRecordNotFound // Explicitly return not found if no rows affected
	}
	log.WithContext(ctx).Infof("Balance and status for gift card code %s updated successfully", code) // Ensure 'code' is used here
	return nil
}


func (c *GiftCardRepository) FullTextSearchGiftCard(ctx context.Context, query string) ([]models.GiftCard, error) {
	log.WithContext(ctx).Info("FullTextSearchGiftCard repository")

	var giftCards []models.GiftCard
	res := c.gorm.WithContext(ctx).Where("MATCH(type, status) AGAINST(? IN BOOLEAN MODE)", query).Find(&giftCards)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error full text searching gift card: %v", res.Error)
		return []models.GiftCard{}, res.Error
	}

	log.WithContext(ctx).Info("Gift card list retrieved successfully")
	return giftCards, nil
}

func (c *GiftCardRepository) DeleteGiftCard(ctx context.Context, id string) error {
	log.WithContext(ctx).Info("DeleteGiftCard repository")

	res := c.gorm.WithContext(ctx).Delete(&models.GiftCard{}, "code = ?", id) // id is code here
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error deleting gift card code %s: %v", id, res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.WithContext(ctx).Warnf("No gift card found with code %s to delete", id)
		return gorm.ErrRecordNotFound // Return specific error
	}

	log.WithContext(ctx).Infof("Gift card with code %s deleted successfully", id)
	return nil
}
