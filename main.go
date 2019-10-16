package main

import (
	"fmt"
	"log"
	"net/http"
	cmd "projectapi/cmd"

	"github.com/gorilla/mux"
)

func main() {
	//mux router
	router := mux.NewRouter()
	router.HandleFunc("/api/garages", cmd.GaragesHandler).Methods("GET")
	router.HandleFunc("/api/getcarposition/{idcar}", cmd.GetCarPositionHandler).Methods("GET")
	router.HandleFunc("/api/getgaragedetail/{idgarage}", cmd.GetGarageDetail).Methods("GET")
	fmt.Println("Server started in http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
