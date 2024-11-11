package app

import "net/http"

func (a *App) configureRoutes() {
	http.HandleFunc(a.route.PostWallet, a.endpoint.HandlerChangeWallet)
	http.HandleFunc(a.route.StatusWallet, a.endpoint.HandlerStatusWallet)
}
