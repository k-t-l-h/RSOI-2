// Package classification Warehouse API
//
//Сервис ответственен за работу cо складом.
//
//     Schemes: http
//     Host: lab2-k-t-l-h-warehouse
//     BasePath: /
//     Version: 1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Me <kochkarova.lelya@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"lab2-microservices-k-t-l-h/internal/pkg/warehouse/delivery"
	rep "lab2-microservices-k-t-l-h/internal/pkg/warehouse/repo"
	use "lab2-microservices-k-t-l-h/internal/pkg/warehouse/usecase"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	godotenv.Load(".env")
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8280"
	}

	conn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("no database url")
	}

	pool, err := pgxpool.Connect(context.Background(),
		conn)

	if err != nil {
		return errors.New("database connection error")
	}

	warranty, ok := os.LookupEnv("warrantyUrl")

	if !ok {
		return errors.New("no warranty url")
	}

	wrhsRepo := rep.NewWrhsRepo(*pool)
	wrhsUse := use.NewWrhsUsecase(wrhsRepo)
	wrhsHandler := delivery.NewWrhsHandler(warranty, wrhsUse)

	r := mux.NewRouter()
	r.Use(utils.InternalServerError)

	r.HandleFunc("/api/v1/warehouse/{UUID}",
		wrhsHandler.GetItemInfo).Methods("GET")

	r.HandleFunc("/api/v1/warehouse/",
		wrhsHandler.GetItem).Methods("POST")

	r.HandleFunc("/api/v1/warehouse/{UUID}",
		wrhsHandler.ReturnItem).Methods("DELETE")

	r.HandleFunc("/api/v1/warehouse/{UUID}/warranty",
		wrhsHandler.GetItemWarranty).Methods("POST")

	r.HandleFunc("/manage/health", utils.HealthCheck).Methods("GET")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	sw := utils.NewSwag(dir + "/cmd/warehouse/warehouse.json")
	r.HandleFunc("/api-docs", sw.Swagger).Methods("GET")

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)
	log.Print("store running on: ", srv.Addr)
	return srv.ListenAndServe()
}
