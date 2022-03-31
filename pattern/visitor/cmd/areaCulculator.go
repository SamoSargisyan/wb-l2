package main

import (
	"fmt"
	"l2/pattern/visitor/pkg"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) VisitForSquare(s *pkg.Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) VisitForCircle(s *pkg.Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) VisitForRectangle(s *pkg.Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
