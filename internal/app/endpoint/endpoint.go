package endpoint

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Prepearer interface {
	Check(*http.Request, string) ([]byte, int)
}

type Endpoint struct {
	prepareResponse Prepearer
}

const (
	msgRequest = "[%s] Получен запрос с методом: %s от URL: %s\n"
)

func New(c Prepearer) *Endpoint {
	return &Endpoint{
		prepareResponse: c,
	}
}

func (e *Endpoint) HandlerChangeWallet(w http.ResponseWriter, r *http.Request) {
	reqID := requestID()
	log.Printf(msgRequest, reqID, r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp, status := e.prepareResponse.Check(r, reqID)
	w.WriteHeader(status)
	w.Write(resp)
}

func (e *Endpoint) HandlerStatusWallet(w http.ResponseWriter, r *http.Request) {
	reqID := requestID()
	log.Printf(msgRequest, reqID, r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp, status := e.prepareResponse.Check(r, reqID)
	w.WriteHeader(status)
	w.Write(resp)
}

func requestID() string {
	return uuid.New().String()
}
