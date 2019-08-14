package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	dbs "projectapi/db"
	"projectapi/delivery"
	postgres "projectapi/repository/postgres/implementasi"
	restapi "projectapi/repository/restapi/implementasi"
	"projectapi/usecase"
)

func main() {
	conn := dbs.Conn()
	httpDelivery := initAll(conn)
	router := httpDelivery.CreateRouter()
	fmt.Println("Server started in http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initAll(db *sql.DB) (httpDelivery delivery.Delivery) {
	carData, err := restapi.New()
	if err != nil {
		fmt.Println("handle Error")
	}
	postgresRepo, errDb := postgres.New(db)
	if errDb != nil {
		fmt.Println("handle Error")
	}
	postgresRepo.PrepareQuery()
	Service := usecase.New(postgresRepo, carData)
	httpDelivery = delivery.New(Service)
	return
}
