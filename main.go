package main

import (
	"awesomeProject/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/home", home).Methods(http.MethodGet)
	router.HandleFunc("/api/{endpoint}", controllers.GetRandom).Methods(http.MethodGet)
	router.HandleFunc("/api/internal/", controllers.GetInternal).Methods(http.MethodGet)
	router.HandleFunc("/api/internal/calculateResponseTime/{endpoint}", controllers.GetResponseTime).Methods(http.MethodGet)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Home sweet home")
	log.Printf("/home Request from: %s", r.RemoteAddr)
}
