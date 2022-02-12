package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// структура описывающая конфига приложения
type AppConfig struct {
	// использовать кэш или нет
	UseCache bool
	// кэш страниц
	TemplateCache map[string]*template.Template
	// запуск сервера в продакшене или нет
	InProduction bool
	// Сессии
	Session *scs.SessionManager
}
