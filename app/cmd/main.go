package main

import (
	"crime-stats/app/control"
	"crime-stats/app/store"
	"crime-stats/app/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	err := utils.LoadCrimeData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting crime-stats service")

	store := store.NewInMemoryCrimeService()
	crimeServer := &control.CrimeServer{
		CrimeService: store,
	}
	log.Fatal(http.ListenAndServe(":8080", crimeServer))
}
