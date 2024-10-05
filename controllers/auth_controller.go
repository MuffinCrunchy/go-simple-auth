package controllers

import (
	"encoding/json"
	"net/http"
	"simple-auth/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("customer_id")
	err := services.Login(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
}
