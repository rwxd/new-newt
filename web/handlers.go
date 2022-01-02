package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/rwxd/new-newt/db"
)

type Home struct {
	Domains map[string]db.DomainStatus
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	domains, err := db.Rdb.GetDomainsToCheck()
	if err != nil {
		log.Println(w, err)
		return
	}
	log.Printf("Found %d domains to check\n", len(domains))

	home := Home{Domains: make(map[string]db.DomainStatus)}
	for _, domain := range domains {
		status, err := db.Rdb.GetDomainStatus(domain)
		if err != nil {
			log.Println(w, err)
			continue
		}
		home.Domains[domain] = status
	}

	log.Printf("found %d domain status\n", len(home.Domains))

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Rendering templates")
	err = t.Execute(w, home)
	if err != nil {
		log.Printf("Error executing template: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type SingleDomain struct {
	Domain string
	Status db.DomainStatus
}

func handlerCheck(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Path[len("/check/"):]
	log.Printf("Checking domain: %s\n", domain)

	status, err := db.Rdb.RecheckDomainStatus(domain)
	if err != nil {
		log.Println(w, err)
		http.Error(w, "error when checking domain", http.StatusInternalServerError)
	}

	log.Printf("Domain %s is %v\n", domain, status)

	files := []string{
		"./ui/html/domain.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Rendering templates")
	err = t.Execute(w, SingleDomain{Domain: domain, Status: status})
	if err != nil {
		log.Printf("Error executing template: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
