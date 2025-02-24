package main

import (
	"errors"
	"fmt"
	"go-technopark/task_2/collections"
	"os"
	"strconv"
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
	for i := 1; i < len(s); i++ {
		if operators[string(s[i])] && operators[string(s[i-1])] {
			return false
		}
	}
	if string(s[0]) != "-" && (operators[string(s[0])] || operators[string(s[len(s)-1])]) {
		return false
	}
	return true
}

func isValid(str string) bool {
	return len(str) > 0 && isValidStr(str) && isValidParentheses(str) && isValidOperators(str)
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
	if !isValidStr(str) {
		return []string{""}
	}
	newExpression := []string{}
	num := ""
	if string(str[0]) == "-" { // заменяем "-" в начале строки на "0-"
		str = "0" + str
	}
	for i, char := range str {
		if unicode.IsDigit(char) || string(char) == "." {
			num += string(char)
		} else {
			if num != "" {
				newExpression = append(newExpression, num)
				newExpression = append(newExpression, string(char))
				num = ""
			} else {
				if string(char) == "-" && string(str[i-1]) == "(" { // обработка случая, когда "-" идет сразу после скобки
					num = "-"
					continue
				}
				newExpression = append(newExpression, string(char))
			}
		}
	}
	if num != "" {
		newExpression = append(newExpression, num)
	}
	return newExpression
}

func calc(operator string, operands *collections.Stack[float64]) error {
	operand1 := operands.Pop()
	operand2 := operands.Pop()
	res := 0.0

	switch operator {
	case "+":
		res = operand2 + operand1

	case "-":
		res = operand2 - operand1

	case "*":
		res = operand2 * operand1

	case "/":
		if operand1 == 0 {
			err := errors.New("division by zero")
			return err
		}
		res = operand2 / operand1
	}
	operands.Push(res)
	return nil
}

func calculate(slicedExpression []string) float64 {
	priorities := getPriorities()
	operators := collections.Stack[string]{}
	operands := collections.Stack[float64]{}

	for _, str := range slicedExpression {
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			operands.Push(num)
		} else if str == "(" {
			operators.Push(str)
		} else if str == ")" {
			for !operators.IsEmpty() && operators.Top() != "(" {
				operator := operators.Pop()
				err = calc(operator, &operands)
				if err != nil {
					fmt.Println(err)
					return 0
				}
			}
			_ = operators.Pop()
		} else {
			for !operators.IsEmpty() && priorities[str] <= priorities[operators.Top()] {
				operator := operators.Pop()
				err := calc(operator, &operands)
				if err != nil {
					return 0
				}
			}
			operators.Push(str)
		}
	}

	for !operators.IsEmpty() {
		operator := operators.Pop()
		err := calc(operator, &operands)
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}
	return operands.Pop()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("empty expression was given")
		return
	}
	expression := os.Args[1]
	if !isValid(expression) {
		fmt.Println("Wrong expression was given")
		return
	}
	fmt.Println(calculate(parseExpression(expression)))
}
