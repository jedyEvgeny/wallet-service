package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
)

// MockWriter эмулирует поведение при записи и чтении
type MockWriter struct {
	writeErr   error
	readAmount *int
}

func (m *MockWriter) Write(wallet *Wallet, requestID string) error {
	return m.writeErr
}

func (m *MockWriter) Read(walletID string, requestID string) (*int, error) {
	return m.readAmount, nil
}

func TestCheckPost_Success(t *testing.T) {
	mockWriter := &MockWriter{}
	service := New(mockWriter)

	reqBody := `{"valletId":"` + uuid.New().String() + `","operationType":"DEPOSIT","amount":1000}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallet", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	response, status := service.CheckPost(req, "request-id")

	if status != http.StatusCreated {
		t.Errorf("ожидался код статуса:\n\t%d\nполучили:\n\t%d\n", http.StatusCreated, status)
	}

	var resp ResponsePost
	if err := json.Unmarshal(response, &resp); err != nil {
		t.Fatalf("ошибка преобразования json-объекта: %v", err)
	}
	if resp.CodeStatus != http.StatusCreated {
		t.Errorf("ожидался код статуса:\n\t%d\nполучили:\n\t%d\n", http.StatusCreated, resp.CodeStatus)
	}
}

func TestCheckPost_WriteError(t *testing.T) {
	mockWriter := &MockWriter{
		writeErr:   errors.New("write error"),
		readAmount: new(int),
	}
	service := New(mockWriter)

	reqBody := `{"valletId":"` + uuid.New().String() + `","operationType":"WITHDRAW","amount":100}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallet", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	response, status := service.CheckPost(req, "request-id")

	if status != http.StatusInternalServerError {
		t.Errorf("ожидался код статуса:\n\t%d\nполучили:\n\t%d\n", http.StatusInternalServerError, status)
	}

	var resp ResponsePost
	if err := json.Unmarshal(response, &resp); err != nil {
		t.Fatalf("ошибка чтения json-объекта с сообщением об ошибке: %v", err)
	}
	if resp.CodeStatus != http.StatusInternalServerError {
		t.Errorf("ожидался код статуса в json-объекте:\n\t%d\nполучили:\n\t%d\n", http.StatusInternalServerError, resp.CodeStatus)
	}
}

func TestCheckGet_Success(t *testing.T) {
	amount := 1000
	mockWriter := &MockWriter{readAmount: &amount}
	service := New(mockWriter)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/"+uuid.New().String(), nil)

	response, status := service.CheckGet(req, "request-id")

	if status != http.StatusOK {
		t.Errorf("ожидался код статуса:\n\t%d\nполучили:\n\t%d\n", http.StatusOK, status)
	}

	var resp ResponseGet
	if err := json.Unmarshal(response, &resp); err != nil {
		t.Fatalf("ошибка преобразования json-объекта: %v", err)
	}
	if resp.CodeStatus != http.StatusOK {
		t.Errorf("ожидался код статуса в json-объекте:\n\t%d\nполучили:\n\t%d\n", http.StatusOK, resp.CodeStatus)
	}
}
