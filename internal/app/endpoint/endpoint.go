package endpoint

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type CheckPrepearer interface {
	Check()
	Prepare()
}

type Endpoint struct {
	checkPrepare CheckPrepearer
}

func New(c CheckPrepearer) *Endpoint {
	return &Endpoint{
		checkPrepare: c,
	}
}

// PrepareResponse
// ParseRequest - преобразуем json
// ValidJson
func (e *Endpoint) HandlerChangeWallet(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	log.Printf(msgRequest, requestID, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method != http.MethodPost {
		msg := fmt.Sprintf(errMethod, http.MethodPost, r.Method)
		dataJson, status := prepareResponse(
			http.StatusMethodNotAllowed, msg, requestID, uuid.Nil)
		w.WriteHeader(status)
		w.Write(dataJson)
		return
	}
	req, err := parseRequest(r.Body)
	if err != nil {
		dataJson, status := prepareResponse(
			http.StatusBadRequest, errDecodeJson, requestID, uuid.Nil)
		w.WriteHeader(status)
		w.Write(dataJson)
		return
	}

	err = isValidJson(req)
	if err != nil {
		dataJson, status := prepareResponse(
			http.StatusBadRequest, fmt.Sprint(err), requestID, uuid.Nil)
		w.WriteHeader(status)
		w.Write(dataJson)
		return
	}
	fmt.Println(req) //Логика обработки базы данных

	dataJson, status := prepareResponse(
		http.StatusCreated, msg201, requestID, req.WalletId)
	w.WriteHeader(status)
	w.Write(dataJson)
}

func prepareResponse(statusCode int, msg, requestID string, id uuid.UUID) ([]byte, int) {
	log.Printf("[%s]  %s\n", requestID, msg)
	msgErr := Response{
		CodeStatus:       statusCode,
		DescritionStatus: msg,
	}
	if id != uuid.Nil {
		msgErr.ID = &id
	}
	dataJson, err := json.Marshal(msgErr)
	if err != nil {
		log.Println(errMarshalJson)
		return nil, http.StatusInternalServerError
	}
	return dataJson, statusCode
}

func parseRequest(body io.ReadCloser) (Wallet, error) {
	var req Wallet
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return Wallet{}, err
	}
	return req, nil
}

func isValidJson(req Wallet) error {
	var err error

	_, parseErr := uuid.Parse(req.WalletId.String())
	if req.WalletId == uuid.Nil || parseErr != nil {
		msgErr := fmt.Sprintf(errIsNotUUID, req.WalletId)
		err = errors.New(msgErr)
	}

	if req.OperationType != "DEPOSIT" && req.OperationType != "WITHDRAW" {
		if err != nil {
			err = fmt.Errorf("%w; %w", err, fmt.Errorf(errOperation, req.OperationType))
		}
		if err == nil {
			err = fmt.Errorf(errOperation, req.OperationType)
		}
	}

	if req.Amount <= 0 {
		if err != nil {
			err = fmt.Errorf("%w; %w", err, fmt.Errorf(errAmount, req.Amount))
		}
		if err == nil {
			err = fmt.Errorf(errAmount, req.Amount)
		}
	}
	return err
}
