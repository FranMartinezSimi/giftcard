package app

import (
	"GiftWize/src/entity/response"
	"context"
)

type CampaignUseCase interface {
	CreateCampaign(ctx context.Context) error
	GetCampaign(ctx context.Context, id int) (*response.CampaignResponse, error)
	UpdateCampaign(ctx context.Context, id int) error
}
