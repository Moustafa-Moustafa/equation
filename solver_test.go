package equation

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("Valid", validEquation)
	t.Run("Invalid", invalidEquation)
	t.Run("NoSolution", noSolutionEquation)
}

func validEquation(t *testing.T) {
	e1 := &Equation{1.5, -.5, 4.25}
	e2 := &Equation{2, 1, 9}
	solver := *NewSolver()
	sol, _ := solver.Solve([]Equation{*e1, *e2})

	fmt.Printf("Solution: x = %f, y = %f\n", sol.FirstUnknown, sol.SecondUnknown)

	if sol.FirstUnknown != 3.5 || sol.SecondUnknown != 2 {
		t.Fail()
	}
}

func invalidEquation(t *testing.T) {
	e1 := &Equation{1, 2, 3}
	e2 := &Equation{2, 4, 6}
	solver := *NewSolver()
	_, err := solver.Solve([]Equation{*e1, *e2})

	if err == nil {
		t.Fail()
	}
}

func noSolutionEquation(t *testing.T) {
	e1 := &Equation{1, 2, 6}
	e2 := &Equation{2, 4, 6}
	solver := *NewSolver()
	_, err := solver.Solve([]Equation{*e1, *e2})

	if err == nil {
		t.Fail()
	}
}
