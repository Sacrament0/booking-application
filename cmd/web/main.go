package main

import (
	"goDev/booking-application/pkg/config"
	"goDev/booking-application/pkg/handlers"
	"goDev/booking-application/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// переменная для хранения КОНФИГА
var app config.AppConfig
// перемення для хранения СЕССИИ
var session *scs.SessionManager

func main() {

	// поменять на true если запуск в продакшене
	app.InProduction = false

	//создание сессии
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	//помещаем сессию в конфиг
	app.Session = session

	// создаем кеш всех страниц
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot careate template cache")
	}
	
	// помещаем кэш в переменную для конфига
	app.TemplateCache = tc
	// задаем указание к использованию кэша
	app.UseCache = false

	// помещаем переменную для конфига в репозиторий, т.е. создаем репозиторий
	repo := handlers.NewRepo(&app)

	// эта функция записывает конфиг в виде репозитория в переменную в пакете handlers, поэтому он становится доступным там
	handlers.NewHandlers(repo)

	// эта функция записывает конфиг в переменную в пакете render, поэтому он становится доступен там
	// поэтому при рендеринге страницы он не будет загружаться снова и снова
	render.NewTemplates(&app)
	

	// создаем и настраиваем сервер
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// запускаем сервер
	err = srv.ListenAndServe()
	log.Fatal(err)
}
