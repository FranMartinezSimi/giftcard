package usecase

import (
	"GiftWize/src/app" // For custom errors
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"GiftWize/src/entity/response"
	"GiftWize/src/infreaestructure/repository" // Used for ICampaignRepository
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockCampaignRepository is a mock type for the ICampaignRepository
type MockCampaignRepository struct {
	mock.Mock
}

// Ensure MockCampaignRepository implements ICampaignRepository
var _ repository.ICampaignRepository = (*MockCampaignRepository)(nil)

func (m *MockCampaignRepository) CreateCampaign(ctx context.Context, data request.CreateCampaignRequest, uuid string) error {
	args := m.Called(ctx, data, uuid)
	return args.Error(0)
}

func (m *MockCampaignRepository) GetCampaign(ctx context.Context, id int) (*models.Campaign, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) UpdateCampaign(ctx context.Context, id int, data *request.UpdateCampaignRequest) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *MockCampaignRepository) DeleteCampaign(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCampaignRepository) FullTextSearchCampaign(ctx context.Context, data *request.FullTextSearchCampaignRequest) (*response.CampaignResponse, error) {
	args := m.Called(ctx, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.CampaignResponse), args.Error(1)
}

func (m *MockCampaignRepository) SearchCampaign(ctx context.Context, query string) ([]*models.Campaign, error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) ListCampaigns(ctx context.Context) ([]*models.Campaign, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Campaign), args.Error(1)
}


func TestCampaignUseCase_UpdateCampaign(t *testing.T) {
	ctx := context.Background()
	campaignID := 1
	updateReq := &request.UpdateCampaignRequest{
		Name: "Updated Campaign",
	}

	t.Run("successful update", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(&models.Campaign{ID: uint(campaignID)}, nil).Once()
		mockRepo.On("UpdateCampaign", ctx, campaignID, updateReq).Return(nil).Once()

		err := useCase.UpdateCampaign(ctx, campaignID, updateReq)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("campaign not found for update (GetCampaign returns gorm.ErrRecordNotFound)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, gorm.ErrRecordNotFound).Once()

		err := useCase.UpdateCampaign(ctx, campaignID, updateReq)
		assert.Error(t, err)
		// This assertion depends on how CampaignUseCase handles gorm.ErrRecordNotFound.
		// Based on previous analysis, it directly returns the error from GetCampaign if err != nil.
		// If it were to translate gorm.ErrRecordNotFound to app.ErrCampaignNotFound, this would change.
		assert.Equal(t, app.ErrCampaignNotFound, err) // Assuming use case translates the error
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "UpdateCampaign", mock.Anything, mock.Anything, mock.Anything)
	})
	
	t.Run("campaign not found for update (GetCampaign returns nil, nil)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, nil).Once()

		err := useCase.UpdateCampaign(ctx, campaignID, updateReq)
		assert.Error(t, err)
		assert.Equal(t, app.ErrCampaignNotFound, err)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "UpdateCampaign", mock.Anything, mock.Anything, mock.Anything)
	})


	t.Run("error from GetCampaign (not RecordNotFound)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		dbError := errors.New("some other DB error")
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, dbError).Once()

		err := useCase.UpdateCampaign(ctx, campaignID, updateReq)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "UpdateCampaign", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("error on actual update", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		dbError := errors.New("db update error")
		mockRepo.On("GetCampaign", ctx, campaignID).Return(&models.Campaign{ID: uint(campaignID)}, nil).Once()
		mockRepo.On("UpdateCampaign", ctx, campaignID, updateReq).Return(dbError).Once()

		err := useCase.UpdateCampaign(ctx, campaignID, updateReq)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCampaignUseCase_DeleteCampaign(t *testing.T) {
	ctx := context.Background()
	campaignID := 1

	t.Run("successful delete", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(&models.Campaign{ID: uint(campaignID)}, nil).Once()
		mockRepo.On("DeleteCampaign", ctx, campaignID).Return(nil).Once()

		err := useCase.DeleteCampaign(ctx, campaignID)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("campaign not found for delete (GetCampaign returns nil, nil)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, nil).Once() 

		err := useCase.DeleteCampaign(ctx, campaignID)
		assert.Error(t, err)
		assert.Equal(t, app.ErrCampaignNotFound, err)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "DeleteCampaign", mock.Anything, mock.Anything)
	})

	t.Run("error from GetCampaign (gorm.ErrRecordNotFound)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, gorm.ErrRecordNotFound).Once()
		
		err := useCase.DeleteCampaign(ctx, campaignID)
		assert.Error(t, err)
		// Assuming use case translates gorm.ErrRecordNotFound to app.ErrCampaignNotFound
		assert.Equal(t, app.ErrCampaignNotFound, err) 
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "DeleteCampaign", mock.Anything, mock.Anything)
	})


	t.Run("error from GetCampaign (not RecordNotFound)", func(t *testing.T) {
		mockRepo := new(MockCampaignRepository)
		useCase := NewCampaignUseCase(mockRepo)
		dbError := errors.New("some other DB error for get")
		mockRepo.On("GetCampaign", ctx, campaignID).Return(nil, dbError).Once()

		err := useCase.DeleteCampaign(ctx, campaignID)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "DeleteCampaign", mock.Anything, mock.Anything)
	})

	t.Run("error on actual delete", func(t *testing.T) {
		mockRepoCampaign := new(MockCampaignRepository) // Corrected mock type
		useCase := NewCampaignUseCase(mockRepoCampaign)  // Use the correct mock

		dbError := errors.New("db delete error")
		mockRepoCampaign.On("GetCampaign", ctx, campaignID).Return(&models.Campaign{ID: uint(campaignID)}, nil).Once()
		mockRepoCampaign.On("DeleteCampaign", ctx, campaignID).Return(dbError).Once()
		
		err := useCase.DeleteCampaign(ctx, campaignID)
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		mockRepoCampaign.AssertExpectations(t)
	})
}
