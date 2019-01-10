package equation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SimpleParser struct{}

// Parse the specified equation and sets the coifficients and the constant
func (p SimpleParser) Parse(equationString string) (Equation, error) {

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

func parseExpression(expression string) (Equation, error) {
	if expression == "" {
		return Equation{0, 0, 0}, fmt.Errorf("invalid expression %s", expression)
	}

	matches, err := findRegexMatches(expression)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	equationConst := 0.0
	if matches[9] != "" {
		equationConst, err = strconv.ParseFloat(matches[9], 64)
		if err != nil {
			return Equation{0, 0, 0}, err
		}

	}

	firstCoef, err := calculateCoefficient(matches, 0)
	if err != nil {
		return Equation{0, 0, 0}, err
	}

	secondCoef, err := calculateCoefficient(matches, 4)
	if err != nil {
		return Equation{0, 0, 0}, err
	}
	eq := &Equation{firstCoef, secondCoef, equationConst}

	return *eq, nil
}

func getEquationFromExpressions(lhsExpression Equation, rhsExpression Equation) Equation {
	return Equation{lhsExpression.FirstCoefiicient - rhsExpression.FirstCoefiicient,
		lhsExpression.SecondCoefiicient - rhsExpression.SecondCoefiicient,
		rhsExpression.EquationConstant - lhsExpression.EquationConstant}
}

func calculateCoefficient(matches []string, startIndex int) (float64, error) {
	xc := matches[startIndex+2]
	if matches[startIndex+1] != "" {
		if matches[startIndex+3] != "" {
			xc += matches[startIndex+3]
		} else {
			xc += "1"
		}
	} else {
		xc += "0"
	}

	return strconv.ParseFloat(xc, 64)
}

func findRegexMatches(equationString string) ([]string, error) {
	regexMatcher, err := regexp.Compile(buildRegexPattern())

	if err != nil {
		return nil, fmt.Errorf("can't parse equation %s", equationString)
	}

	matches := regexMatcher.FindStringSubmatch(equationString)

	if matches == nil {
		return nil, fmt.Errorf("can't parse equation %s", equationString)
	}

	// Check the operands exist
	if (matches[1] != "" && matches[5] != "" && matches[4] == "" && matches[6] == "") ||
		(matches[1] != "" && matches[9] != "" && matches[4] == "" && matches[10] == "") ||
		(matches[5] != "" && matches[9] != "" && matches[8] == "" && matches[10] == "") {
		return nil, fmt.Errorf("can't parse equation %s", equationString)
	}

	return matches, nil
}

func buildRegexPattern() string {
	return `^` + buildCoefficientMatchRegexPattern("x") + `(\+)?` + buildCoefficientMatchRegexPattern("y") + `(\+)?` + `((-)?(\d*\.?\d*))?` + `$`
}

func buildCoefficientMatchRegexPattern(unknown string) string {
	return `((-)?(\d*\.?\d*)?` + unknown + `)?`
}
