package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Alihanc/web-develoment/model"
)

func NewBunker(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var bunker model.Bunker
	json.Unmarshal(reqBody, &bunker)
	// update our global model.IniteilData array to include
	// our new Bunker
	model.InitialData = append(model.InitialData, bunker)

	json.NewEncoder(w).Encode(bunker)
}
