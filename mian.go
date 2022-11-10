package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

type Movie struct {
	Name string
	Year int
}


func Index(w http.ResponseWritter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Use /api/v1/<name> \n")
}


func MovieApi(w http.ResponseWritter, r *http.Request, ps httprouter.Params){
	m := Movie {
		Name : ps.ByName("name"),
		Year: 1982
	}
	json.NewEncoder(w).Encode(m)
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/api/v1/name", MovieApi)

	log.Fatal(http.ListenAndServe(":8080" , router))
}