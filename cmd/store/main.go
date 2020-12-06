// Package classification Store API
//
// Сервис является gateway
// Именно здесь происходят проверка пользователя
//
//
//     Schemes: http
//     Host: lab2-k-t-l-h-store
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
	_ "github.com/swaggo/http-swagger"
	"lab2-microservices-k-t-l-h/internal/pkg/store/delivery"
	rep "lab2-microservices-k-t-l-h/internal/pkg/store/repo"
	use "lab2-microservices-k-t-l-h/internal/pkg/store/usecase"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//go:generate swagger generate spec
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
		port = "8480"
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

	orders, ok := os.LookupEnv("orderUrl")

	if !ok {
		return errors.New("no order url")
	}

	warehouse, ok := os.LookupEnv("warehouseUrl")

	if !ok {
		return errors.New("no warehouse url")
	}

	warranty, ok := os.LookupEnv("warrantyUrl")

	if !ok {
		return errors.New("no warranty url")
	}

	strRepo := rep.NewStrRepo(*pool)
	strUse := use.NewStrUsecase(strRepo)
	strHandler := delivery.NewStrHandler(orders, warehouse, warranty, strUse)

	r := mux.NewRouter()
	r.Use(utils.InternalServerError)

	//check
	r.HandleFunc("/api/v1/store/{UUID}/orders",
		strHandler.Orders).Methods("GET")

	r.HandleFunc("/api/v1/store/{UUID}/{ORDER_UUID}",
		strHandler.OrdersInfo).Methods("GET")

	r.HandleFunc("/api/v1/store/{UUID}/{ORDER_UUID}/warranty",
		strHandler.Warranty).Methods("POST")

	r.HandleFunc("/api/v1/store/{UUID}/purchase",
		strHandler.Purchase).Methods("POST")

	r.HandleFunc("/api/v1/store/{UUID}/{ORDER_UUID}/refund",
		strHandler.Refund).Methods("DELETE")

	r.HandleFunc("/manage/health", utils.HealthCheck).Methods("GET")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	sw := utils.NewSwag(dir + "/cmd/store/store.json")
	r.HandleFunc("/api-docs", sw.Swagger).Methods("GET")

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)
	log.Print("store running on: ", srv.Addr)
	return srv.ListenAndServe()
}
