package shuntingyard

import (
	"strconv"
)

type Parser struct {
	tokens []string
}

func NewParser(input string) Parser {
	return Parser{
		tokens: tokenize(input),
	}
}

func tokenize(input string) []string {
	var tokens []string
	var currentNumber string

	for i := 0; i < len(input); i++ {
		char := string(input[i])

		// Skip whitespace
		if char == " " {
			continue
		}

		// If it's a digit or decimal point, build the number
		if isDigit(char) || char == "." {
			currentNumber += char
			// If it's the last character or next char is not a digit/decimal
			if i == len(input)-1 || !(isDigit(string(input[i+1])) || string(input[i+1]) == ".") {
				tokens = append(tokens, currentNumber)
				currentNumber = ""
			}
			continue
		}

		// If we have a pending number, add it to tokens
		if currentNumber != "" {
			tokens = append(tokens, currentNumber)
			currentNumber = ""
		}

		// Handle operators and parentheses
		if isOperator(char) || isLeftParenthesis(char) || isRightParenthesis(char) {
			tokens = append(tokens, char)
		}
	}

	return tokens
}

func isDigit(s string) bool {
	if len(s) != 1 {
		return false
	}
	return s[0] >= '0' && s[0] <= '9'
}

func (p *Parser) Parse() []string {
	outputStack := []string{}
	operatorStack := []string{}

	for _, token := range p.tokens {
		token := string(token)
		if isNumber(token) {
			outputStack = append(outputStack, token)
		}
		if isOperator(token) {
			// While there's an operator on top of the stack with greater or equal precedence
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				if !isLeftParenthesis(top) && hasPrecedence(top, token) {
					// Pop the operator from the stack and add to output
					operatorStack = operatorStack[:len(operatorStack)-1]
					outputStack = append(outputStack, top)
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, token)
		}
		if isLeftParenthesis(token) {
			operatorStack = append(operatorStack, token)
		}
		if isRightParenthesis(token) {
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				if isLeftParenthesis(top) {
					break
				}
				outputStack = append(outputStack, top)
			}
		}
	}

	// Pop any remaining operators from the stack
	for len(operatorStack) > 0 {
		top := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		outputStack = append(outputStack, top)
	}

	return outputStack
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^" || token == "%"
}

func isLeftParenthesis(token string) bool {
	return token == "("
}

func isRightParenthesis(token string) bool {
	return token == ")"
}

func isNumber(token string) bool {
	// should handle floating point numbers
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func hasPrecedence(operator, token string) bool {
	return precendence(operator) >= precendence(token)
}

func precendence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/", "%":
		return 2
	case "^":
		return 3
	}
	return 0
}
