package handlers

import (
	"encoding/json"
	"fmt"
	"goDev/booking-application/internal/config"
	"goDev/booking-application/internal/models"
	"goDev/booking-application/internal/render"
	"log"
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})

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

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//получение данных из формы на странице
	// аргумент Get - это name формы в html
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// структура json ответа для проверки доступности комнаты
type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

// обрабатывает запрос о доступности бронирования и отправляет ответ в виде JSON
func (m *Repository) AvailabilitySJON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK: false,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	// создаем хедер, что будет ответ в формате json
	w.Header().Set("Content-Type", "application/json")
	// отправляем ответ
	w.Write(out)
	
}
