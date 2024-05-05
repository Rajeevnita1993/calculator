package calc

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Calculate(expression string) {

	postfix, err := infixToPostfix(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Evaluate the postfix expression
	result, err := evaluatePostfix(postfix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// operator, a, b, err := parseExpression(expression)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// // Perform the arithmetic operations
	// var result float64
	// switch operator {
	// case "+":
	// 	result = a + b
	// case "-":
	// 	result = a - b
	// case "*":
	// 	result = a * b
	// case "/":
	// 	if b == 0 {
	// 		fmt.Println("Error: Division by zero")
	// 		return
	// 	}
	// 	result = a / b
	// default:
	// 	fmt.Println("Error: Invalid Operator")
	// 	return

	// }

	fmt.Println(result)
}

func parseExpression(expr string) (string, float64, float64, error) {
	// Split the expression into operator & operands
	var operator string
	var a, b float64

	n, err := fmt.Sscanf(expr, "%f %s %f", &a, &operator, &b)
	if err != nil {
		return "", 0, 0, err
	}

	if n != 3 {
		return "", 0, 0, fmt.Errorf("invalid expression")
	}
	return operator, a, b, nil
}

func infixToPostfix(expression string) (string, error) {
	var output strings.Builder
	var stack []string

	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
	tokens := tokenize(expression)

	for _, token := range tokens {

		switch token {
		case "+", "-", "*", "/":
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] && stack[len(stack)-1] != "(" {
				output.WriteString(stack[len(stack)-1])
				output.WriteString(" ")
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output.WriteString(stack[len(stack)-1])
				output.WriteString(" ")
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return "", fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		default:
			output.WriteString(token)
			output.WriteString(" ")
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return "", fmt.Errorf("mismatched parentheses")
		}
		output.WriteString(stack[len(stack)-1])
		output.WriteString(" ")
		stack = stack[:len(stack)-1]
	}

	return strings.TrimSpace(output.String()), nil
}

func evaluatePostfix(expression string) (float64, error) {
	var stack []float64

	scanner := bufio.NewScanner(strings.NewReader(expression))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()

		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

func tokenize(expression string) []string {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expression {
		if char == ' ' {
			continue
		}

		if char == '(' || char == ')' || char == '+' || char == '-' || char == '*' || char == '/' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
