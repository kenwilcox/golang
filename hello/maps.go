package main

import "fmt"

type Geo struct {
	Lat, Long float64
}

var m = map[string]Geo{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

var g map[string]Geo

func main() {
	g = make(map[string]Geo)
	g["Bell Labs"] = Geo{40.68433, -74.39967}
	fmt.Println(g["Bell Labs"])

	fmt.Println(m)
}
