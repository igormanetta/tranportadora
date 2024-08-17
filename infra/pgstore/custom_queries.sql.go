package pgstore

import (
	"context"
)

func (q *Queries) ListVeiculos(ctx context.Context, sql string) ([]Veiculo, error) {
	rows, err := q.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Veiculo
	for rows.Next() {
		var i Veiculo
		if err := rows.Scan(
			&i.ID,
			&i.Placa,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *Queries) ListMotoristas(ctx context.Context, sql string) ([]Motorista, error) {
	rows, err := q.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Motorista
	for rows.Next() {
		var i Motorista
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.VeiculoID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
