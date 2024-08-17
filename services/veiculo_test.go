package services_test

import (
	"context"
	"testing"
	"transportadora/controller/di"
	"transportadora/models"
	"transportadora/services"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		id, err := s.InsertVeiculo(ctx, mockInsertVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestUpdateVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		id, err := s.InsertVeiculo(ctx, mockInsertVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		var mockUpdateVeiculo = models.UpdateVeiculo{
			Placa: "AAA1234",
		}

		id, err = s.UpdateVeiculo(ctx, id, mockUpdateVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		veiculo, err := s.GetVeiculo(ctx, id)
		require.NoError(t, err)

		require.Equal(t, veiculo.Placa, "AAA1234")

	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestDeleteVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		id, err := s.InsertVeiculo(ctx, mockInsertVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		err = s.DeleteVeiculo(ctx, id)
		require.NoError(t, err)

		veiculo, err := s.GetVeiculo(ctx, id)
		require.Error(t, err)
		require.Nil(t, veiculo)
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestGetVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		id, err := s.InsertVeiculo(ctx, mockInsertVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		veiculo, err := s.GetVeiculo(ctx, id)
		require.NoError(t, err)
		require.NotNil(t, veiculo)

		assert.Equal(t, veiculo.Placa, "ABC1234")
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestListVeiculos(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		for i := 0; i < 10; i++ {
			_, err := s.InsertVeiculo(ctx, mockInsertVeiculo)
			require.NoError(t, err)
		}

		var mockSearchVeiculo = models.SearchVeiculo{
			PaginationRequest: models.PaginationRequest{
				Page:  1,
				Limit: 5,
			},
		}

		list, err := s.ListVeiculos(ctx, mockSearchVeiculo)
		veiculos := *list.Data

		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Equal(t, list.Pagination.TotalRecord, 10)
		assert.Len(t, veiculos, 5)
	})
	require.NoError(t, err)

	defer d.Close(true)
}
