package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alihanc/web-develoment/model"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")

}
func All(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.InitialData)
}

func Get(w http.ResponseWriter, r *http.Request) {

	Bunker := model.InitialData
	vars := mux.Vars(r)["id"]
	//vars := r.URL.Query().Get("id")

	// Loop over all of our Bunker
	// if the bunker.Id equals the key we pass in
	// return the bunker encoded as JSON
	for _, bunker := range Bunker {
		if bunker.ID == vars {

			err := json.NewEncoder(w).Encode(&bunker)
			if err != nil {
				log.Fatalln("There was an error encoding the initialized struct")
			}
		}
	}
}
