package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/nico-mayer/go-elements/models"
)

func ElementByAtomicNumber(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(req.URL.Path, "/")
	atomicNumberString := pathParts[len(pathParts)-1]
	atomicNumber, err := strconv.Atoi(atomicNumberString)

	if err != nil {
		http.Error(res, "Invalid atomic number", http.StatusBadRequest)
		return
	}

	element := models.ElementByAtomicNumber(atomicNumber)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(element)

}
