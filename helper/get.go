package helper

import (
	"log"
	"net/http"
	"text/template"

	model "ht/model"

	"http"
)

func Get(w http.ResponseWriter, r *http.Request) {

	Bunker := model.InitialData

	t, err := template.ParseFiles("bunker.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, Bunker) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {            // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
