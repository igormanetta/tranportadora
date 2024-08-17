package services

import (
	"context"
	"transportadora/models"

	"github.com/google/uuid"
)

type Veiculo interface {
	InsertVeiculo(ctx context.Context, dto models.InsertVeiculo) (uuid.UUID, error)
	UpdateVeiculo(ctx context.Context, id uuid.UUID, dto models.UpdateVeiculo) (uuid.UUID, error)
	DeleteVeiculo(ctx context.Context, id uuid.UUID) error

	GetVeiculo(ctx context.Context, id uuid.UUID) (*models.Veiculo, error)
	ListVeiculos(ctx context.Context, dto models.SearchVeiculo) (*models.ListVeiculo, error)
}
