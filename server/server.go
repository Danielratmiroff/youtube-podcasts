package server

import (
	"io/ioutil"
	"log"
	"net/http"
	helpers "podcasts/helpers"
	"podcasts/template"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got a %s request for: %v", r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	log.Println("Status response:", http.StatusOK)

	search := r.URL.RawQuery
	template.BuildTemplate(search)
	fileBytes, err := ioutil.ReadFile("./build/index.html")
	helpers.HandleError(err, "reading search")
	w.Write(fileBytes)
}

func StartServer(useDummyData bool) {

	template.BuildTemplate("")

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer((http.Dir("./build"))))
	mux.HandleFunc("/search", searchHandler)

	log.Print("Listening on :8080...")
	err := http.ListenAndServe("localhost:8080", mux)
	helpers.HandleError(err, "starting the server")

}
