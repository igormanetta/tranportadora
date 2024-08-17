package pgstore

import (
	"context"
	"fmt"
	"strings"
	"transportadora/infra/pgstore"
	"transportadora/models"
	"transportadora/utils"

	"github.com/google/uuid"
)

type veiculo struct {
	q *pgstore.Queries
}

func NewVeiculo(q *pgstore.Queries) *veiculo {
	return &veiculo{q: q}
}

func (r *veiculo) InsertVeiculo(ctx context.Context, dto models.InsertVeiculo) (uuid.UUID, error) {
	return r.q.InsertVeiculo(ctx, dto.Placa)
}

func (r *veiculo) UpdateVeiculo(ctx context.Context, id uuid.UUID, dto models.UpdateVeiculo) (uuid.UUID, error) {
	params := pgstore.UpdateVeiculoParams{
		ID:    id,
		Placa: dto.Placa,
	}

	return r.q.UpdateVeiculo(ctx, params)
}

func (r *veiculo) DeleteVeiculo(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVeiculo(ctx, id)
}

func (r *veiculo) GetVeiculo(ctx context.Context, id uuid.UUID) (*models.Veiculo, error) {
	veiculo, err := r.q.GetVeiculo(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Veiculo{
		ID:    veiculo.ID.String(),
		Placa: veiculo.Placa,
	}, nil
}

func (r *veiculo) ListVeiculos(ctx context.Context, dto models.SearchVeiculo) (*models.ListVeiculo, error) {

	limitOffset := pgstore.MountLimitOffset(dto.Limit, dto.Page)
	where := r.mountWhere(dto)

	sql := "SELECT id, placa FROM veiculo"

	query := sql + where + limitOffset
	pagination, err := pgstore.Pagination(sql, where, dto.Limit, dto.Page, r.q, ctx)
	if err != nil {
		return nil, err
	}

	veiculos, err := r.q.ListVeiculos(ctx, query)
	if err != nil {
		return nil, err
	}

	toMap := func(veiculo pgstore.Veiculo) (models.Veiculo, error) {
		return models.Veiculo{
			ID:    veiculo.ID.String(),
			Placa: veiculo.Placa,
		}, nil
	}

	veiculosDto, _ := utils.Map(veiculos, toMap)

	list := models.ListVeiculo{
		Pagination: *pagination,
		Data:       &veiculosDto,
	}

	return &list, nil
}

func (r *veiculo) mountWhere(dto models.SearchVeiculo) string {
	var where string

	if dto.Placa != "" {
		where += fmt.Sprintf(" and placa = '%s'", strings.ToUpper(dto.Placa))
	}

	if where != "" {
		where = " where" + where[4:]
	}

	return where
}
