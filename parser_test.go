package equation

import (
	"testing"
)

func TestParse(t *testing.T) {
	// Valid equations
	t.Run("x=5.3", func(t *testing.T) { validateEquationParsing(t, "x=5.3", 1, 0, 5.3) })
	t.Run("y=-5", func(t *testing.T) { validateEquationParsing(t, "y=-5", 0, 1, -5) })
	t.Run("-x=y", func(t *testing.T) { validateEquationParsing(t, "-x=y", -1, -1, 0) })
	t.Run("y=-x", func(t *testing.T) { validateEquationParsing(t, "y=-x", 1, 1, 0) })
	t.Run("X+y=.3", func(t *testing.T) { validateEquationParsing(t, "x+y=.3", 1, 1, .3) })
	t.Run("x + -y= 0", func(t *testing.T) { validateEquationParsing(t, "x + -y= 0", 1, -1, 0) })

	// Invalid equations
	t.Run("= 0", func(t *testing.T) { checkInvaidEquationParsing(t, "= 0") })
	t.Run("x+y=", func(t *testing.T) { checkInvaidEquationParsing(t, "x+y=") })
	t.Run("x-=+", func(t *testing.T) { checkInvaidEquationParsing(t, "x-=+") })
	t.Run("x+.=8", func(t *testing.T) { checkInvaidEquationParsing(t, "x+.=8") })
	t.Run("x=y+", func(t *testing.T) { checkInvaidEquationParsing(t, "x=y+") })
	t.Run("+x=0", func(t *testing.T) { checkInvaidEquationParsing(t, "+x=0") })
	t.Run("+-y=0", func(t *testing.T) { checkInvaidEquationParsing(t, "+-y=0") })
}

func validateEquationParsing(t *testing.T, equationString string, firstCoef float64, secondCoef float64, equationConst float64) {
	parser := *NewParser()
	equation, _ := parser.Parse(equationString)

	if equation.FirstCoefiicient != firstCoef || equation.SecondCoefiicient != secondCoef || equation.EquationConstant != equationConst {
		t.Fail()
	}
}

func checkInvaidEquationParsing(t *testing.T, equationString string) {
	parser := *NewParser()
	_, err := parser.Parse(equationString)

	if err == nil {
		t.Fail()
	}
}
