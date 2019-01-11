package equation

import (
	"errors"
	"math"
)

type simpleSolver struct{}

// Solve the specified equation and sets the coifficients and the constant
func (s simpleSolver) Solve(equation []Equation) (Solution, error) {

	if len(equation) != 2 {
		return Solution{0, 0}, errors.New("two equations are required")
	}

	e1 := equation[0]
	e2 := equation[1]

	if almostEqual(e1.FirstCoefiicient*e2.SecondCoefiicient, e1.SecondCoefiicient*e2.FirstCoefiicient) {
		return Solution{0, 0}, errors.New("the equation system doesn't have a determinstic output")
	}

	denomirator := (e1.FirstCoefiicient*e2.SecondCoefiicient - e1.SecondCoefiicient*e2.FirstCoefiicient)
	x := (e1.EquationConstant*e2.SecondCoefiicient - e1.SecondCoefiicient*e2.EquationConstant) / denomirator
	y := (e1.FirstCoefiicient*e2.EquationConstant - e1.EquationConstant*e2.FirstCoefiicient) / denomirator

	var sol = &Solution{x, y}
	return *sol, nil
}

func almostEqual(a float64, b float64) bool {
	if a == b {
		return true
	}

	absDiff := math.Abs(a - b)
	epsilon := .000000001
	if absDiff < epsilon {
		return true
	}

	return false
}
