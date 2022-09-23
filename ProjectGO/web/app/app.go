package app

import (
	"log"
	"net/http"
)

type App struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./web/ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/cats", app.catView)
	mux.HandleFunc("/details", app.catDetails)

	return mux
}
