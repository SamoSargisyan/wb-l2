package pkg

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}
