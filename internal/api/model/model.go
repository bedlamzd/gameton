package model

type (
	GarbageID    = string
	GarbageShape = [][]int
	GarbageMap   = map[GarbageID]GarbageShape
)

type Planet struct {
	Name string
	Garbage GarbageMap
}

type Ship struct {
	CapacityX int
	CapacityY int
	FuelUsed int
	Garbage GarbageMap
	Planet Planet
}

type UniverseEdge struct {
	From string
	To string
	Cost int
}

type Universe = []UniverseEdge

type UniverseResponse struct {
	Name string
	Ship Ship
	Universe [][]any
}

