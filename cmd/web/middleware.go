package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// добавляет CSRF защиту для POST запросов
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// загружает и сохраняет сессии от каждого запроса
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
