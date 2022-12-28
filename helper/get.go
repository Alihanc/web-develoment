package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Alihanc/web-develoment/model"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)["id"]
	//fmt.Fprintf(w, "key"+vars)

	// Loop over all of our Bunker
	// if the bunker.Id equals the key we pass in
	// return the bunker encoded as JSON
	for _, bunker := range model.InitialData {
		if bunker.ID == vars {

			json.NewEncoder(w).Encode(bunker)
		}
	}
}
