package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type VdovinDTO struct {
	FirstName    string                 `json:"first_name"`
	LastName     string                 `json:"last_name"`
}

func handleRequests() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/vdovin", VdovinHandler)

	http.Handle("/", router)

	return cors.AllowAll().Handler(router)
}

func VdovinHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(VdovinDTO{
		FirstName:    "Ivan",
		LastName:     "Vdovin",
	})
}

func main() {
	fmt.Println("жопа порт :7171")
	err := http.ListenAndServe(":7171", handleRequests())
	if err != nil {
		log.Fatal("Internal error!")
	}
}
