package routes

import (
	"github.com/jacktrusler/movie-delphia/go-backend/pkg/app"
)

func RegisterRoutes(a *app.App) {
	a.Echo.GET("/users", GetUsers(a))
	a.Echo.POST("/users", PostUser(a))
	a.Echo.PUT("/users/:id", PutUser(a))
	a.Echo.DELETE("/users/:id", DeleteUser(a))
}
