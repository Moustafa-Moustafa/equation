package equation

import "errors"

type SimpleSolver struct{}

// Solve the specified equation and sets the coifficients and the constant
func (s SimpleSolver) Solve(equation []Equation) (Solution, error) {

	if len(equation) != 2 {
		// error
		return Solution{0, 0}, errors.New("two equations are required")
	}

	e1 := equation[0]
	e2 := equation[1]

	denomirator := (e1.FirstCoefiicient*e2.SecondCoefiicient - e1.SecondCoefiicient*e2.FirstCoefiicient)

	if denomirator == 0 {
		// error
		return Solution{0, 0}, errors.New("the equation system doesn't have a determinstic output")
	}

	x := (e1.EquationConstant*e2.SecondCoefiicient - e1.SecondCoefiicient*e2.EquationConstant) / denomirator
	y := (e1.FirstCoefiicient*e2.EquationConstant - e1.EquationConstant*e2.FirstCoefiicient) / denomirator

	var sol = &Solution{x, y}
	return *sol, nil
}
