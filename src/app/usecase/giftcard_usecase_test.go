package usecase

import (
	"GiftWize/src/app" // For custom errors
	"GiftWize/src/entity/models"
	"GiftWize/src/entity/request"
	"GiftWize/src/entity/response"
	"GiftWize/src/infreaestructure/repository" // Used for IGiftCardRepository
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockGiftCardRepository is a mock type for the IGiftCardRepository
type MockGiftCardRepository struct {
	mock.Mock
}

// Ensure MockGiftCardRepository implements IGiftCardRepository
var _ repository.IGiftCardRepository = (*MockGiftCardRepository)(nil)

func (m *MockGiftCardRepository) GiftCardNumberExists(ctx context.Context, giftCardNumber string) (bool, error) {
	args := m.Called(ctx, giftCardNumber)
	return args.Bool(0), args.Error(1)
}

func (m *MockGiftCardRepository) CreateGiftCard(ctx context.Context, data request.CreateGiftCardRequest, code string, giftCardNumber string) error {
	args := m.Called(ctx, data, code, giftCardNumber)
	return args.Error(0)
}

func (m *MockGiftCardRepository) GetGiftCardByCode(ctx context.Context, code string) (*models.GiftCard, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GiftCard), args.Error(1)
}

func (m *MockGiftCardRepository) GetByGiftCardNumber(ctx context.Context, giftCardNumber string) (*models.GiftCard, error) {
	args := m.Called(ctx, giftCardNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GiftCard), args.Error(1)
}

func (m *MockGiftCardRepository) GetAllGiftCardList(ctx context.Context) ([]models.GiftCard, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1) 
	}
	val, ok := args.Get(0).([]models.GiftCard)
	if !ok && args.Get(0) != nil { 
		// Consider logging or returning an error if type assertion fails unexpectedly
	}
	return val, args.Error(1)
}

func (m *MockGiftCardRepository) UpdateGiftCard(ctx context.Context, code string, data request.UpdateGiftCardRequest) error {
	args := m.Called(ctx, code, data)
	return args.Error(0)
}

func (m *MockGiftCardRepository) UpdateGiftCardBalanceAndStatus(ctx context.Context, code string, balance float64, status string) error {
	args := m.Called(ctx, code, balance, status)
	return args.Error(0)
}

func (m *MockGiftCardRepository) FullTextSearchGiftCard(ctx context.Context, query string) ([]models.GiftCard, error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	val, ok := args.Get(0).([]models.GiftCard)
	if !ok && args.Get(0) != nil {
		// Consider logging or returning an error
	}
	return val, args.Error(1)
}

func (m *MockGiftCardRepository) DeleteGiftCard(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}


