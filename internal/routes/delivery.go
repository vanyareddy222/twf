package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"twf1/internal/services"
)

type OrderRequest struct {
	Products map[string]int `json:"products"`
}

func CalculateDeliveryCost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var order OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	cost := services.CalculateMinimumCost(order.Products)
	response := map[string]interface{}{
		"minimum_cost": cost,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
