package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"twf1/internal/services"
	"twf1/internal/structs"
)

func CalculateMinCost(writer http.ResponseWriter, request *http.Request) {

	// Process the JSON input and sum duplicate keys
	order, err := processAndSumDuplicateKeys(request.Body)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	//// Validate keys
	//order, err := validateKeys(summedInput)
	//if err != nil {
	//	log.Printf("Invalid input: %v", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	writer.Write([]byte(err.Error()))
	//	return
	//}

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

func processAndSumDuplicateKeys1(body io.ReadCloser) (map[string]interface{}, error) {
	defer body.Close()

	// Read the raw JSON input
	rawBody, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	// Use a custom decoder to parse the JSON
	decoder := json.NewDecoder(bytes.NewReader(rawBody))
	result := make(map[string]float64)

	// Expect the JSON object to start
	token, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	if delim, ok := token.(json.Delim); !ok || delim != '{' {
		return nil, fmt.Errorf("expected JSON object to start with '{'")
	}

	// Process key-value pairs
	var currentKey string
	for decoder.More() {
		// Read the key
		token, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("failed to read JSON token: %w", err)
		}

		// Keys should always be strings
		if key, ok := token.(string); ok {
			currentKey = key
		} else {
			return nil, fmt.Errorf("expected string key, got %T", token)
		}

		// Read the value
		token, err = decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("failed to read value for key %s: %w", currentKey, err)
		}

		// Only handle numeric values
		if value, ok := token.(float64); ok {
			// Sum values for duplicate keys
			result[currentKey] += value
		} else {
			return nil, fmt.Errorf("expected numeric value for key %s, got %T", currentKey, token)
		}
	}

	// Finalize the JSON object
	token, err = decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to parse closing JSON object: %w", err)
	}
	if delim, ok := token.(json.Delim); !ok || delim != '}' {
		return nil, fmt.Errorf("expected JSON object to end with '}'")
	}

	// Convert the result map to map[string]interface{}
	finalResult := make(map[string]interface{})
	for key, value := range result {
		finalResult[key] = value
	}

	return finalResult, nil
}

func processAndSumDuplicateKeys(body io.ReadCloser) (structs.ProductDemandQuantity, error) {
	defer body.Close()

	// Read the raw JSON input
	rawBody, err := io.ReadAll(body)
	if err != nil {
		return structs.ProductDemandQuantity{}, fmt.Errorf("failed to read request body: %w", err)
	}

	// Use a custom decoder to parse the JSON
	decoder := json.NewDecoder(bytes.NewReader(rawBody))
	result := make(map[string]float64)

	// Expect the JSON object to start
	token, err := decoder.Token()
	if err != nil {
		return structs.ProductDemandQuantity{}, fmt.Errorf("failed to parse JSON: %w", err)
	}
	if delim, ok := token.(json.Delim); !ok || delim != '{' {
		return structs.ProductDemandQuantity{}, fmt.Errorf("expected JSON object to start with '{'")
	}

	// Process key-value pairs
	var currentKey string
	var invalidProduct []string
	var invalidQuantity []string

	for decoder.More() {
		// Read the key
		token, err := decoder.Token()
		if err != nil {
			return structs.ProductDemandQuantity{}, fmt.Errorf("failed to read JSON token: %w", err)
		}

		// Keys should always be strings
		if key, ok := token.(string); ok {
			currentKey = key
		} else {
			invalidProduct = append(invalidProduct, fmt.Sprintf("%v", token))
			continue
		}

		// Read the value
		token, err = decoder.Token()
		if err != nil {
			return structs.ProductDemandQuantity{}, fmt.Errorf("failed to read value for key %s: %w", currentKey, err)
		}

		// Only handle numeric values
		if value, ok := token.(float64); ok {
			// Validate the key and value
			if !structs.ValidInputKeys[currentKey] {
				invalidProduct = append(invalidProduct, currentKey)
				continue
			}
			if value < 0 || value != float64(int(value)) {
				invalidQuantity = append(invalidQuantity, currentKey)
			} else {
				// Sum values for duplicate keys
				result[currentKey] += value
			}
		} else {
			invalidQuantity = append(invalidQuantity, currentKey)
		}
	}

	// Finalize the JSON object
	token, err = decoder.Token()
	if err != nil {
		return structs.ProductDemandQuantity{}, fmt.Errorf("failed to parse closing JSON object: %w", err)
	}
	if delim, ok := token.(json.Delim); !ok || delim != '}' {
		return structs.ProductDemandQuantity{}, fmt.Errorf("expected JSON object to end with '}'")
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

		return structs.ProductDemandQuantity{}, fmt.Errorf("%s%s", invalidProductList, invalidQuantityList)
	}

	// Convert the result map to structs.ProductDemandQuantity
	order := structs.ProductDemandQuantity{}
	for key, value := range result {
		intValue := int(value)
		switch key {
		case "A":
			order.A = &intValue
		case "B":
			order.B = &intValue
		case "C":
			order.C = &intValue
		case "D":
			order.D = &intValue
		case "E":
			order.E = &intValue
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

	return order, nil
}
