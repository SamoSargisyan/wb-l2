package pkg

type Shape interface {
	GetType() string
	Accept(Visitor)
}
