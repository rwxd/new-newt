package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/rwxd/new-newt/db"
	"github.com/rwxd/new-newt/utils"
)

type Index struct {
	Domains map[string]db.DomainStatus
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(w, "Cannot load config:", err)
		return
	}

	redisClient := db.NewRedisClient(config.RedisHost)
	domains, err := redisClient.GetDomainsToCheck()
	if err != nil {
		log.Fatal(w, err)
		return
	}
	index := Index{Domains: make(map[string]db.DomainStatus)}
	for _, domain := range domains {
		status, err := redisClient.GetDomainStatus(domain)
		if err != nil {
			log.Fatal(w, err)
			continue
		}
		index.Domains[domain] = status
	}

	log.Printf("Found %d domains to check\n", len(domains))

	t, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Println(err)
	}
	log.Println("Rendering template index.html")
	err = t.Execute(w, index)
	if err != nil {
		log.Fatalf("Error executing template: %s", err)
	}
}

func WebServer(port string, config utils.Config) {
	log.Println("Serving Web Server on port ", port)
	http.HandleFunc("/", handlerIndex)
	http.ListenAndServe(":"+port, nil)
}
