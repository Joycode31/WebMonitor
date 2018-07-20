
package main

import (
	"net/http"
	"fmt"
	"../pkg/mux"
)

// The new router function creates the router and
// returns it to us.
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./views/")
	staticFileHandler := http.StripPrefix("/views", http.FileServer(staticFileDirectory))
	r.PathPrefix("/views").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/getsites", getSiteHandler).Methods("GET")
	r.HandleFunc("/getsites", createSiteHandler).Methods("POST")
	r.HandleFunc("/deleteurl", DeleteUrlHandler).Methods("POST")
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go !!!")
}

var sitesMap map[string]SiteMonitor
func main() {
	r := newRouter()
	sitesMap = make(map[string]SiteMonitor)
	go CheckUrlStatus()
	http.ListenAndServe(":8080", r)
}

