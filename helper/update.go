package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alihanc/web-develoment/model"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bunker model.Bunker
	err := json.NewDecoder(r.Body).Decode(&bunker)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	for index, bunkers := range model.InitialData {
		if bunkers.NAME == bunker.NAME {
			model.InitialData = append(model.InitialData[:index], model.InitialData[index+1:]...)
		}
	}
	model.InitialData = append(model.InitialData, bunker)
	err = json.NewEncoder(w).Encode(&bunker)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
