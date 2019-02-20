package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main()  {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/me", Me)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Index")
}

func Me(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Me")
}