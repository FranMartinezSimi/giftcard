package repository

import (
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type GiftCardRepository struct {
	gorm *gorm.DB
}

func NewGiftCardRepository(gorm *gorm.DB) *GiftCardRepository {
	return &GiftCardRepository{gorm: gorm}
}

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

func (c *GiftCardRepository) GetGiftCardByID(ctx context.Context, id string) (models.GiftCard, error) {
	log.WithContext(ctx).Info("GetGiftCardByUUID repository")

	var giftCard models.GiftCard
	res := c.gorm.WithContext(ctx).Where("uuid = ?", id).First(&giftCard)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error getting gift card: %v", res.Error)
		return models.GiftCard{}, res.Error
	}

	log.WithContext(ctx).Info("Gift card retrieved successfully")
	return giftCard, nil
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

func (c *GiftCardRepository) UpdateGiftCard(ctx context.Context, id string, data request.UpdateGiftCardRequest) error {
	log.WithContext(ctx).Info("UpdateGiftCard repository")

	var giftCard models.GiftCard
	res := c.gorm.WithContext(ctx).Where("id = ?", id).First(&giftCard)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error getting gift card: %v", res.Error)
		return res.Error
	}

	giftCard.Type = data.Type
	giftCard.Balance = data.Balance
	expirationDate, expirationErr := time.Parse("2006-01-02", data.ExpirationDate)
	if expirationErr != nil {
		log.WithContext(ctx).Error("update-giftcard-repository Error parsing ExpirationDate:", expirationErr)
		return expirationErr
	}
	giftCard.ExpirationDate = expirationDate
	giftCard.ExpirationDate = expirationDate
	giftCard.Status = data.Status
	giftCard.IsPromotional = data.IsPromotional

	res = c.gorm.WithContext(ctx).Model(&giftCard).Updates(giftCard)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error updating gift card: %v", res.Error)
		return res.Error
	}

	log.WithContext(ctx).Info("Gift card updated successfully")
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

	res := c.gorm.WithContext(ctx).Delete(&models.GiftCard{}, "uuid = ?", id)
	if res.Error != nil {
		log.WithContext(ctx).Errorf("Error eliminando tarjeta de regalo: %v", res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.WithContext(ctx).Warn("No se encontró la tarjeta de regalo para eliminar")
		return gorm.ErrRecordNotFound
	}

	log.WithContext(ctx).Info("Tarjeta de regalo eliminada exitosamente")
	return nil
}
