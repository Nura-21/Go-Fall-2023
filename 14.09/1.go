package main

import "fmt"

type Vertex struct {
	X     int
	Y     int
	Value int
}

type BigVertex struct {
	vertex Vertex
}

func main() {
	var test = Vertex{1, 2, 3}
	fmt.Println(test.Value)
	var bigTest = BigVertex{Vertex{1, 2, 3}}
	fmt.Println(bigTest)
}
