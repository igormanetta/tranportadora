package services

import (
	"context"
	"transportadora/infra/repositories"
	"transportadora/models"
	"transportadora/utils"

	"github.com/google/uuid"
)

type veiculo struct {
	repo          repositories.Veiculo
	repoMotorista repositories.Motorista
}

func NewVeiculo(repo repositories.Veiculo, repoMotorista repositories.Motorista) *veiculo {
	return &veiculo{
		repo:          repo,
		repoMotorista: repoMotorista,
	}
}

func (s *veiculo) InsertVeiculo(ctx context.Context, dto models.InsertVeiculo) (uuid.UUID, error) {
	return s.repo.InsertVeiculo(ctx, dto)
}

func (s *veiculo) UpdateVeiculo(ctx context.Context, id uuid.UUID, dto models.UpdateVeiculo) (uuid.UUID, error) {
	return s.repo.UpdateVeiculo(ctx, id, dto)
}

func (s *veiculo) DeleteVeiculo(ctx context.Context, id uuid.UUID) error {

	motoristas, err := s.repoMotorista.GetMotoristasByVeiculo(ctx, id)
	if err != nil {
		return err
	}

	if len(motoristas) > 0 {
		return utils.ErrMotAssociado
	}

	return s.repo.DeleteVeiculo(ctx, id)
}

func (s *veiculo) GetVeiculo(ctx context.Context, id uuid.UUID) (*models.Veiculo, error) {
	return s.repo.GetVeiculo(ctx, id)
}

func (s *veiculo) ListVeiculos(ctx context.Context, dto models.SearchVeiculo) (*models.ListVeiculo, error) {
	return s.repo.ListVeiculos(ctx, dto)
}
