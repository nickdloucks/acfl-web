package main

import (
	"log"
	"net/http"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mainHandler := http.HandlerFunc(IndexPageHandler)
	log.Fatal(http.ListenAndServeTLS(
			":5000", 
			"/run/secrets/acfl-cert.pem", 
			"/run/secrets/acfl-key.pem", 
			mainHandler,
		),
	)
}