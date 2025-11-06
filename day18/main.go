package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main()  {
	router := mux.NewRouter();
	router.HandleFunc("/",serverHome).Methods("GET");
	log.Fatal(http.ListenAndServe(":3000",router))

}



func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang page</h1>"))
}