package main

import (
	"fmt"

	route "github.com/x0n4d0/code-delivery-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:        "1",
		ClientID:  "1",
		Positions: nil,
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
