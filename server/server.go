package server

import (
	// "fmt"
	"log"
	"net/http"
	helpers "podcasts/helpers"
	"podcasts/template"
)

// todo: disabled for now
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
// }

func StartServer(useDummyData bool) {

	// fmt.Println(useDummyData)

	// if !useDummyData {
	template.BuildTemplate()
	// }

	fs := http.FileServer((http.Dir("./tmp")))

	// http.HandleFunc("/", fs)
	http.Handle("/", fs)
	log.Print("Listening on :8080...")

	err := http.ListenAndServe("localhost:8080", nil)
	helpers.HandleError(err, "starting the server")

}
