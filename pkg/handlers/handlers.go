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

//-------------------------------------------------------------
// (m *Repository) - каждый хэндлер имеет доступ к репозиторию
//-------------------------------------------------------------

// Хендлер для home.page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

	remoteIP := r.RemoteAddr
	// сохраняем ip запроса в сессии
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
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

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}
