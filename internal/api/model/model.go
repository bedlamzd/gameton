package model

import (
	"encoding/json"
	"fmt"
)

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

func (t *UniverseEdge) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "" {
		return nil
	}

	var edge []interface{}
	if err := json.Unmarshal(data, &edge); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(edge)
	from := edge[0].(string)
	to := edge[1].(string)
	cost := edge[2].(float64)

	*t = UniverseEdge{
		From: from,
		To:   to,
		Cost: int(cost),
	}
	return nil
}

type UniverseResponse struct {
	Name     string
	Ship     Ship
	Universe []UniverseEdge
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
