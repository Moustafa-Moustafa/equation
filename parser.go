package equation

// Parser is an interface used to parse equations
type Parser interface {
	Parse(equationString string) (Equation, error)
}

func NewParser() *Parser {
	var p Parser = simpleParser{}
	return &p
}
