package server

import (
	"net/http"

	"socialNetwork/pkg/config"
)

type Route struct {
	Path    string
	handler http.HandlerFunc
	Role    uint
}

const (
	Auth uint = iota
	User
)

func (app *App) InitRoutes(conf *config.Conf) http.Handler {
	mux := http.NewServeMux()
	routes := app.createRoutes()
	for _, route := range routes {
		if requireLogin(route.Role) {
			mux.Handle(route.Path, app.SessionManager.LoadAndSave(app.authenticate(app.requireAuthentication(http.HandlerFunc(route.handler)))))
		} else {
			mux.Handle(route.Path, app.SessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(route.handler))))
		}
	}
	return mux
}

func (app *App) createRoutes() []Route {
	return []Route{
		{
			Path:    "/api/register",
			handler: app.User.Register,
			Role:    Auth,
		},
		{
			Path:    "/api/login",
			handler: app.User.Login,
			Role:    Auth,
		},
		{
			Path:    "/api/logout",
			handler: app.Logout,
			Role:    User,
		},
		{
			Path:    "/ping",
			handler: ping,
			Role:    Auth,
		},
		{
			Path:    "/api/post/create",
			handler: app.CreatePost,
			Role:    Auth,
		},
		{
			Path:    "/api/post/{id}",
			handler: app.GetPost,
			Role:    Auth,
		},
		{
			Path:    "/api/post/react",
			handler: app.ReactPost,
			Role:    Auth,
		},
	}
}

func requireLogin(Auth uint) bool {
	return Auth == User
}
