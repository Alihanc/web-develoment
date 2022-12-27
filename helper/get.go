package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Alihanc/web-develoment/model"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {

	//Bunker := model.InitialData

	/*t, err := template.ParseFiles("model") //parse the html file homepage.html
	if err != nil {                        // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, Bunker) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {            // if there is an error
		log.Print("template executing error: ", err) //log it
	}*/
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, bunker := range model.InitialData {
		if bunker.Id == key {
			json.NewEncoder(w).Encode(bunker)
		}
	}
}
