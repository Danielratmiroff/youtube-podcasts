package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Print("Listening server...")
	err := http.ListenAndServe(":"+port, mux)
	helpers.HandleError(err, "starting the server")

}
