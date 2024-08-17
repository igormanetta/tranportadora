package services

import (
	"context"
	"transportadora/models"

	"github.com/google/uuid"
)

type Motorista interface {
	InsertMotorista(ctx context.Context, dto models.InsertMotorista) (uuid.UUID, error)
	UpdateMotorista(ctx context.Context, id uuid.UUID, dto models.UpdateMotorista) (uuid.UUID, error)
	DeleteMotorista(ctx context.Context, id uuid.UUID) error

	GetMotorista(ctx context.Context, id uuid.UUID) (*models.Motorista, error)
	ListMotoristas(ctx context.Context, dto models.SearchMotorista) (*models.ListMotorista, error)

	SetMotoristaVeiculo(ctx context.Context, motoristaID, veiculoID uuid.UUID) (uuid.UUID, error)
}
