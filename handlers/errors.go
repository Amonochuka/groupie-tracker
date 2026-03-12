package handlers

import (
	"html/template"
	"net/http"
)

type ErrorPageData struct {
	Code    int
	Message string
}

func RenderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, message, code)
		return
	}

	tmpl.Execute(w, ErrorPageData{
		Code:    code,
		Message: message,
	})
}