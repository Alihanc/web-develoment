package main

import (
	helper "ht/helper"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bunker", helper.Get)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
