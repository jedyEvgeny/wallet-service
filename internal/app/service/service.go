package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Writer interface {
	Write(*Wallet, string) error
	Read(string, string) (*int, error)
}

type Service struct {
	writerData Writer
}

func New(w Writer) *Service {
	return &Service{
		writerData: w,
	}
}

func (s *Service) CheckPost(r *http.Request, requestID string) ([]byte, int) {
	if r.Method != http.MethodPost {
		msg := fmt.Sprintf(errMethod, http.MethodPost, r.Method)
		dataJson, status := prepareResponsePost(
			http.StatusMethodNotAllowed, msg, requestID, uuid.Nil)
		return dataJson, status
	}
	req, err := parseRequest(r.Body)
	if err != nil {
		dataJson, status := prepareResponsePost(
			http.StatusBadRequest, errDecodeJson, requestID, uuid.Nil)
		return dataJson, status
	}

	err = isValidJson(req)
	if err != nil {
		dataJson, status := prepareResponsePost(
			http.StatusBadRequest, fmt.Sprint(err), requestID, uuid.Nil)
		return dataJson, status
	}
	err = s.writerData.Write(&req, requestID)
	if err != nil {
		dataJson, status := prepareResponsePost(
			http.StatusInternalServerError, fmt.Sprint(err), requestID, uuid.Nil)
		return dataJson, status
	}

	dataJson, status := prepareResponsePost(
		http.StatusCreated, msg201, requestID, req.WalletId)
	return dataJson, status
}

func (s *Service) CheckGet(r *http.Request, requestID string) ([]byte, int) {
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf(errMethod, http.MethodGet, r.Method)
		dataJson, status := prepareResponseGet(
			http.StatusMethodNotAllowed, msg, requestID, nil)
		return dataJson, status
	}

	id, err := parseURL(r)
	if err != nil {
		dataJson, status := prepareResponseGet(
			http.StatusBadRequest, errIsNotUUIDInURL, requestID, nil)
		return dataJson, status
	}

	amount, err := s.writerData.Read(id, requestID)
	if err != nil {
		dataJson, status := prepareResponseGet(
			http.StatusInternalServerError, fmt.Sprint(err), requestID, nil)
		return dataJson, status
	}
	dataJson, status := prepareResponseGet(
		http.StatusOK, msg200, requestID, amount)
	return dataJson, status
}

func prepareResponsePost(statusCode int, msg, requestID string, id uuid.UUID) ([]byte, int) {
	log.Printf("[%s]  %s\n", requestID, msg)
	resp := ResponsePost{
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

func parseURL(r *http.Request) (string, error) {
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/wallets/") //Рефакторить

	uuid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		return "", fmt.Errorf(errIsNotUUID, id)
	}
	fmt.Println("ID: ", uuid)
	return id, nil
}

func prepareResponseGet(statusCode int, msg, requestID string, amount *int) ([]byte, int) {
	log.Printf("[%s]  %s\n", requestID, msg)
	resp := ResponseGet{
		CodeStatus:       statusCode,
		DescritionStatus: msg,
	}
	if amount != nil {
		resp.Amount = amount
	}
	dataJson, err := json.Marshal(resp)
	if err != nil {
		log.Println(errMarshalJson)
		return nil, http.StatusInternalServerError
	}
	return dataJson, statusCode
}
