package equation

// Solver is an interface that solves linear equations
type Solver interface {
	Solve(equation []Equation) (Solution, error)
}

func NewSolver() *Solver {
	var s Solver = simpleSolver{}
	return &s
}
