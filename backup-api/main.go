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
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	port   int
	dsn    string
	store  Store
	logger *slog.Logger
}

// 运行，如：go run . -dsn=./sqlite3/xxx.sqlite
func main() {
	var app application

	app.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(app.logger)

	flag.IntVar(&app.port, "port", 6001, "服务端口")
	flag.StringVar(&app.dsn, "dsn", "./sqlite3/db.sqlite", "sqlite3 数据库")
	flag.Parse()

	app.store = NewStore(app.dsn)

	app.store.CreateTable()

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", app.port),
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelWarn),
		IdleTimeout:  time.Minute,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("start server on " + srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("server stop: " + err.Error())
	}
}

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.recoverPanic)
	mux.Use(middleware.Logger)
	mux.Use(app.enableCORS)

	mux.Post("/get", app.getHandler)   // 获取
	mux.Post("/save", app.saveHandler) // 添加

	router := chi.NewRouter()
	router.Mount("/v1", mux)

	return router
}
