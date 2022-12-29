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
	myRouter.HandleFunc("/all", helper.All)
	myRouter.HandleFunc("/bunker/{?:id}", helper.Get)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
