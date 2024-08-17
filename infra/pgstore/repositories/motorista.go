package pgstore

import (
	"context"
	"fmt"
	"strings"
	"transportadora/infra/pgstore"
	"transportadora/infra/repositories"
	"transportadora/models"
	"transportadora/utils"

	"github.com/google/uuid"
)

type motorista struct {
	q           *pgstore.Queries
	repoVeiculo repositories.Veiculo
}

func NewMotorista(q *pgstore.Queries, repoVeiculo repositories.Veiculo) *motorista {
	return &motorista{
		q:           q,
		repoVeiculo: repoVeiculo,
	}
}

func (r *motorista) InsertMotorista(ctx context.Context, dto models.InsertMotorista) (uuid.UUID, error) {
	return r.q.InsertMotorista(ctx, dto.Nome)
}

func (r *motorista) UpdateMotorista(ctx context.Context, id uuid.UUID, dto models.UpdateMotorista) (uuid.UUID, error) {
	params := pgstore.UpdateMotoristaParams{
		ID:   id,
		Nome: dto.Nome,
	}

	return r.q.UpdateMotorista(ctx, params)
}

func (r *motorista) DeleteMotorista(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMotorista(ctx, id)
}

func (r *motorista) GetMotorista(ctx context.Context, id uuid.UUID) (*models.Motorista, error) {
	motorista, err := r.q.GetMotorista(ctx, id)
	if err != nil {
		return nil, err
	}

	var veiculo *models.Veiculo

	if uuid.Nil != motorista.VeiculoID {
		veiculo, err = r.repoVeiculo.GetVeiculo(ctx, motorista.VeiculoID)
		if err != nil {
			return nil, err
		}
	}

	return &models.Motorista{
		ID:      motorista.ID.String(),
		Nome:    motorista.Nome,
		Veiculo: veiculo,
	}, nil
}

func (r *motorista) ListMotoristas(ctx context.Context, dto models.SearchMotorista) (*models.ListMotorista, error) {

	limitOffset := pgstore.MountLimitOffset(dto.Limit, dto.Page)
	where := r.mountWhere(dto)

	sql := "SELECT id, nome, veiculo_id FROM motorista"

	query := sql + where + limitOffset
	pagination, err := pgstore.Pagination(sql, where, dto.Limit, dto.Page, r.q, ctx)
	if err != nil {
		return nil, err
	}

	motoristas, err := r.q.ListMotoristas(ctx, query)
	if err != nil {
		return nil, err
	}

	toMap := func(motorista pgstore.Motorista) (models.Motorista, error) {
		var veiculo *models.Veiculo

		if uuid.Nil != motorista.VeiculoID {
			veiculo, err = r.repoVeiculo.GetVeiculo(ctx, motorista.VeiculoID)
			if err != nil {
				return models.Motorista{}, err
			}
		}

		return models.Motorista{
			ID:      motorista.ID.String(),
			Nome:    motorista.Nome,
			Veiculo: veiculo,
		}, nil
	}

	motoristasDto, err := utils.Map(motoristas, toMap)
	if err != nil {
		return nil, err
	}

	return &models.ListMotorista{
		Pagination: *pagination,
		Data:       &motoristasDto,
	}, nil
}

func (r *motorista) GetMotoristasByVeiculo(ctx context.Context, veiculoID uuid.UUID) ([]models.Motorista, error) {
	motoristas, err := r.q.GetMotoristaByVeiculo(ctx, veiculoID)
	if err != nil {
		return nil, err
	}

	toMap := func(motorista pgstore.GetMotoristaByVeiculoRow) (models.Motorista, error) {
		return models.Motorista{
			ID:   motorista.ID.String(),
			Nome: motorista.Nome,
		}, nil
	}

	motoristasDto, err := utils.Map(motoristas, toMap)
	if err != nil {
		return nil, err
	}

	return motoristasDto, nil
}

func (r *motorista) SetMotoristaVeiculo(ctx context.Context, motoristaID, veiculoID uuid.UUID) (uuid.UUID, error) {
	params := pgstore.SetMotoristaVeiculoParams{
		VeiculoID: veiculoID,
		ID:        motoristaID,
	}
	return r.q.SetMotoristaVeiculo(ctx, params)
}

func (r *motorista) mountWhere(dto models.SearchMotorista) string {
	var where string

	if dto.Nome != "" {
		where += fmt.Sprintf(" and placa = '%s'", strings.ToUpper(dto.Nome))
	}

	if where != "" {
		where = " where" + where[4:]
	}

	return where
}
