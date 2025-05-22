package request

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func TestCreateCampaignRequest_Validation(t *testing.T) {
	now := time.Now()
	validStartDate := now.Add(time.Hour * 24)
	validEndDate := validStartDate.Add(time.Hour * 48)

	tests := []struct {
		name          string
		request       CreateCampaignRequest
		expectedError bool
		errorFields   []string // Fields expected to have validation errors
	}{
		{
			name: "valid request",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				Description:        "A great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				IsEnabled:          true,
				DiscountPercentage: 10.5,
			},
			expectedError: false,
		},
		{
			name: "missing name",
			request: CreateCampaignRequest{
				Description:        "A great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"Name"},
		},
		{
			name: "name too short",
			request: CreateCampaignRequest{
				Name:               "S",
				Description:        "A great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"Name"},
		},
		{
			name: "name too long",
			request: CreateCampaignRequest{
				Name:               string(make([]byte, 256)), // 256 chars
				Description:        "A great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"Name"},
		},
		{
			name: "description too long",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				Description:        string(make([]byte, 1001)), // 1001 chars
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"Description"},
		},
		{
			name: "missing start date",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				EndDate:            validEndDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"StartDate"},
		},
		{
			name: "missing end date",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				StartDate:          validStartDate,
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"EndDate"},
		},
		{
			name: "end date before start date",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				StartDate:          validStartDate,
				EndDate:            validStartDate.Add(-time.Hour * 24), // End date is before start date
				DiscountPercentage: 10.5,
			},
			expectedError: true,
			errorFields:   []string{"EndDate"},
		},
		{
			name: "missing discount percentage",
			request: CreateCampaignRequest{
				Name:      "Summer Sale",
				StartDate: validStartDate,
				EndDate:   validEndDate,
			},
			expectedError: true,
			errorFields:   []string{"DiscountPercentage"},
		},
		{
			name: "discount percentage too low",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: -5.0,
			},
			expectedError: true,
			errorFields:   []string{"DiscountPercentage"},
		},
		{
			name: "discount percentage too high",
			request: CreateCampaignRequest{
				Name:               "Summer Sale",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 101.0,
			},
			expectedError: true,
			errorFields:   []string{"DiscountPercentage"},
		},
		{
			name: "multiple errors",
			request: CreateCampaignRequest{
				Name:               "S",                                      // Too short
				StartDate:          validStartDate,                           // Valid
				EndDate:            validStartDate.Add(-time.Hour * 24),      // End before start
				DiscountPercentage: 200,                                    // Too high
			},
			expectedError: true,
			errorFields:   []string{"Name", "EndDate", "DiscountPercentage"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.request)
			if tt.expectedError {
				assert.Error(t, err, "Expected validation error")
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok, "Error should be of type validator.ValidationErrors")

				foundFields := make(map[string]bool)
				for _, e := range validationErrors {
					foundFields[e.Field()] = true
				}

				for _, expectedField := range tt.errorFields {
					assert.True(t, foundFields[expectedField], "Expected error on field %s", expectedField)
				}
				// Optional: Check if only expected fields have errors
				// assert.Equal(t, len(tt.errorFields), len(foundFields), "Mismatch in number of error fields")

			} else {
				assert.NoError(t, err, "Expected no validation error")
			}
		})
	}
}

func TestUpdateCampaignRequest_Validation(t *testing.T) {
	// Similar structure to TestCreateCampaignRequest_Validation
	// For brevity, only a few cases are shown here. A full test suite would be more comprehensive.
	now := time.Now()
	validStartDate := now.Add(time.Hour * 24)
	validEndDate := validStartDate.Add(time.Hour * 48)

	tests := []struct {
		name          string
		request       UpdateCampaignRequest
		expectedError bool
		errorFields   []string
	}{
		{
			name: "valid update request",
			request: UpdateCampaignRequest{
				Name:               "Updated Summer Sale",
				Description:        "An updated great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				IsEnabled:          false,
				DiscountPercentage: 15.0,
			},
			expectedError: false,
		},
		{
			name: "update with name too short",
			request: UpdateCampaignRequest{
				Name:               "U",
				Description:        "An updated great summer sale.",
				StartDate:          validStartDate,
				EndDate:            validEndDate,
				DiscountPercentage: 15.0,
			},
			expectedError: true,
			errorFields:   []string{"Name"},
		},
		{
			name: "update with end date before start date",
			request: UpdateCampaignRequest{
				Name:               "Updated Summer Sale",
				StartDate:          validStartDate,
				EndDate:            validStartDate.Add(-time.Hour * 24),
				DiscountPercentage: 15.0,
			},
			expectedError: true,
			errorFields:   []string{"EndDate"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.request)
			if tt.expectedError {
				assert.Error(t, err)
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				foundFields := make(map[string]bool)
				for _, e := range validationErrors {
					foundFields[e.Field()] = true
				}
				for _, expectedField := range tt.errorFields {
					assert.True(t, foundFields[expectedField], "Expected error on field %s for test '%s'", expectedField, tt.name)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