func TestGiftCardUseCase_UseGiftCardAmount(t *testing.T) {
	ctx := context.Background()

	activeStatus := "active"
	expiredStatus := "expired"
	usedStatus := "used"
	inactiveStatus := "inactive"

	validCardCode := "test-code-123"
	validCardNumber := "GC1234567890123456"

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	tests := []struct {
		name                    string
		giftCardNumber          string
		amountToUse             float64
		mockSetup               func(mockRepo *MockGiftCardRepository) 
		expectedResponse        response.UseGiftCardAmountResponse
		expectedError           error
		expectUpdateCall        bool 
		expectedNewBalance      float64
		expectedNewStatus       string
	}{
		{
			name:           "successful use of gift card",
			giftCardNumber: validCardNumber,
			amountToUse:    50.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
					Code:           validCardCode, 
					GiftCardNumber: validCardNumber,
					Balance:        100.0,
					Status:         activeStatus,
					ExpirationDate: tomorrow,
				}, nil).Once()
				mockRepo.On("UpdateGiftCardBalanceAndStatus", ctx, validCardCode, 50.0, activeStatus).Return(nil).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 50.0, IsUsed: true, Message: "Gift card amount used successfully."},
			expectedError:    nil,
			expectUpdateCall: true,
			expectedNewBalance: 50.0, 
			expectedNewStatus: activeStatus,
		},
		{
			name:           "gift card not found",
			giftCardNumber: "GCUNKNOWN",
			amountToUse:    10.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, "GCUNKNOWN").Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 0, IsUsed: false, Message: "Gift card not found."},
			expectedError:    app.ErrGiftCardNotFound,
			expectUpdateCall: false,
		},
		{
			name:           "gift card not active",
			giftCardNumber: validCardNumber,
			amountToUse:    10.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
					Code:           validCardCode, 
					GiftCardNumber: validCardNumber,
					Balance:        100.0,
					Status:         inactiveStatus, 
					ExpirationDate: tomorrow,
				}, nil).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 100.0, IsUsed: false, Message: "Gift card is not active. Status: inactive."},
			expectedError:    app.ErrGiftCardNotActive,
			expectUpdateCall: false,
		},
		{
			name:           "gift card expired",
			giftCardNumber: validCardNumber,
			amountToUse:    10.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
					Code:           validCardCode, 
					GiftCardNumber: validCardNumber,
					Balance:        100.0,
					Status:         activeStatus,
					ExpirationDate: yesterday, 
				}, nil).Once()
				mockRepo.On("UpdateGiftCardBalanceAndStatus", ctx, validCardCode, 100.0, expiredStatus).Return(nil).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 100.0, IsUsed: false, Message: "Gift card has expired."},
			expectedError:    app.ErrGiftCardExpired,
			expectUpdateCall: true,
			expectedNewBalance: 100.0, 
			expectedNewStatus: expiredStatus,
		},
		{
			name:           "insufficient balance",
			giftCardNumber: validCardNumber,
			amountToUse:    150.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
					Code:           validCardCode, 
					GiftCardNumber: validCardNumber,
					Balance:        100.0, 
					Status:         activeStatus,
					ExpirationDate: tomorrow,
				}, nil).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 100.0, IsUsed: false, Message: "Insufficient balance."},
			expectedError:    app.ErrInsufficientBalance,
			expectUpdateCall: false,
		},
		{
			name:           "using exact balance",
			giftCardNumber: validCardNumber,
			amountToUse:    100.0,
			mockSetup: func(mockRepo *MockGiftCardRepository) {
				mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
					Code:           validCardCode, 
					GiftCardNumber: validCardNumber,
					Balance:        100.0,
					Status:         activeStatus,
					ExpirationDate: tomorrow,
				}, nil).Once()
				mockRepo.On("UpdateGiftCardBalanceAndStatus", ctx, validCardCode, 0.0, usedStatus).Return(nil).Once()
			},
			expectedResponse: response.UseGiftCardAmountResponse{Balance: 0.0, IsUsed: true, Message: "Gift card amount used successfully."},
			expectedError:    nil,
			expectUpdateCall: true,
			expectedNewBalance: 0.0, 
			expectedNewStatus: usedStatus,
		},
        {
            name:           "error from GetByGiftCardNumber (not RecordNotFound)",
            giftCardNumber: "GCERROR",
            amountToUse:    10.0,
            mockSetup: func(mockRepo *MockGiftCardRepository) {
                mockRepo.On("GetByGiftCardNumber", ctx, "GCERROR").Return(nil, errors.New("generic DB error")).Once()
            },
            expectedResponse: response.UseGiftCardAmountResponse{Balance: 0, IsUsed: false, Message: "Error retrieving gift card."},
            expectedError:    errors.New("generic DB error"),
            expectUpdateCall: false,
        },
        {
            name:           "error updating status when card expires",
            giftCardNumber: validCardNumber,
            amountToUse:    10.0,
            mockSetup: func(mockRepo *MockGiftCardRepository) {
                mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
                    Code:           validCardCode, 
                    GiftCardNumber: validCardNumber,
                    Balance:        100.0,
                    Status:         activeStatus,
                    ExpirationDate: yesterday, 
                }, nil).Once()
                mockRepo.On("UpdateGiftCardBalanceAndStatus", ctx, validCardCode, 100.0, expiredStatus).Return(errors.New("update failed")).Once();
            },
            expectedResponse: response.UseGiftCardAmountResponse{Balance: 100.0, IsUsed: false, Message: "Gift card has expired."},
            expectedError:    app.ErrGiftCardExpired, 
			expectUpdateCall: true,
			expectedNewBalance: 100.0,
			expectedNewStatus: expiredStatus,
        },
		{
            name:           "error from UpdateGiftCardBalanceAndStatus on successful use",
            giftCardNumber: validCardNumber,
            amountToUse:    50.0,
            mockSetup: func(mockRepo *MockGiftCardRepository) {
                mockRepo.On("GetByGiftCardNumber", ctx, validCardNumber).Return(&models.GiftCard{
                    Code:           validCardCode, 
                    GiftCardNumber: validCardNumber,
                    Balance:        100.0,
                    Status:         activeStatus,
                    ExpirationDate: tomorrow,
                }, nil).Once()
                mockRepo.On("UpdateGiftCardBalanceAndStatus", ctx, validCardCode, 50.0, activeStatus).Return(errors.New("update failed")).Once()
            },
            expectedResponse: response.UseGiftCardAmountResponse{Balance: 100.0, IsUsed: false, Message: "Failed to update gift card after use."}, 
            expectedError:    errors.New("update failed"),
			expectUpdateCall: true,
			expectedNewBalance: 50.0,
			expectedNewStatus: activeStatus,
        },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockGiftCardRepository) 
			useCase := NewGiftCardUseCase(mockRepo) 
			tt.mockSetup(mockRepo)

			resp, err := useCase.UseGiftCardAmount(ctx, tt.giftCardNumber, tt.amountToUse)

			assert.Equal(t, tt.expectedResponse, resp, "Response struct does not match for test: %s", tt.name)
			
			if tt.expectedError != nil {
				assert.Error(t, err, "Expected an error for test: %s", tt.name)
				assert.True(t, errors.Is(err, tt.expectedError) || err.Error() == tt.expectedError.Error(), "Error type mismatch for test: %s. Expected: %v, Got: %v", tt.name, tt.expectedError, err)
			} else {
				assert.NoError(t, err, "Did not expect an error for test: %s", tt.name)
			}
			
			if tt.expectUpdateCall { 
				mockRepo.AssertCalled(t, "UpdateGiftCardBalanceAndStatus", ctx, validCardCode, tt.expectedNewBalance, tt.expectedNewStatus)
			} else {
				mockRepo.AssertNotCalled(t, "UpdateGiftCardBalanceAndStatus", mock.AnythingOfTypeArgument("context.backgroundCtx"), mock.AnythingOfTypeArgument("string"), mock.AnythingOfTypeArgument("float64"), mock.AnythingOfTypeArgument("string"))
			}
			mockRepo.AssertExpectations(t) 
		})
	}
}

