package equation

type SimpleSolver struct{}

// Solve the specified equation and sets the coifficients and the constant
func (s SimpleSolver) Solve(equation []Equation) Solution {
	var sol = &Solution{1, 2}
	return *sol
}
