package utils

import (
	"encoding/json"
	"lab2-microservices-k-t-l-h/internal/models"
	"net/http"
)

func InternalServerError(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				answer := models.ErrorMessage{
					Message: "Sorry :(",
				}
				jsn, _ := json.Marshal(answer)
				w.Write(jsn)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
