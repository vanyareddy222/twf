package services

import (
	"twf1/internal/structs"
)

// GetWarehouseDemandWeight calculates the weighted demand for each warehouse center.
func GetWarehouseDemandWeight(input structs.ProductDemandQuantity) structs.WarehouseCenterDemandQuantity {
	var c1, c2, c3 float64

	if input.A != nil {
		c1 += float64(*input.A) * structs.ProductWeights["A"]
	}
	if input.B != nil {
		c1 += float64(*input.B) * structs.ProductWeights["B"]
	}
	if input.C != nil {
		c1 += float64(*input.C) * structs.ProductWeights["C"]
	}
	if input.D != nil {
		c2 += float64(*input.D) * structs.ProductWeights["D"]
	}
	if input.E != nil {
		c2 += float64(*input.E) * structs.ProductWeights["E"]
	}
	if input.F != nil {
		c2 += float64(*input.F) * structs.ProductWeights["F"]
	}
	if input.G != nil {
		c3 += float64(*input.G) * structs.ProductWeights["G"]
	}
	if input.H != nil {
		c3 += float64(*input.H) * structs.ProductWeights["H"]
	}
	if input.I != nil {
		c3 += float64(*input.I) * structs.ProductWeights["I"]
	}

	return structs.WarehouseCenterDemandQuantity{
		C1: &c1,
		C2: &c2,
		C3: &c3,
	}
}

// Helper function to find the minimum non-zero cost
func minNonZero(values ...float64) float64 {
	minCost := float64(0)
	for _, v := range values {
		if v > 0 && (minCost == 0 || v < minCost) {
			minCost = v
		}
	}
	return minCost
}

func GetMinCostService(inputDemand structs.WarehouseCenterDemandQuantity) float64 {
	var cost1, cost2, cost3 float64

	if inputDemand.C1 != nil && *inputDemand.C1 > 0 {
		cost1 = GetMinCostStartingAtWarehouse("C1", inputDemand)
	}
	if inputDemand.C2 != nil && *inputDemand.C2 > 0 {
		cost2 = GetMinCostStartingAtWarehouse("C2", inputDemand)
	}
	if inputDemand.C3 != nil && *inputDemand.C3 > 0 {
		cost3 = GetMinCostStartingAtWarehouse("C3", inputDemand)
	}

	// Return the minimum non-zero cost
	return minNonZero(cost1, cost2, cost3)
}

