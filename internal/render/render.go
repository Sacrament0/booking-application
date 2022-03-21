package render

import (
	"bytes"
	"fmt"
	"goDev/booking-application/internal/config"
	"goDev/booking-application/internal/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
)

// переменная, содержащая функции для использования в  шаблоне (в renderTemplateTest)
var functions = template.FuncMap{}

// переменная для хранения конфига, полученного из main
var app *config.AppConfig

// когда эта функция выполняется в main, в переменную app записывается конфиг
func NewTemplates(a *config.AppConfig) {
	app = a
}

// функция добавления дефолтных данных на страницу, нужна чтобы добавлять в нее, а не в RenderTemplate
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// renderTemplate рендерит страницу (второй аргумент - название страницы), третий - данные для страницы
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	// переменная для хранени кэша
	var tc map[string]*template.Template

	// если конфиг true, использовать сгенереный кэш
	if app.UseCache {
		// вытаскиваем из записанного конфига кэш
		tc = app.TemplateCache
	} else {
		// если false, постоянно генерить страницы при загрузке
		tc, _ = CreateTemplateCache()
	}
	// выбираем шаблон в соответствии с нужным именем
	t, ok := tc[tmpl]
	// если такого шаблона нет, то падаем
	if !ok {
		log.Fatal("Could not get template from template Cache")
	}
	// создаем буфер для хранения шаблона
	buf := new(bytes.Buffer)

	//добавляем данные на страницу
	td = AddDefaultData(td, r)

	// записываем шаблон в буфер, второй аргумент - переданные данные
	// td - CSRF токен, чтобы можно было постить
	_ = t.Execute(buf, td)
	// отправляем браузеру все, что хранится в буфере
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache находит страницы, преобразовывает в шаблоны и склеивает с layout
func CreateTemplateCache() (map[string]*template.Template, error) {

	// мапа, где хранятся все склеенные c layout шаблоны
	myCache := map[string]*template.Template{}

	// получить все названия страниц (т.е. без layout base) (...page.html)
	// Glob получает пути ко всем файлам в соответствии с заданным паттерном
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// отрезает путь к файлу и сохраняет только название страницы
		name := filepath.Base(page)
		// new  создает новый шаблон с названием name из файла лежащего по пути page. Func поставляет функции в шаблон
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// проверяем, нужно ли использовать какие-нибудь layouts вместе с templates
		// сначала находим layout
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		// если что-то нашлось
		if len(matches) > 0 {
			// склеиваем шаблон и layout
			ts, err := ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}
