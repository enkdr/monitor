package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/enkdr/monitor/config"
	"github.com/enkdr/monitor/database"
	"github.com/jmoiron/sqlx"
)

type App struct {
	templates *template.Template
	port      int
	db        *sqlx.DB
}

func NewApp() *http.Server {

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	currentDir, _ := os.Getwd()
	templatesPath := currentDir + config.TEMPLATE_PATH

	NewApp := &App{
		templates: template.Must(template.ParseFiles(templatesPath)),
		port:      config.PORT,
		db:        db,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewApp.port),
		Handler: NewApp.RegisterRoutes(),
	}

	return server

}
