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

func TestHandleInsertMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleInsertMotoristaError(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "AA",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status bad request; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleUpdateMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		var mockUpdateMotorista = models.UpdateMotorista{
			Nome: "João",
		}

		jsonData, err = json.Marshal(mockUpdateMotorista)
		require.NoError(t, err)
		body = tests.ByteToReadCloser(jsonData)

		req = httptest.NewRequest("PUT", "/motorista/"+returnID.ID.String(), body)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleDeleteMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		req = httptest.NewRequest("DELETE", "/motorista/"+returnID.ID.String(), nil)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusNoContent {
			t.Errorf("expected status no content; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleGetMotorista(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		req = httptest.NewRequest("GET", "/motorista/"+returnID.ID.String(), nil)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleListMotoristas(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		req := httptest.NewRequest("GET", "/motorista", nil)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}

func TestHandleSetMotoristaVeiculo(t *testing.T) {
	d := di.New()
	d.Inject(true)

	err := d.Dig.Invoke(func(c *controller.Motorista) {
		var mockInsertMotorista = models.InsertMotorista{
			Nome: "João",
		}

		jsonData, err := json.Marshal(mockInsertMotorista)
		require.NoError(t, err)
		body := tests.ByteToReadCloser(jsonData)

		req := httptest.NewRequest("POST", "/motorista", body)
		w := httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res := w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnID := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnID)
		require.NoError(t, err)

		var mockInsertVeiculo = models.InsertVeiculo{
			Placa: "ABC1234",
		}

		jsonData, err = json.Marshal(mockInsertVeiculo)
		require.NoError(t, err)
		body = tests.ByteToReadCloser(jsonData)

		req = httptest.NewRequest("POST", "/veiculo", body)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status created; got %v", res.Status)
		}

		returnIDVeiculo := models.ReturnID{}
		err = json.NewDecoder(res.Body).Decode(&returnIDVeiculo)
		require.NoError(t, err)

		req = httptest.NewRequest("PATCH", "/motorista/"+returnID.ID.String()+"/veiculo/"+returnIDVeiculo.ID.String(), nil)
		w = httptest.NewRecorder()

		c.API.R.ServeHTTP(w, req)

		res = w.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
	require.NoError(t, err)
}
