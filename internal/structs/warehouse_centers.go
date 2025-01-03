package structs

var ProductWeights = map[string]float64{
	"A": 3.0,
	"B": 2.0,
	"C": 8.0,
	"D": 12.0,
	"E": 25.0,
	"F": 15.0,
	"G": 0.5,
	"H": 1.0,
	"I": 2.0,
}

type WarehouseCenterDemandQuantity struct {
	C1 *float64
	C2 *float64
	C3 *float64
}

type ProductDemandQuantity struct {
	A *int
	B *int
	C *int
	D *int
	E *int
	F *int
	G *int
	H *int
	I *int
}

var WarehouseClientDistance = map[string]float64{
	"C1": 3.0,
	"C2": 2.5,
	"C3": 2.0,
}

var WarehouseDistance = map[string]float64{
	"C1C2": 4.0,
	"C2C1": 4.0,
	"C1C3": 5.0,
	"C3C1": 5.0,
	"C2C3": 3.0,
	"C3C2": 3.0,
}
