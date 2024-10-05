package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-auth/controllers"
)

func RegisterRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
