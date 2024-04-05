package main

import (
	"fmt"

	"github.com/bedlamzd/gameton/internal/api/client"
	"github.com/bedlamzd/gameton/internal/api/model"
)

func main() {
	fmt.Printf("%#v", client.GetUniverse())
	// fmt.Printf("%#v", client.GetRounds())
	res, err := client.Travel(&[]string{"Earth", "Moen"})
	fmt.Printf("%#v", res)
	fmt.Printf("%#v", err)
	gb := model.GarbageMap{}
	for key, sh := range res.PlanetGarbage {
		gb[key] = sh
		break
	}
	fmt.Printf("%#v", client.Collect(&gb))
}
