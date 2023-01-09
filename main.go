package main

import (
	"log"
	"net/http"

	"github.com/Alihanc/web-develoment/helper"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helper.HomePage)
	myRouter.HandleFunc("/bunkers", helper.All)
	myRouter.HandleFunc("/bunker", helper.NewBunker).Methods("POST")
	myRouter.HandleFunc("/delete/{id}", helper.DeleteBunker).Methods("DELETE")
	myRouter.HandleFunc("/bunker/update", helper.Update).Methods("PUT")
	myRouter.HandleFunc("/bunker/{id}", helper.Get)
	myRouter.HandleFunc("/bunker/3600", helper.Maintance_3600)
	myRouter.HandleFunc("/bunker/7200", helper.Maintance_7200)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
