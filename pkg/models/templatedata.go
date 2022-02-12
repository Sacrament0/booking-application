package models

// структура, описывающая данные, посылаемые на страницу
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	// безопасность
	CSRFToken string
	//короткое сообщение для пользователя
	Flash   string
	Warning string
	Error   string
}