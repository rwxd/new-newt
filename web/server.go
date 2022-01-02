package web

import (
	"log"
	"net/http"

	"github.com/rwxd/new-newt/utils"
)

func WebServer(port string, config utils.Config) {
	log.Println("Serving Web Server on port ", port)
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/check/", handlerCheck)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting web server: ", err)
	}
}
