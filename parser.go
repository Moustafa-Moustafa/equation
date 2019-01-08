package equation

type Parser interface {
	Parse(equationString string) Equation
}

func NewParser() *Parser {
	var p Parser = SimpleParser{}
	return &p
}
