package request

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// Re-using the validator instance from campaign_request_test.go
// if these tests are in the same package, it's already initialized.
// If in a different package, validator.New() would be needed here too.
// For this structure, they are in the same 'request' package.
// var validate *validator.Validate // Already declared in campaign_request_test.go

func TestCreateGiftCardRequest_Validation(t *testing.T) {
	// Initializing validate here if it's not accessible or if running tests in isolation
	// For this subtask, assuming it's accessible from campaign_request_test.go's init()
	// If not, uncomment:
	// validate = validator.New() 

	futureDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")

	tests := []struct {
		name          string
		request       CreateGiftCardRequest
		expectedError bool
		errorFields   []string
	}{
		{
			name: "valid request",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        100.00,
				ExpirationDate: futureDate,
				Status:         "active",
				IsPromotional:  false,
				CampaignID:     1,
			},
			expectedError: false,
		},
		{
			name: "valid request - promotional without campaign_id", // campaign_id is omitempty, gt=0
			request: CreateGiftCardRequest{
				Type:           "physical",
				Balance:        50.00,
				ExpirationDate: futureDate,
				Status:         "active",
				IsPromotional:  true,
				// CampaignID is 0 (omitted)
			},
			expectedError: false,
		},
		{
			name: "missing type",
			request: CreateGiftCardRequest{
				Balance:        100.00,
				ExpirationDate: futureDate,
				Status:         "active",
			},
			expectedError: true,
			errorFields:   []string{"Type"},
		},
		{
			name: "missing balance", // Balance is float64, 'required' means it can't be zero value if not using pointers.
									 // The validation `min=0` allows 0. Let's test if required means it must be present.
									 // Actually, for float64, 'required' means it cannot be the zero value (0.0).
									 // If 0 is a valid balance that needs to be explicitly set, the field should be a pointer or use `isset` tag.
									 // Given `min=0`, 0 is allowed. The `required` tag on a non-pointer float64 means it must not be 0.
									 // This might be a slight conflict or nuance in validation tags.
									 // Let's assume "required" means it must be explicitly provided if it were a pointer.
									 // For a non-pointer, it means it cannot be its zero value (0 for float64).
									 // If a balance of 0 is valid and distinct from "not provided", then `*float64` would be better.
									 // Test assuming current struct: required means not 0.
			request: CreateGiftCardRequest{
				Type:           "virtual",
				ExpirationDate: futureDate,
				Status:         "active",
				// Balance is 0.0
			},
			expectedError: true, 
			errorFields:   []string{"Balance"}, // This will fail if 'required' on float64 allows 0. It typically doesn't.
		},
        {
            name: "balance is exactly 0 (allowed by min=0, but required might make it fail)",
            request: CreateGiftCardRequest{
                Type:           "virtual",
                Balance:        0, // Explicitly 0
                ExpirationDate: futureDate,
                Status:         "active",
            },
            // If 'required' on float64 means 'not the zero value (0.0)', this will fail.
            // If 'min=0' takes precedence or 'required' is for presence (for pointers), it will pass.
            // Validator behavior: 'required' for non-pointer numeric types means != 0.
            expectedError: true, 
            errorFields:   []string{"Balance"},
        },
		{
			name: "balance less than zero",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        -10.00,
				ExpirationDate: futureDate,
				Status:         "active",
			},
			expectedError: true,
			errorFields:   []string{"Balance"},
		},
		{
			name: "missing expiration date",
			request: CreateGiftCardRequest{
				Type:    "virtual",
				Balance: 100.00,
				Status:  "active",
			},
			expectedError: true,
			errorFields:   []string{"ExpirationDate"},
		},
		{
			name: "missing status",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        100.00,
				ExpirationDate: futureDate,
			},
			expectedError: true,
			errorFields:   []string{"Status"},
		},
		{
			name: "invalid campaign_id (less than 1)",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        100.00,
				ExpirationDate: futureDate,
				Status:         "active",
				CampaignID:     0, // Invalid because of gt=0, but omitempty should make it pass if it's the zero value
			},
			// This should NOT be an error because CampaignID has `omitempty` and 0 is the zero value for uint.
			// If CampaignID was set to 0 explicitly by the user, it would be considered "omitted" by validator.
			expectedError: false, 
		},
        {
			name: "invalid campaign_id (explicitly 0 but gt=0)",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        100.00,
				ExpirationDate: futureDate,
				Status:         "active",
				CampaignID:     0, // Explicitly 0, should be omitted by `omitempty`
			},
			expectedError: false, // `omitempty` means if it's the zero value, validation for gt=0 is skipped
		},
         {
			name: "valid campaign_id (greater than 0)",
			request: CreateGiftCardRequest{
				Type:           "virtual",
				Balance:        10.00, // Not 0, so "required" passes for Balance
				ExpirationDate: futureDate,
				Status:         "active",
				CampaignID:     123, 
			},
			expectedError: false, 
		},
		{
			name: "multiple errors",
			request: CreateGiftCardRequest{
				// Type missing
				// Balance negative
				Balance: -5.0,
				// ExpirationDate missing
				// Status missing
				CampaignID: 0, // This is fine due to omitempty
			},
			expectedError: true,
			errorFields:   []string{"Type", "Balance", "ExpirationDate", "Status"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ensure validator is initialized for each run if tests run in parallel or specific conditions
			 v := validator.New()
			err := v.Struct(tt.request)
			if tt.expectedError {
				assert.Error(t, err, "Expected validation error for test: %s", tt.name)
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok, "Error should be of type validator.ValidationErrors for test: %s", tt.name)

				foundFields := make(map[string]bool)
				for _, e := range validationErrors {
					foundFields[e.Field()] = true
				}
				for _, expectedField := range tt.errorFields {
					assert.True(t, foundFields[expectedField], "Expected error on field %s for test '%s'", expectedField, tt.name)
				}
			} else {
				assert.NoError(t, err, "Expected no validation error for test: %s", tt.name)
			}
		})
	}
}

