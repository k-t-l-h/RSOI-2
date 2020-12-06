package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Response(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		jsn, _ := json.Marshal(body)
		w.Write(jsn)
	}
}

func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	if resp.StatusCode == http.StatusInternalServerError {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		w.WriteHeader(resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err == nil || len(bodyBytes) != 0 {
		w.Write(bodyBytes)
	}
}
