// Package classification Warranty API
//
//Сервис ответственен за работу cо складом.
//
//     Schemes: http
//     Host: lab2-k-t-l-h-warranty
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
	"lab2-microservices-k-t-l-h/internal/pkg/warranty/delivery"
	rep "lab2-microservices-k-t-l-h/internal/pkg/warranty/repo"
	use "lab2-microservices-k-t-l-h/internal/pkg/warranty/usecase"
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
		port = "8180"
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

	wrRepo := rep.NewWrntRepo(*pool)
	wrUse := use.NewWrntUsecase(wrRepo)
	wrHandler := delivery.NewWrntHandler(wrUse)

	r := mux.NewRouter()
	r.Use(utils.InternalServerError)

	r.HandleFunc("/api/v1/warranty/{UUID}",
		wrHandler.Info).Methods("GET")

	r.HandleFunc("/api/v1/warranty/{UUID}/warranty",
		wrHandler.InfoResult).Methods("POST")

	r.HandleFunc("/api/v1/warranty/{UUID}",
		wrHandler.StartWarranty).Methods("POST")

	r.HandleFunc("/api/v1/warranty/{UUID}",
		wrHandler.EndWarranty).Methods("DELETE")

	r.HandleFunc("/manage/health", utils.HealthCheck).Methods("GET")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	sw := utils.NewSwag(dir + "/cmd/warranty/warranty.json")
	r.HandleFunc("/api-docs", sw.Swagger).Methods("GET")

	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)
	log.Print("store running on: ", srv.Addr)
	return srv.ListenAndServe()
}