func TestGiftCardUseCase_UpdateGiftCard(t *testing.T) {
    ctx := context.Background()
    testCode := "test-code-for-update" 
    updateReq := request.UpdateGiftCardRequest{
        Type:    "virtual",
        Balance: 50.0,
        Status:  "inactive",
    }

    t.Run("successful update", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(&models.GiftCard{Code: testCode}, nil).Once() 
        mockRepo.On("UpdateGiftCard", ctx, testCode, updateReq).Return(nil).Once()

        err := useCase.UpdateGiftCard(ctx, testCode, updateReq)
        assert.NoError(t, err)
        mockRepo.AssertExpectations(t)
    })

	t.Run("update returns error", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(&models.GiftCard{Code: testCode}, nil).Once() 
        mockRepo.On("UpdateGiftCard", ctx, testCode, updateReq).Return(errors.New("db update error")).Once()

        err := useCase.UpdateGiftCard(ctx, testCode, updateReq)
        assert.Error(t, err)
        assert.EqualError(t, err, "db update error")
        mockRepo.AssertExpectations(t)
    })

    t.Run("gift card not found for update", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(nil, gorm.ErrRecordNotFound).Once() 

        err := useCase.UpdateGiftCard(ctx, testCode, updateReq)
        assert.Error(t, err)
        assert.Equal(t, app.ErrGiftCardNotFound, err)
        mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "UpdateGiftCard", mock.Anything, mock.Anything, mock.Anything)
    })

	t.Run("error from GetGiftCardByCode (not RecordNotFound)", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(nil, errors.New("other db error")).Once() 

        err := useCase.UpdateGiftCard(ctx, testCode, updateReq)
        assert.Error(t, err)
        assert.EqualError(t, err, "other db error")
        mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "UpdateGiftCard", mock.Anything, mock.Anything, mock.Anything)
    })
}


func TestGiftCardUseCase_DeleteGiftCard(t *testing.T) {
    ctx := context.Background()
    testCode := "test-code-for-delete" 

    t.Run("successful delete", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(&models.GiftCard{Code: testCode}, nil).Once() 
        mockRepo.On("DeleteGiftCard", ctx, testCode).Return(nil).Once()

        err := useCase.DeleteGiftCard(ctx, testCode)
        assert.NoError(t, err)
        mockRepo.AssertExpectations(t)
    })

	t.Run("delete returns error", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(&models.GiftCard{Code: testCode}, nil).Once() 
        mockRepo.On("DeleteGiftCard", ctx, testCode).Return(errors.New("db delete error")).Once()

        err := useCase.DeleteGiftCard(ctx, testCode)
        assert.Error(t, err)
        assert.EqualError(t, err, "db delete error")
        mockRepo.AssertExpectations(t)
    })

    t.Run("gift card not found for delete", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(nil, gorm.ErrRecordNotFound).Once() 

        err := useCase.DeleteGiftCard(ctx, testCode)
        assert.Error(t, err)
        assert.Equal(t, app.ErrGiftCardNotFound, err)
        mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "DeleteGiftCard", mock.Anything, mock.Anything)
    })

	t.Run("error from GetGiftCardByCode (not RecordNotFound) on delete", func(t *testing.T) {
		mockRepo := new(MockGiftCardRepository)
		useCase := NewGiftCardUseCase(mockRepo)
        mockRepo.On("GetGiftCardByCode", ctx, testCode).Return(nil, errors.New("another db error")).Once() 

        err := useCase.DeleteGiftCard(ctx, testCode)
        assert.Error(t, err)
        assert.EqualError(t, err, "another db error")
        mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "DeleteGiftCard", mock.Anything, mock.Anything)
    })
}
