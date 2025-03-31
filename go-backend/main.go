package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jacktrusler/movie-delphia/go-backend/pkg/app"
	"github.com/jacktrusler/movie-delphia/go-backend/pkg/database"
	"github.com/jacktrusler/movie-delphia/go-backend/pkg/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		slog.Error("failed to get CWD", "error", err)
	}
	fullPath := filepath.Join(cwd, "pkg", "database", "main.db")

	db, err := database.InitDB(fullPath)
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	application := app.NewApp(db)

	routes.RegisterRoutes(application)

	defer db.Close()

	if err := application.Echo.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		application.L.Error("failed to start server", "error", err)
	}
}
