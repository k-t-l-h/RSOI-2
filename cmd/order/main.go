// Package classification Order API
//
//Сервис ответственен за работу с заказом, получение товара со склада (запрос к Warehouse) и создание гарантии (запрос к Warranty).
//
//
//     Schemes: http
//     Host: lab2-k-t-l-h-order
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
	"lab2-microservices-k-t-l-h/internal/pkg/orders/delivery"
	rep "lab2-microservices-k-t-l-h/internal/pkg/orders/repo"
	use "lab2-microservices-k-t-l-h/internal/pkg/orders/usecase"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
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
		port = "8380"
	}

	conn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("no database url")
	}

	pool, err := pgxpool.Connect(context.Background(),
		conn)

	if err != nil {
		return err
	}

	warehouse, ok := os.LookupEnv("warehouseUrl")

	if !ok {
		return errors.New("no warehouse url")
	}

	warranty, ok := os.LookupEnv("warrantyUrl")

	if !ok {
		return errors.New("no warranty url")
	}

	orderRepo := rep.NewOdrRepo(*pool)
	orderUse := use.NewOdrUsecase(orderRepo)
	orderHandler := delivery.NewOdrHandler(orderUse, warehouse, warranty)

	r := mux.NewRouter()
	r.Use(utils.InternalServerError)

	r.HandleFunc("/api/v1/orders/{UUID}",
		orderHandler.MakeOrders).Methods("POST")

	r.HandleFunc("/api/v1/orders/{UUID}/{orderUUID}",
		orderHandler.OrderInfo).Methods("GET")

	r.HandleFunc("/api/v1/orders/{UUID}",
		orderHandler.UsersOrders).Methods("GET")

	r.HandleFunc("/api/v1/orders/{orderUUID}/warranty",
		orderHandler.GetWarranty).Methods("POST")

	r.HandleFunc("/api/v1/orders/{UUID}",
		orderHandler.ReturnOrder).Methods("DELETE")

	r.HandleFunc("/manage/health", utils.HealthCheck).Methods("GET")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	sw := utils.NewSwag(dir + "/cmd/order/order.json")
	r.HandleFunc("/api-docs", sw.Swagger).Methods("GET")


	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)
	log.Print("store running on: ", srv.Addr)
	return srv.ListenAndServe()
}
