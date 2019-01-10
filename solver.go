package equation

type Solver interface {
	Solve(equation []Equation) (Solution, error)
}

func NewSolver() *Solver {
	var s Solver = SimpleSolver{}
	return &s
}
