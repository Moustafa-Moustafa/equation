package equation

import "errors"

var (
	// ErrEquationWithNoOrManyEqualSign is returned when the equation has a different number of equal signs than one.
	ErrEquationWithNoOrManyEqualSign = errors.New("equation must have one equal sign")
	// ErrNotFound is returned when no route match is found.
	ErrNotFound = errors.New("no matching route was found")
)

type Parser interface {
	Parse(equationString string) (Equation, error)
}

func NewParser() *Parser {
	var p Parser = SimpleParser{}
	return &p
}
