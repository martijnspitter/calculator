package calculator

import (
	"errors"
	"math"
	"strconv"
)

type Calculator struct {
	input []string
}

func NewCalculator(input []string) Calculator {
	return Calculator{
		input: input,
	}
}

func (c *Calculator) Calculate() (float64, error) {
	output := []float64{}

	for len(c.input) > 0 {
		token := c.input[0]
		c.input = c.input[1:]

		if isOperator(token) {
			b := output[len(output)-1]
			output = output[:len(output)-1]
			a := output[len(output)-1]
			output = output[:len(output)-1]

			result := performOperation(a, b, token)
			output = append(output, result)
		} else {
			number, _ := strconv.ParseFloat(token, 64)
			output = append(output, number)
		}
	}

	if len(output) != 1 {
		return 0, errors.New("invalid input")
	}
	return round(output[0], 10), nil
}

func round(num float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(num*multiplier) / multiplier
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func performOperation(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	}

	return 0
}
