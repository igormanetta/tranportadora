package services

import (
	"context"

	"transportadora/infra/repositories"
	"transportadora/models"

	"github.com/google/uuid"
)

type motorista struct {
	repo        repositories.Motorista
	repoVeiculo repositories.Veiculo
}

func NewMotorista(repo repositories.Motorista, repoVeiculo repositories.Veiculo) *motorista {
	return &motorista{
		repo:        repo,
		repoVeiculo: repoVeiculo,
	}
}

func (s *motorista) InsertMotorista(ctx context.Context, dto models.InsertMotorista) (uuid.UUID, error) {
	return s.repo.InsertMotorista(ctx, dto)
}

func (s *motorista) UpdateMotorista(ctx context.Context, id uuid.UUID, dto models.UpdateMotorista) (uuid.UUID, error) {
	return s.repo.UpdateMotorista(ctx, id, dto)
}

func (s *motorista) DeleteMotorista(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteMotorista(ctx, id)
}

func (s *motorista) GetMotorista(ctx context.Context, id uuid.UUID) (*models.Motorista, error) {
	return s.repo.GetMotorista(ctx, id)
}

func (s *motorista) ListMotoristas(ctx context.Context, dto models.SearchMotorista) (*models.ListMotorista, error) {
	return s.repo.ListMotoristas(ctx, dto)
}

func (s *motorista) SetMotoristaVeiculo(ctx context.Context, motoristaID, veiculoID uuid.UUID) (uuid.UUID, error) {
	_, err := s.repoVeiculo.GetVeiculo(ctx, veiculoID)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = s.repo.GetMotorista(ctx, motoristaID)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.SetMotoristaVeiculo(ctx, motoristaID, veiculoID)
}
