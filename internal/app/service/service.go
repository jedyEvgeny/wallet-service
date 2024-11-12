package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Writer interface {
	Write(*Wallet, string) error
}

type Service struct {
	writerData Writer
}

func New(w Writer) *Service {
	return &Service{
		writerData: w,
	}
}

func (s *Service) Check(r *http.Request, requestID string) ([]byte, int) {
	if r.Method != http.MethodPost {
		msg := fmt.Sprintf(errMethod, http.MethodPost, r.Method)
		dataJson, status := prepareResponse(
			http.StatusMethodNotAllowed, msg, requestID, uuid.Nil)
		return dataJson, status
	}
	req, err := parseRequest(r.Body)
	if err != nil {
		dataJson, status := prepareResponse(
			http.StatusBadRequest, errDecodeJson, requestID, uuid.Nil)
		return dataJson, status
	}

	err = isValidJson(req)
	if err != nil {
		dataJson, status := prepareResponse(
			http.StatusBadRequest, fmt.Sprint(err), requestID, uuid.Nil)
		return dataJson, status
	}
	fmt.Println(req) //Логика обработки базы данных

	err = s.writerData.Write(&req, requestID)
	if err != nil {
		dataJson, status := prepareResponse(
			http.StatusInternalServerError, fmt.Sprint(err), requestID, uuid.Nil)
		return dataJson, status
	}

	dataJson, status := prepareResponse(
		http.StatusCreated, msg201, requestID, req.WalletId)
	return dataJson, status
}

func prepareResponse(statusCode int, msg, requestID string, id uuid.UUID) ([]byte, int) {
	log.Printf("[%s]  %s\n", requestID, msg)
	resp := Response{
		CodeStatus:       statusCode,
		DescritionStatus: msg,
	}
	if id != uuid.Nil {
		resp.ID = &id
	}
	dataJson, err := json.Marshal(resp)
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

	if req.OperationType != Deposit && req.OperationType != Withdrow {
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
