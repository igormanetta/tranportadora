package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transportadora/controller"
	"transportadora/controller/di"
	"transportadora/models"
	"transportadora/tests"

	"github.com/stretchr/testify/require"
)

func TestHandleInsertVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		jsonData, err := json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/veiculo", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestHandleInsertVeiculoError(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC12345",
		}

		jsonData, err := json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/veiculo", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status bad request; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestHandleUpdateVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		jsonData, err := json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/veiculo", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		var mockUpdateVeiculo = models.UpdateVeiculo{
			Placa: "AAA1234",
		}

		jsonData, err = json.Marshal(mockUpdateVeiculo)
		require.NoError(t, err)
		body = tests.ByteToReadCloser(jsonData)

		req = httptest.NewRequest("PUT", "/veiculo/"+returnID.ID.String(), body)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestHandleDeleteVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		jsonData, err := json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/veiculo", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		req = httptest.NewRequest("DELETE", "/veiculo/"+returnID.ID.String(), nil)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusNoContent {
			t.Errorf("expected status no content; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestHandleGetVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		jsonData, err := json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/veiculo", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		req = httptest.NewRequest("GET", "/veiculo/"+returnID.ID.String(), nil)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}

func TestHandleListVeiculos(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Veiculo) {
		req := httptest.NewRequest("GET", "/veiculo", nil)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)

	defer d.Close(true)
}
