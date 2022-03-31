package main

import (
	"fmt"
	"l2/pattern/visitor/pkg"
)

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) VisitForSquare(s *pkg.Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) VisitForCircle(c *pkg.Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) VisitForRectangle(t *pkg.Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
