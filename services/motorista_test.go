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

func TestInsertMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		id, err := s.InsertMotorista(ctx, mockInsertMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestUpdateMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		id, err := s.InsertMotorista(ctx, mockInsertMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		var mockUpdateMotorista = models.UpdateMotorista{
			Nome: "Maria",
		}

		id, err = s.UpdateMotorista(ctx, id, mockUpdateMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		motorista, err := s.GetMotorista(ctx, id)
		require.NoError(t, err)
		require.Equal(t, "Maria", motorista.Nome)
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestDeleteMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		id, err := s.InsertMotorista(ctx, mockInsertMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		err = s.DeleteMotorista(ctx, id)
		require.NoError(t, err)

		_, err = s.GetMotorista(ctx, id)
		require.Error(t, err)
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestGetMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		id, err := s.InsertMotorista(ctx, mockInsertMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, id)

		motorista, err := s.GetMotorista(ctx, id)
		require.NoError(t, err)
		require.NotNil(t, motorista)

		assert.Equal(t, motorista.Nome, "João")
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestListMotoristas(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	err := d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		for i := 0; i < 10; i++ {
			_, err := s.InsertMotorista(ctx, mockInsertMotorista)
			require.NoError(t, err)
		}

		var mockSearchMotorista = models.SearchMotorista{
			PaginationRequest: models.PaginationRequest{
				Page:  1,
				Limit: 5,
			},
		}

		list, err := s.ListMotoristas(ctx, mockSearchMotorista)
		motoristas := *list.Data

		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Equal(t, list.Pagination.TotalRecord, 10)
		require.Len(t, motoristas, 5)
	})
	require.NoError(t, err)
}

func TestSetMotoristaVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	ctx := context.Background()

	var motoristaID uuid.UUID
	var veiculoID uuid.UUID
	var err error

	err = d.Dig.Invoke(func(s services.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		veiculoID, err = s.InsertVeiculo(ctx, mockInsertVeiculo)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, veiculoID)
	})
	require.NoError(t, err)

	err = d.Dig.Invoke(func(s services.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		motoristaID, err = s.InsertMotorista(ctx, mockInsertMotorista)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, motoristaID)

		motoristaID, err = s.SetMotoristaVeiculo(ctx, motoristaID, veiculoID)
		require.NoError(t, err)
		require.IsType(t, uuid.UUID{}, motoristaID)
	})
	require.NoError(t, err)

	defer d.Close(true)
}
