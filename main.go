package main

import (
	"log"
	"net/http"

	"github.com/Alihanc/web-develoment/helper"
)

func main() {
	http.HandleFunc("/bunker", helper.Get)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
