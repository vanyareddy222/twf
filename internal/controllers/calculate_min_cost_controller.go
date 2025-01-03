package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"twf1/internal/services"
	"twf1/internal/structs"
)

func CalculateMinCost(writer http.ResponseWriter, request *http.Request) {

	order := structs.ProductDemandQuantity{}
	err := json.NewDecoder(request.Body).Decode(&order)

	if err != nil {
		log.Printf("error decoding request body: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	warehouseOrder := services.GetWarehouseDemandWeight(order)
	minCost := services.GetMinCostService(warehouseOrder)

	data, _ := json.Marshal(minCost)
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
	return
}
