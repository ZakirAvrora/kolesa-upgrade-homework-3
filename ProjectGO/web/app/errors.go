package app

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"text/template"
)

func (app *App) Errors(w http.ResponseWriter, status int, err error) {
	if status == http.StatusInternalServerError {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.ErrorLog.Output(2, trace)
	}

	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("web/ui/templates/error.html")
	if err != nil {
		http.Error(w, strconv.Itoa(status)+" "+http.StatusText(status), status)
		return
	}
	statusint := strconv.Itoa(status) + " " + http.StatusText(status)
	tmpl.ExecuteTemplate(w, "error.html", statusint)
}
