package main

import (
	"goDev/booking-application/internal/config"
	"goDev/booking-application/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// создаем мультиплексор
	mux := chi.NewRouter()

	// MIDDLEWARE

	mux.Use(middleware.Recoverer)
	// ignore any request that is a post that doesnt have proper cross site request forgery token
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// хэндлеры
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilitySJON)

	mux.Get("/contact", handlers.Repo.Contact)

	// обработка статических файлов
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
