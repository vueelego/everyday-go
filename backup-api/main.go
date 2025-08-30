package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	port   int
	dsn    string
	logger *slog.Logger
}

func main() {
	var app application

	app.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(app.logger)

	flag.IntVar(&app.port, "port", 6000, "服务端口")
	flag.StringVar(&app.dsn, "dsn", "./db.sqlite", "sqlite3 数据库")
	flag.Parse()

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", app.port),
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelWarn),
		IdleTimeout:  time.Minute,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("server stop: " + err.Error())
	}
}

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.recoverPanic)
	mux.Use(app.enableCORS)

	mux.Post("/save", app.saveHandler)
	mux.Get("/get", app.getHandler)

	return mux
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
