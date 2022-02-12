package handlers

import (
	"goDev/booking-application/pkg/config"
	"goDev/booking-application/pkg/models"
	"goDev/booking-application/pkg/render"
	"net/http"
)

// ПРИКРУТИТЬ КОНФИГ К ХЭНДЛЕРАМ -----------------------------------------------------------
// переменная для репозитория
var Repo *Repository

// структура описывающая репозиторий (по сути оборачивает структуру конфига в другую структуру)
type Repository struct {
	App *config.AppConfig
}

// записывает структуру конфига в переменную, т.о. создавая новый репозиторий
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// записывает созданный в main репозиторий в переменную в пакете handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// КОНЕЦ ПРИКРУТИТЬ КОНФИГ К ХЭНДЛЕРАМ -----------------------------------------------------------

//-------------------------------------------------------------
// (m *Repository) - каждый хэндлер имеет доступ к репозиторию
//-------------------------------------------------------------

// Хендлер для home.page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

	remoteIP := r.RemoteAddr
	// сохраняем ip запроса в сессии
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP )
}

// Хендлер для about.page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
