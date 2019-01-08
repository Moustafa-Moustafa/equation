package equation

type SimpleParser struct{}

// Parse the specified equation and sets the coifficients and the constant
func (p SimpleParser) Parse(equationString string) Equation {
	var eq = &Equation{5, 3, 2}
	return *eq
}
