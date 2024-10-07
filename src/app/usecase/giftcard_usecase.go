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

type GiftCardUseCase struct {
	giftCardRepo repository.GiftCardRepository
}

func NewGiftCardUseCase(giftCardRepo repository.GiftCardRepository) *GiftCardUseCase {
	return &GiftCardUseCase{
		giftCardRepo: giftCardRepo,
	}
}

func (g *GiftCardUseCase) CreateGiftCard(ctx context.Context, data request.CreateGiftCardRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("CreateGiftCard use case")

	UUID, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("Error generating UUID: %v", err)
		return err
	}
	giftCardUUID := fmt.Sprintf("GIFT-%d", UUID)

	err = g.giftCardRepo.CreateGiftCard(ctx, data, giftCardUUID)
	if err != nil {
		log.Errorf("Error creating gift card: %v", err)
		return err
	}

	log.Info("Gift card created successfully")
	return nil
}

func (g *GiftCardUseCase) GetAllGiftCardList(ctx context.Context) ([]response.GetAllGiftCardResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("GetAllGiftCardList use case")

	results, err := g.giftCardRepo.GetAllGiftCardList(ctx)
	if err != nil {
		log.Errorf("Error getting gift card list: %v", err)
		return []response.GetAllGiftCardResponse{}, err
	}
	log.Info("Lista de tarjetas de regalo recuperada exitosamente")

	var responseList []response.GetAllGiftCardResponse
	for _, giftCard := range results {
		responseItem := response.GetAllGiftCardResponse{
			ID:             giftCard.ID,
			Type:           giftCard.Type,
			Balance:        giftCard.Balance,
			ExpirationDate: giftCard.ExpirationDate.Format("2006-01-02"),
			Status:         giftCard.Status,
			IsPromotional:  giftCard.IsPromotional,
		}
		responseList = append(responseList, responseItem)
	}

	return responseList, nil
}

func (g *GiftCardUseCase) GetGiftCardByID(ctx context.Context, id string) (response.GetAllGiftCardResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("GetGiftCardByID use case")

	giftCard, err := g.giftCardRepo.GetGiftCardByID(ctx, id)
	if err != nil {
		log.Errorf("Error getting gift card: %v", err)
		return response.GetAllGiftCardResponse{}, err
	}

	responseItem := response.GetAllGiftCardResponse{
		ID:             giftCard.ID,
		Type:           giftCard.Type,
		Balance:        giftCard.Balance,
		ExpirationDate: giftCard.ExpirationDate.Format("2006-01-02"),
		Status:         giftCard.Status,
		IsPromotional:  giftCard.IsPromotional,
	}

	log.Info("Gift card retrieved successfully")
	return responseItem, nil
}

func (g *GiftCardUseCase) UpdateGiftCard(ctx context.Context, id string, data request.UpdateGiftCardRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("UpdateGiftCard use case")

	err := g.giftCardRepo.UpdateGiftCard(ctx, id, data)
	if err != nil {
		log.Errorf("Error updating gift card: %v", err)
		return err
	}

	log.Info("Gift card updated successfully")

	return nil
}

func (g *GiftCardUseCase) FullTextSearchGiftCard(ctx context.Context, query string) ([]response.GetAllGiftCardResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("FullTextSearchGiftCard use case")

	results, err := g.giftCardRepo.FullTextSearchGiftCard(ctx, query)
	if err != nil {
		log.Errorf("Error searching gift card: %v", err)
		return []response.GetAllGiftCardResponse{}, err
	}

	var responseList []response.GetAllGiftCardResponse
	for _, giftCard := range results {
		responseItem := response.GetAllGiftCardResponse{
			ID:             giftCard.ID,
			Type:           giftCard.Type,
			Balance:        giftCard.Balance,
			ExpirationDate: giftCard.ExpirationDate.Format("2006-01-02"),
			Status:         giftCard.Status,
			IsPromotional:  giftCard.IsPromotional,
		}
		responseList = append(responseList, responseItem)
	}

	log.Info("Gift card search results retrieved successfully")
	return responseList, nil
}

func (g *GiftCardUseCase) DeleteGiftCard(ctx context.Context, id string) error {
	log := logrus.WithContext(ctx)
	log.Info("DeleteGiftCard use case")

	err := g.giftCardRepo.DeleteGiftCard(ctx, id)
	if err != nil {
		log.Errorf("Error deleting gift card: %v", err)
		return err
	}

	log.Info("Gift card deleted successfully")
	return nil
}
