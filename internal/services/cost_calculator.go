package services

import (
	"math"
)

func GetDistance(from, to string) float64 {
	distances := map[string]map[string]float64{
		"C1": {"L1": 10, "C2": 20, "C3": 30},
		"C2": {"L1": 15, "C1": 20, "C3": 25},
		"C3": {"L1": 20, "C1": 30, "C2": 25},
	}
	return distances[from][to]
}

func CalculateMinimumCost(products map[string]int) float64 {
	warehouses := []string{"C1", "C2", "C3"}
	numWarehouses := len(warehouses)
	subsetCount := 1 << numWarehouses

	// Map of products available in each warehouse
	warehouseStock := map[string]map[string]bool{
		"C1": {"A": true, "B": true, "C": true},
		"C2": {"D": true, "E": true, "F": true},
		"C3": {"G": true, "H": true, "I": true},
	}

	// Calculate the weight of products required from each warehouse
	weights := make(map[string]float64)
	for product, qty := range products {
		for warehouse, stock := range warehouseStock {
			if stock[product] {
				weights[warehouse] += float64(qty) * 0.5 // Assuming 0.5kg per product
			}
		}
	}

	dp := make([]float64, subsetCount)
	for i := range dp {
		dp[i] = math.Inf(1)
	}
	dp[0] = 0

	// Dynamic programming to calculate minimum cost
	for mask := 1; mask < subsetCount; mask++ {
		for last := 0; last < numWarehouses; last++ {
			if mask&(1<<last) == 0 {
				continue
			}
			prevMask := mask ^ (1 << last)
			for prev := 0; prev < numWarehouses; prev++ {
				if prevMask&(1<<prev) == 0 && prev != last {
					continue
				}
				dp[mask] = math.Min(
					dp[mask],
					dp[prevMask]+weights[warehouses[last]]*GetDistance(warehouses[last], "L1"),
				)
			}
		}
	}

	return dp[subsetCount-1]
}
