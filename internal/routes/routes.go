package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"twf1/internal/controllers"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/calculateMinCost", controllers.CalculateMinCost).Methods(http.MethodPost) //22 DONE

	return router
}
