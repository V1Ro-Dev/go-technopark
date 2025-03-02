package main

import (
	"errors"
	"fmt"
	"go-technopark/task_2/collections"
	"log"
	"math/big"
	"os"
	"unicode"
)

func isValidParentheses(s string) bool {

	cnt := 0
	for _, ch := range s {
		if string(ch) == "(" {
			cnt++
		} else if string(ch) == ")" {
			cnt--
		} else if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func isValidStr(s string) bool {

	chars := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"(": true,
		")": true,
		".": true,
		"0": true,
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
	}

	for _, ch := range s {
		if !chars[string(ch)] {
			return false
		}
	}

	return true
}

func isValidOperators(s string) bool {

	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		".": true,
	}

	for i, _ := range s[:len(s)-1] {
		if operators[string(s[i])] && operators[string(s[i+1])] {
			return false
		}
	}

	if string(s[0]) != "-" && (operators[string(s[0])] || operators[string(s[len(s)-1])]) {
		return false
	}

	return true
}

func isValid(str string) error {

	switch {
	case len(str) == 0:
		return errors.New("empty expression was given")

	case !isValidStr(str):
		return errors.New("incorrect chars were used in the expression")

	case !isValidOperators(str):
		return errors.New("incorrect usage of operators")

	case !isValidParentheses(str):
		return errors.New("incorrect parenthesis sequence was given")

	default:
		return nil
	}
}

func getPriorities() map[string]int {

	return map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
		"(": -1,
		")": -1,
	}
}

func parseExpression(str string) []string {

	var newExpression []string
	num := ""

	if string(str[0]) == "-" { // заменяем "-" в начале строки на "0-"
		str = "0" + str
	}

	for i, char := range str {
		if unicode.IsDigit(char) || string(char) == "." {
			num += string(char)
			continue
		}

		if num != "" {
			newExpression = append(newExpression, num)
			num = ""
		}

		if string(char) == "-" && string(str[i-1]) == "(" { // обработка случая, когда "-" идет сразу после скобки
			num = "-"
			continue
		}
		newExpression = append(newExpression, string(char))
	}

	if num != "" {
		newExpression = append(newExpression, num)
	}

	return newExpression
}

func calc(operator string, operands *collections.Stack[*big.Float]) error {

	operand1 := operands.Pop()
	operand2 := operands.Pop()
	res := new(big.Float)

	switch operator {
	case "+":
		res = new(big.Float).Add(operand2, operand1)

	case "-":
		res = new(big.Float).Sub(operand2, operand1)

	case "*":
		res = new(big.Float).Mul(operand2, operand1)

	case "/":
		if operand1.Cmp(big.NewFloat(0)) == 0 {
			err := errors.New("division by zero")
			return err
		}
		res = new(big.Float).Quo(operand2, operand1)
	}
	operands.Push(res)

	return nil
}

func calculate(expression string) (*big.Float, error) {

	if err := isValid(expression); err != nil {
		return big.NewFloat(0), err
	}

	slicedExpression := parseExpression(expression)
	priorities := getPriorities()
	operators := collections.Stack[string]{}
	operands := collections.Stack[*big.Float]{}

	for _, str := range slicedExpression {
		switch num, _, err := big.ParseFloat(str, 10, 53, big.ToNearestEven); {

		case err == nil:
			operands.Push(num)

		case str == "(":
			operators.Push(str)

		case str == ")":
			for !operators.IsEmpty() && operators.Top() != "(" {
				operator := operators.Pop()
				if err = calc(operator, &operands); err != nil {
					return big.NewFloat(0), err
				}
			}
			_ = operators.Pop()

		default:
			for !operators.IsEmpty() && priorities[str] <= priorities[operators.Top()] {
				operator := operators.Pop()
				if err = calc(operator, &operands); err != nil {
					return big.NewFloat(0), err
				}
			}
			operators.Push(str)
		}
	}

	for !operators.IsEmpty() {
		operator := operators.Pop()
		if err := calc(operator, &operands); err != nil {
			return big.NewFloat(0), err
		}
	}

	return operands.Pop(), nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("empty expression was given")
	}

	expression := os.Args[1]
	result, err := calculate(expression)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
