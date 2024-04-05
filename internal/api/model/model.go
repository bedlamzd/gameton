package model

type (
	GarbageID    = string
	GarbageShape = [][]int
	GarbageMap   = map[GarbageID]GarbageShape
)

type Planet struct {
	Name    string
	Garbage GarbageMap
}

type Ship struct {
	CapacityX int
	CapacityY int
	FuelUsed  int
	Garbage   GarbageMap
	Planet    Planet
}

type UniverseEdge struct {
	From string
	To   string
	Cost int
}

type Universe = []UniverseEdge

type UniverseResponse struct {
	Name     string
	Ship     Ship
	Universe [][]any
}

type PlanetDiff struct {
	From string
	Fuel int
	To   string
}

type TravelResponse struct {
	FuelDiff      int
	PlanetDiffs   []PlanetDiff
	PlanetGarbage GarbageMap
	ShipGarbage   GarbageMap
}
type CollectResponse struct {
	Garbage GarbageMap
	Leaved  []GarbageID
}

type RoundInfo struct {
	StartAt     string
	EndAt       string
	IsCurrent   bool
	Name        string
	PlanetCount int
}

type RoundsResponse struct {
	Rounds []RoundInfo
}

type ResetRoundResponse struct {
	Success bool
}
