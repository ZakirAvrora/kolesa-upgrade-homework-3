package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ZakirAvrora/kolesa-upgrade-homework-3/web/app"
)

func main() {
	port := flag.String("port", "8080", "Network port")

	flag.Parse()

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)

	app := &app.App{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,

		Addr:     ":" + *port,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	app.InfoLog.Printf("Starting server on localhost:%s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
