package equation

type (
	Coefiicient      float32
	EquationConstant float32
	Unknown          float32
)

// Equation represntation
type Equation struct {
	FirstCoefiicient  Coefiicient
	SecondCoefiicient Coefiicient
	EquationConstant  EquationConstant
}

// Solution represntation
type Solution struct {
	FirstUnknown  Unknown
	SecondUnknown Unknown
}
