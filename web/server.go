package web

import (
	"log"
	"net/http"

	"github.com/rwxd/new-newt/utils"
)

func WebServer(port string, config utils.Config) {
	log.Println("Serving Web Server on port ", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc(("/check/"), handlerCheck)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Error starting web server: ", err)
	}
}