func TestUpdateGiftCardRequest_Validation(t *testing.T) {
	// validate = validator.New() // Ensure initialized if needed
	futureDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")

	tests := []struct {
		name          string
		request       UpdateGiftCardRequest
		expectedError bool
		errorFields   []string
	}{
		{
			name: "valid update request",
			request: UpdateGiftCardRequest{
				Type:           "physical",
				Balance:        75.00,
				ExpirationDate: futureDate,
				Status:         "inactive",
				IsPromotional:  true,
			},
			expectedError: false,
		},
		{
			name: "update missing type",
			request: UpdateGiftCardRequest{
				Balance:        75.00,
				ExpirationDate: futureDate,
				Status:         "inactive",
			},
			expectedError: true,
			errorFields:   []string{"Type"},
		},
		{
            name: "update balance is exactly 0 (fails 'required')",
            request: UpdateGiftCardRequest{
                Type:           "virtual",
                Balance:        0, 
                ExpirationDate: futureDate,
                Status:         "active",
            },
            expectedError: true, 
            errorFields:   []string{"Balance"},
        },
		{
			name: "update balance less than zero",
			request: UpdateGiftCardRequest{
				Type:           "physical",
				Balance:        -5.00,
				ExpirationDate: futureDate,
				Status:         "inactive",
			},
			expectedError: true,
			errorFields:   []string{"Balance"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 v := validator.New()
			err := v.Struct(tt.request)
			if tt.expectedError {
				assert.Error(t, err, "Expected validation error for test: %s", tt.name)
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok, "Error should be of type validator.ValidationErrors for test: %s", tt.name)
				
				foundFields := make(map[string]bool)
				for _, e := range validationErrors {
					foundFields[e.Field()] = true
				}
				for _, expectedField := range tt.errorFields {
					assert.True(t, foundFields[expectedField], "Expected error on field %s for test '%s'", expectedField, tt.name)
				}
			} else {
				assert.NoError(t, err, "Expected no validation error for test: %s", tt.name)
			}
		})
	}
}
