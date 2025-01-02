package models

var Distances = map[string]map[string]float64{
	"C1": {"C2": 4, "C3": 5, "L1": 3},
	"C2": {"C1": 4, "C3": 3, "L1": 2.5},
	"C3": {"C1": 5, "C2": 3, "L1": 2},
}

// GetDistance fetches the distance between two locations.
func GetDistance(from, to string) float64 {
	if dist, ok := Distances[from][to]; ok {
		return dist
	}
	return 1e9 // Large value for unreachable nodes
}
