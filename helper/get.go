package helper

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Bunke := &model.InitialData
	vars := mux.Vars(r)["id"]
	fmt.Fprintf(w, "key"+vars)

	// Loop over all of our Bunker
	// if the bunker.Id equals the key we pass in
	// return the bunker encoded as JSON
	/*for _, bunker := range *Bunke {
		if bunker.Id == vars {

			json.NewEncoder(w).Encode(bunker)
		}
	}*/
}
