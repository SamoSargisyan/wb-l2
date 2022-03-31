package main

import (
	"fmt"
	"l2/pattern/visitor/pkg"
)

func main() {
	square := &pkg.Square{Side: 2}
	circle := &pkg.Circle{Radius: 3}
	rectangle := &pkg.Rectangle{L: 2, B: 3}

	areaCalculator := &areaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
