package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twf1/internal/services"
	"twf1/internal/structs"
)

func CalculateMinCost(writer http.ResponseWriter, request *http.Request) {

	// Decode the JSON request into a generic map to fetch all keys
	var rawInput map[string]interface{}

	err := json.NewDecoder(request.Body).Decode(&rawInput)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	// Validate keys
	order, err := validateKeys(rawInput)
	if err != nil {
		log.Printf("Invalid input: %v", err)
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

func validateKeys(rawInput map[string]interface{}) (structs.ProductDemandQuantity, error) {
	var invalidProduct []string
	var invalidQuantity []string
	order := structs.ProductDemandQuantity{}

	// Iterate over rawInput keys
	for key, value := range rawInput {
		// Check if the key is valid
		if !structs.ValidInputKeys[key] {
			invalidProduct = append(invalidProduct, key)
		} else {
			// Check if the value is of type int
			switch v := value.(type) {
			case float64:
				if v < 0 {
					// If the value is negative
					invalidQuantity = append(invalidQuantity, key)
				} else if v != float64(int(v)) {
					// If the value is not a whole number, add to invalid keys
					invalidQuantity = append(invalidQuantity, key)
				} else {
					intValue := int(v)
					switch key {
					case "A":
						order.A = &intValue
					case "B":
						order.B = &intValue
					case "C":
						order.C = &intValue
					case "D":
						order.D = &intValue // for product D, we have custom error message
					case "E":
						order.E = &intValue // for product E, we have custom error message
					case "F":
						order.F = &intValue
					case "G":
						order.G = &intValue
					case "H":
						order.H = &intValue
					case "I":
						order.I = &intValue
					}
				}
			case int:
				intValue := v
				switch key {
				case "A":
					order.A = &intValue
				case "B":
					order.B = &intValue
				case "C":
					order.C = &intValue
				case "D":
					order.D = &intValue // for product D, we have custom error message
				case "E":
					order.E = &intValue // for product E, we have custom error message
				case "F":
					order.F = &intValue
				case "G":
					order.G = &intValue
				case "H":
					order.H = &intValue
				case "I":
					order.I = &intValue
				}
			default:
				// Skip non-numeric values
				invalidQuantity = append(invalidQuantity, key)
			}
		}
	}

	// If there are invalid keys, return an error with specific details
	if len(invalidProduct) > 0 || len(invalidQuantity) > 0 {
		invalidProductList := ""
		invalidQuantityList := ""

		if len(invalidProduct) > 0 {
			invalidProductList = fmt.Sprintf("invalid products: %v \n", invalidProduct)
		}

		if len(invalidQuantity) > 0 {
			invalidQuantityList = fmt.Sprintf("invalid order quantities for products %v", invalidQuantity)
		}

		return order, fmt.Errorf("%s%s", invalidProductList, invalidQuantityList)
	}

	return order, nil

}
