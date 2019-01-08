package equation

type Solver interface {
	Solve(equation []Equation) Solution
}

func NewSolver() *Solver {
	var s Solver = SimpleSolver{}
	return &s
}