func GetMinCostStartingAtWarehouse(initialWarehouse string, inputDemand structs.WarehouseCenterDemandQuantity) float64 {
	var cost float64

	// Validate the initial warehouse and calculate cost based on its demand
	switch initialWarehouse {
	case "C1":
		if inputDemand.C1 != nil && *inputDemand.C1 > 0 {
			C1L1InputDemand := inputDemand
			C1L1InputDemand.C1 = nil
			C1L1 := GetPathCost(structs.WarehouseClientDistance["C1"], *inputDemand.C1) + GetMinCostStartingAtClient(C1L1InputDemand)
			//C1L1 := *inputDemand.C1*structs.WarehouseClientDistance["C1"] + GetMinCostStartingAtClient(C1L1InputDemand)
			if inputDemand.C2 != nil && *inputDemand.C2 > 0 {
				// Create a copy of inputDemand for C1C2
				C1C2InputDemand := inputDemand
				*C1C2InputDemand.C2 += *inputDemand.C1
				C1C2 := GetPathCost(structs.WarehouseClientDistance["C1C2"], *inputDemand.C1) + GetMinCostStartingAtWarehouse("C2", C1C2InputDemand)
				//C1C2 := *inputDemand.C1*structs.WarehouseClientDistance["C1C2"] + GetMinCostStartingAtWarehouse("C2", C1C2InputDemand)
				cost = min(C1L1, C1C2)
			} else {
				cost = C1L1
			}
		} else {
			return 0.0
		}

	case "C2":
		if inputDemand.C2 != nil && *inputDemand.C2 > 0 {
			C2L1InputDemand := inputDemand
			C2L1InputDemand.C2 = nil
			C2L1 := GetPathCost(structs.WarehouseClientDistance["C2"], *inputDemand.C2) + GetMinCostStartingAtClient(C2L1InputDemand)
			//C2L1 := *inputDemand.C2*structs.WarehouseClientDistance["C2"] + GetMinCostStartingAtClient(C2L1InputDemand)
			var C2C1, C2C3 float64
			if inputDemand.C1 != nil && *inputDemand.C1 > 0 {
				C2C1InputDemand := inputDemand
				*C2C1InputDemand.C1 += *inputDemand.C2
				C2C1 = GetPathCost(structs.WarehouseClientDistance["C2C1"], *inputDemand.C1) + GetMinCostStartingAtWarehouse("C3", C2C1InputDemand)

			}
			if inputDemand.C3 != nil && *inputDemand.C3 > 0 {
				C2C3InputDemand := inputDemand
				*C2C3InputDemand.C3 += *inputDemand.C2
				C2C3 = GetPathCost(structs.WarehouseClientDistance["C2C3"], *inputDemand.C1) + GetMinCostStartingAtWarehouse("C1", C2C3InputDemand)

			}
			cost = minNonZero(C2L1, C2C3, C2C1)
		} else {
			return 0.0
		}

	case "C3":
		if inputDemand.C3 != nil && *inputDemand.C3 > 0 {
			C3L1InputDemand := inputDemand
			C3L1InputDemand.C3 = nil
			C3L1 := GetPathCost(structs.WarehouseClientDistance["C3"], *inputDemand.C3) + GetMinCostStartingAtClient(C3L1InputDemand)
			//C3L1 := *inputDemand.C3*structs.WarehouseClientDistance["C3"] + GetMinCostStartingAtClient(C3L1InputDemand)
			if inputDemand.C1 != nil && *inputDemand.C1 > 0 {
				// Create a copy of inputDemand for C1C2
				C3C1InputDemand := inputDemand
				*C3C1InputDemand.C2 += *inputDemand.C1
				C1C2 := GetPathCost(structs.WarehouseClientDistance["C3C1"], *inputDemand.C3) + GetMinCostStartingAtWarehouse("C3", C3C1InputDemand)
				//	C1C2 := *inputDemand.C3*structs.WarehouseClientDistance["C3C1"] + GetMinCostStartingAtWarehouse("C3", C3C1InputDemand)
				cost = min(C3L1, C1C2)
			} else {
				cost = C3L1
			}
		} else {
			return 0.0
		}

	default:
		// Invalid warehouse identifier
		return 0.0
	}

	return cost
}

func GetMinCostStartingAtClient(inputDemand structs.WarehouseCenterDemandQuantity) float64 {
	var cost1 float64
	var cost2 float64
	var cost3 float64

	if inputDemand.C1 != nil && *inputDemand.C1 > 0 {
		cost1 = 10*structs.WarehouseClientDistance["C1"] + GetMinCostStartingAtWarehouse("C1", inputDemand)
	}
	if inputDemand.C2 != nil && *inputDemand.C2 > 0 {
		cost2 = 10*structs.WarehouseClientDistance["C2"] + GetMinCostStartingAtWarehouse("C2", inputDemand)
	}
	if inputDemand.C3 != nil && *inputDemand.C3 > 0 {
		cost3 = 10*structs.WarehouseClientDistance["C3"] + GetMinCostStartingAtWarehouse("C3", inputDemand)
	}

	return min(cost1, cost2, cost3)
}

func GetPathCost(distance, weight float64) float64 {
	if weight <= 0 || distance <= 0 {
		return 0.0 // No cost for invalid or zero inputs
	}

	cost := 0.0

	// Calculate cost for the initial 5 kgs
	if weight <= 5 {
		cost = 10 * distance
	} else {
		// Cost for the first 5 kgs
		cost = 10 * distance

		// Calculate remaining weight
		remainingWeight := weight - 5

		// Calculate the multiple and remainder of dividing remaining weight by 5
		multiple := float64(int(remainingWeight) / 5)
		remainder := float64(int(remainingWeight) % 5)

		//cost += abc * (8 * float64(int(remainingWeight/5)*5))
		cost += multiple * (8 * distance)
		// Handle fractional weights (e.g., remaining weight less than 5)
		if remainder > 0 {
			cost += 8 * distance
		}
	}

	return cost
}
