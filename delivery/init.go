package delivery

import (
	"projectapi/usecase"
)

// Delivery is struct to handle delivery http
type Delivery struct {
	Service *usecase.GarageServiceStruct
}

// New is Consturctor create service
func New(Service *usecase.GarageServiceStruct) (httpDelivery Delivery) {
	httpDelivery.Service = Service
	return httpDelivery
}
