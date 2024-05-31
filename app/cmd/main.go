package main

import (
	"crime-stats/app/control"
	"crime-stats/app/store"
	"log"
	"net/http"
)

func main() {
	store := store.NewInMemoryCrimeService()
	crimeServer := &control.CrimeServer{
		CrimeService: store,
	}
	log.Fatal(http.ListenAndServe(":8080", crimeServer))
}
