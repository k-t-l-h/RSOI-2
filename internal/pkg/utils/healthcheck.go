package utils

import (
	"lab2-microservices-k-t-l-h/internal/models"
	"net/http"
	"os"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	name, password, state := r.BasicAuth()
	if !state {
		Response(w, http.StatusForbidden, nil)
		return
	}

	adminName, _ := os.LookupEnv("ADMIN")
	adminPsw, _ := os.LookupEnv("ADMIN_PASSWORD")

	if name != adminName || password != adminPsw {
		Response(w, http.StatusForbidden, nil)
		return
	}

	answer := models.HealthCheck{
		Status: "UP",
	}
	Response(w, http.StatusOK, answer)

}
