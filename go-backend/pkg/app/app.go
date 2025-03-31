package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/jacktrusler/movie-delphia/go-backend/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Db   *sql.DB
	Echo *echo.Echo
	L    *slog.Logger
}

func NewApp(dbInstance *sql.DB) *App {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${uri} ${status} ${latency_human}\n",
		Output: os.Stdout,
	}))

	// Cors for localhost:3000 and default vite port: 5173
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	l := utils.GetLogger()

	return &App{
		Db:   dbInstance,
		Echo: e,
		L:    l,
	}
}
