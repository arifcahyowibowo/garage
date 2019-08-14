package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCarPositionHandler handler for reques garages list
func (D Delivery) GetCarPositionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idcar := params["idcar"]
	carPositionDatas := D.Service.GetCarPosition(idcar)
	jsonCarPositionDatas, err := json.Marshal(carPositionDatas)
	if err != nil {
		log.Fatal("failed to fetch json")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonCarPositionDatas))
}

// GetGarageDetailHandler handler for reques garages detail
func (D Delivery) GetGarageDetailHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDGarage, _ := strconv.Atoi(params["idgarage"])
	retData := D.Service.GetGarageDetail(IDGarage)
	jsonRetData, _ := json.Marshal(retData)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonRetData))
}

// CreateRouter ...
func (D Delivery) CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/getcarposition/{idcar}", D.GetCarPositionHandler).Methods("GET")
	router.HandleFunc("/api/getgaragedetail/{idgarage}", D.GetGarageDetailHandler).Methods("GET")
	return router
}
