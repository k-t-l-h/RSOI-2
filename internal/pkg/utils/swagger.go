package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Swag struct {
	filename string
}

func NewSwag(filename string) *Swag {
	return &Swag{filename: filename}
}

func (s *Swag) Swagger(w http.ResponseWriter, r *http.Request) {

	jsonFile, err := os.Open(s.filename)
	if err != nil {
		log.Print(err)
		Response(w, http.StatusInternalServerError, nil)
		return
	}
	defer jsonFile.Close()
	byte, err := ioutil.ReadAll(jsonFile)

	w.Write(byte)
}
