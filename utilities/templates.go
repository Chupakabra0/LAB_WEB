package utilities

import (
	"html/template"
	"net/http"
)

// Инстанс парсера файлов
var templateParser *template.Template

func ParseTemplates(pattern string) {
	// Парсим все html файлы в текущей директории
	templateParser = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(writter http.ResponseWriter, fileName string, data interface{}) error {
	return templateParser.ExecuteTemplate(writter, fileName, data)
}