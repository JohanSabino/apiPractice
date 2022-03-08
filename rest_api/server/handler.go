package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "Method not allowed")
		if err != nil {
			return
		}
		return
	}
	_, err := fmt.Fprintf(w, "First time using http library %s", "visitor")
	if err != nil {
		return
	}
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(countries)
	if err != nil {
		return
	}
}

func addCountries(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "%v", err)
		if err != nil {
			return
		}
		return
	}

	countries = append(countries, country)
	_, err = fmt.Fprintf(w, "Country was added")
	if err != nil {
		return
	}
}
