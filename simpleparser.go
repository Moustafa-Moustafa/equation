package equation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type simpleParser struct{}

// Parse the specified equation and sets the coifficients and the constant
func (p simpleParser) Parse(equationString string) (Equation, error) {

	equationString = strings.ToLower(strings.Join(strings.Fields(equationString), ""))
	equationSides := strings.Split(equationString, "=")

	if len(equationSides) != 2 {
		return Equation{0, 0, 0}, fmt.Errorf("invalid equation %s: equation must have one equal sign", equationString)
	}

	lhsExpression, err := parseExpression(equationSides[0])

	if err != nil {
		return Equation{0, 0, 0}, fmt.Errorf("invalid left hand side of the equation %s", equationString)
	}

	rhsExpression, err := parseExpression(equationSides[1])

	if err != nil {
		return Equation{0, 0, 0}, fmt.Errorf("invalid right hand side of the equation %s", equationString)
	}

	simpleEquation := getEquationFromExpressions(lhsExpression, rhsExpression)

	return simpleEquation, nil
}

// Parse the specified expression
func parseExpression(expression string) (Equation, error) {
	if expression == "" || strings.Index(expression, "+") == 0 {
		return Equation{0, 0, 0}, fmt.Errorf("invalid expression %s", expression)
	}

	matches, err := findRegexMatches(expression)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	// Calculates the equation constant
	equationConst, err := calculateConstant(matches)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	// Calculates the equation first coefficient
	firstCoef, err := calculateCoefficient(matches, 0)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	// Calculates the equation second coefficient
	secondCoef, err := calculateCoefficient(matches, 3)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	eq := &Equation{firstCoef, secondCoef, equationConst}

	return *eq, nil
}

// Simplifies the equasion by moving the unknowns to LHS and the constants to RHS
func getEquationFromExpressions(lhsExpression Equation, rhsExpression Equation) Equation {
	return Equation{lhsExpression.FirstCoefiicient - rhsExpression.FirstCoefiicient,
		lhsExpression.SecondCoefiicient - rhsExpression.SecondCoefiicient,
		rhsExpression.EquationConstant - lhsExpression.EquationConstant}
}

// Calculates the equation constant based on the submatches
func calculateConstant(matches []string) (float64, error) {
	equationConstant := "0.0"
	if matches[7] != "" {
		if matches[8] == "+" || matches[8] == "+-" {
			equationConstant = matches[7][1:]
		} else {
			equationConstant = matches[7]
		}
	}

	return strconv.ParseFloat(equationConstant, 64)
}

// Calculates the equation coefficient based on the submatches
func calculateCoefficient(matches []string, startIndex int) (float64, error) {
	coefficient := ""

	if matches[startIndex+1] == "" {
		coefficient += "0"
	} else {
		// simplify it for the parsing
		if matches[startIndex+2] == "+-" || matches[startIndex+2] == "-" {
			coefficient = "-"
		}

		if matches[startIndex+3] != "" {
			coefficient += matches[startIndex+3]
		} else {
			// in case of unknown with no coeeficient (e.g "x")
			coefficient += "1"
		}
	}

	return strconv.ParseFloat(coefficient, 64)
}

// Finds the submatches used for equation construction
func findRegexMatches(equationString string) ([]string, error) {
	regexMatcher, err := regexp.Compile(buildRegexPattern())

	if err != nil {
		return nil, fmt.Errorf("can't parse equation %s", equationString)
	}

	matches := regexMatcher.FindStringSubmatch(equationString)

	if matches == nil {
		return nil, fmt.Errorf("can't parse equation %s", equationString)
	}

	return matches, nil
}

func buildRegexPattern() string {
	return `^` + buildCoefficientMatchRegexPattern("x") + buildCoefficientMatchRegexPattern("y") + `(` + buildNumberPattern() + `)?` + `$`
}

func buildCoefficientMatchRegexPattern(unknown string) string {
	return `(` + buildNumberPattern() + unknown + `)?`
}

func buildNumberPattern() string {
	return `(\+|\+-|-)?(\d*\.?\d*)?`
}
