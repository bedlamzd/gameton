package main

import (
	"fmt"

	"github.com/bedlamzd/gameton/internal/api/client"
)

func main() {
	fmt.Printf("%#v", client.GetUniverse())
}
