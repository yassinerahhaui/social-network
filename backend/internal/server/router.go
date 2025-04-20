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
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}

func (app *App) createRoutes() []Route {
	return []Route{
		/*testing handlers*/
		{
			Path:    "/ping",
			handler: ping,
			Role:    Auth,
		},
		{
			Path:    "/ping/user",
			handler: ping,
			Role:    User,
		},
		/*user handlers*/
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
			Path:    "/api/follow/",
			handler: app.Follow,
			Role:    User,
		},
		{
			Path:    "/api/profile",
			handler: app.Profile,
			Role:    User,
		},
		/*group handlers*/
		{
			Path:    "/api/user/groups",
			handler: app.GetUserGroups,
			Role:    User,
		},
		{
			Path:    "/api/group",
			handler: app.GetGroupById,
			Role:    User,
		},
		{
			Path:    "/api/group/create",
			handler: app.CreateGroup,
			Role:    User,
		},
		{
			Path:    "/api/groups",
			handler: app.GetAllGroups,
			Role:    User,
		},
		{
			Path:    "/api/group/members",
			handler: app.GetGroupMembers,
			Role:    User,
		},
		{
			Path: "/api/group/invite/response",
			handler: app.InvitationResponse,
			Role:    User,
		},
		{
			Path: "/api/group/invite/request",
			handler: app.InviteToJoinGroup,
			Role:    User,
		},
		{
			Path: "/api/group/join/request",
			handler: app.RequestToJoinGroup,
			Role:    User,
		},
		{
			Path: "/api/group/join/response",
			handler: app.RequestToJoinResponse,
			Role:		User,
		},
		/*post handlers*/
		{
			Path:    "/api/posts",
			handler: app.GetPosts,
			Role:    User,
		},
		{
			Path:    "/api/ws",
			handler: app.WebSocket,
			Role:    User,
		},
		// {
		// 	Path:    "/api/event/create",
		// 	handler: app.CreateEvent,
		// 	Role:    User,
		// },
		// {
		// 	Path:    "/api/event/vote",
		// 	handler: app.VoteEvent,
		// 	Role:    User,
		// },
		// {
		// 	Path:    "/api/event/get",
		// 	handler: app.GetEvent,
		// 	Role:    User,
		// },
		/*... handlers*/
	}
}

func requireLogin(Auth uint) bool {
	return Auth == User
}
