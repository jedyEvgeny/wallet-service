package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jedyEvgeny/wallet-service/internal/app/endpoint"
	"github.com/jedyEvgeny/wallet-service/internal/app/service"
	"github.com/jedyEvgeny/wallet-service/internal/config"
	storage "github.com/jedyEvgeny/wallet-service/internal/storage/postgresql"
)

type route struct {
	PostWallet   string
	StatusWallet string
}

type App struct {
	cfg      *config.Config
	db       *storage.DataBase
	route    *route
	service  *service.Service
	endpoint *endpoint.Endpoint
}

func New() (*App, error) {
	a := &App{}
	a.route = createRoute()
	a.cfg = config.MustLoad()
	a.db = storage.MustNew(a.cfg)
	a.service = service.New(a.db)
	a.endpoint = endpoint.New(a.service)

	return a, nil
}

func createRoute() *route {
	return &route{
		PostWallet:   "/api/v1/wallet",
		StatusWallet: "/api/v1/wallets/",
	}
}

func (a *App) Run() error {
	defer func() {
		if err := a.db.Close(); err != nil {
			log.Printf("Ошибка при закрытии базы данных: %v", err)
		}
	}()

	a.configureRoutes()

	log.Printf("Запустили сервер на хосте: %s и порту: %s\n%s\n",
		a.cfg.Server.Host, a.cfg.Server.Port, strings.Repeat("-", 70))

	err := http.ListenAndServe(a.serverAdress(), nil)
	if err != nil {
		return fmt.Errorf("ошибка прослушивания порта: %w", err)
	}
	return nil
}

func (a *App) serverAdress() string {
	return a.cfg.Server.Host + ":" + a.cfg.Server.Port
}
