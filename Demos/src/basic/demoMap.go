package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var n = map[string]Vertex {
	"Bell Labs" : Vertex {
		40.68433, -74.39967,
	},
	"Google": Vertex {
		37.42202, -122.08408,
	},
}

var l = map[string]Vertex {
	"Bell Labs" :  {
		40.68433, -74.39967,
	},
	"Google":{
		37.42202, -122.08408,
	},
}

func main () {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967,}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)

	fmt.Println(l)

	mapA := make(map[string]int)

	mapA["Answer"] = 42
	fmt.Println("The value:", mapA["Answer"])

	mapA["Answer"] = 48
	fmt.Println("The value:", mapA["Answer"])

	delete(mapA, "Answer")
	fmt.Println("The value:", mapA["Answer"])

	v, ok := mapA["Answer"]
	fmt.Println("The value:", v, "Prsent?", ok)
}