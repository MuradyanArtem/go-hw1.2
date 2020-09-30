package rpn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var numbers = "0123456789"
var priorities = map[string]uint8{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

var (
	errInvalidExpr  = errors.New("invalid expression")
	errDivideByZero = errors.New("invalid expression, divide by zero")
	errInvalidCond  = errors.New("invalid condition")
)

// New create reverse polish notation from stirng
func New(in string) (out []string, err error) {
	operatorStack := stack.New()
	for idx := 0; idx < len(in); idx++ {
		switch in[idx] {
		case ' ':

		case ')':
			for operator := operatorStack.Pop(); (operatorStack.Len() > 0) && (operator.(string) != "("); {
				out = append(out, operator.(string))
				operator = operatorStack.Pop()
			}

		case '(':
			operatorStack.Push(string(in[idx]))

		case '*', '/', '+', '-':
			if operatorStack.Len() > 0 {
				if operator := operatorStack.Peek(); priorities[string(in[idx])] <= priorities[operator.(string)] {
					out = append(out, operator.(string))
					operatorStack.Pop()
				}
			}
			operatorStack.Push(string(in[idx]))

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf := []byte{}
			for idx < len(in) && strings.Contains(numbers, string(in[idx])) {
				buf = append(buf, in[idx])
				idx++
			}
			idx--

			out = append(out, string(buf))

		default:
			return nil, fmt.Errorf("input has invalid token - %v", in[idx])
		}
	}

	for operatorStack.Len() > 0 {
		operator := operatorStack.Pop()
		out = append(out, operator.(string))
	}
	return
}

// Calculate expression in rpn.
func Calculate(in []string) (res int, err error) {
	values := stack.New()
	for _, el := range in {
		switch el {
		case "+":
			if values.Len() < 2 {
				return 0, errInvalidExpr
			}
			firstVal := values.Pop()
			secondVal := values.Pop()
			values.Push(firstVal.(int) + secondVal.(int))

		case "-":
			if values.Len() < 2 {
				return 0, errInvalidExpr
			}
			firstVal := values.Pop()
			secondVal := values.Pop()
			values.Push(secondVal.(int) - firstVal.(int))

		case "*":
			if values.Len() < 2 {
				return 0, errInvalidExpr
			}
			firstVal := values.Pop()
			secondVal := values.Pop()
			values.Push(firstVal.(int) * secondVal.(int))

		case "/":
			if values.Len() < 2 {
				return 0, errInvalidExpr
			}
			firstVal := values.Pop()
			secondVal := values.Pop()
			if firstVal == 0 {
				return 0, errDivideByZero
			}
			values.Push(secondVal.(int) / firstVal.(int))

		default:
			num, err := strconv.Atoi(el)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %v", err)
			}
			values.Push(num)
		}
	}

	if values.Len() != 1 {
		return res, errInvalidCond
	}
	result := values.Pop()
	res = result.(int)

	return
}
