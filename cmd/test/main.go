package main

import (
	"fmt"

	"github.com/bedlamzd/gameton/internal/api/client"
)

func main() {
	universe := client.GetUniverse()
	fmt.Printf("%#v\n", universe)
	fmt.Printf("%#v\n", universe.Universe)
}
