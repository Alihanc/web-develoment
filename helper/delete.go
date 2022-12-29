package helper

import (
	"net/http"

	"github.com/Alihanc/web-develoment/model"
	"github.com/gorilla/mux"
)

func DeleteBunker(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the bunker we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our model.InitialData
	for index, bunker := range model.InitialData {
		// if our id path parameter matches one of our
		// model.InitialData
		if bunker.ID == id {
			// updates our model.InitialData array to remove the
			// bunker
			model.InitialData = append(model.InitialData[:index], model.InitialData[index+1:]...)
		}
	}

}
