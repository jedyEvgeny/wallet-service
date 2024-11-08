package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT" env-default:"7777" env-description:"Порт сервера"`
		Host string `env:"SERVER_HOST" env-default:"localhost" env-description:"Хост сервера"`
	}
	Database struct {
		Port string `env:"DB_PORT" env-default:"7777" env-description:"Порт базы данных"`
		Host string `env:"DB_HOST" env-default:"localhost" env-description:"Хост базы данных"`
	}
}

type Wallet struct {
	WalletId      uuid.UUID `json:"valletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}

type Response struct {
	StatusCode    int        `json:"statusCode"`
	DescritionErr string     `json:"descriptionError"`
	ID            *uuid.UUID `json:"resourceID,omitempty"`
}

const (
	pathPost      = "/api/v1/wallet"
	pathGet       = "/api/v1/wallets/"
	pathDirConfig = "/home/ev/Документы/Практика/JavaCode/.env"
)

const (
	msgRequest = "[%s] Получен запрос с методом: %s от URL: %s\n"
	msg201     = "Ресурс создан"
)

const (
	errMarshalJson = "ошибка создания json-объекта"
	errDecodeJson  = "ошибка декодирования json-объекта"
	errMethod      = "ошибка метода. Ожидался: %s, имеется: %s"
	errIsNotUUID   = "поле json valletID ожидалось с уникальным UUID. Имеется: %v"
	errOperation   = "поле json operationType: %s. Ожидалось 'DEPOSIT' или 'WITHDRAW'"
	errAmount      = "поле json amount должно быть больше нуля. Имеется: %d"
)

func main() {
	cfg := mustReadConfig()
	http.HandleFunc(pathPost, HandlerChangeWallet)
	// http.HandleFunc(pathGet, HandleStatusWallet)

	fmt.Printf("Запустили сервер на хосте: %s и порту: %s\n%s\n",
		cfg.Server.Host, cfg.Server.Port, strings.Repeat("-", 16))

	err := http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mustReadConfig() Config {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFilePath := filepath.Join(workingDir, ".env")

	var cfg Config
	err = cleanenv.ReadConfig(envFilePath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func HandlerChangeWallet(w http.ResponseWriter, r *http.Request) {
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
		StatusCode:    statusCode,
		DescritionErr: msg,
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
