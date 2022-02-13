package main

import (
	"goDev/booking-application/pkg/config"
	"goDev/booking-application/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// создаем мультиплексор
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// хэндлеры
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// обработка статических файлов
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
