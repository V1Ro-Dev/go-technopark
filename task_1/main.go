package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-technopark/task_1/uniq"
	"os"
)

func getStrings(in *os.File) ([]string, error) {
	scanner := bufio.NewScanner(in)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	options, err := uniq.GetParsedFlags()
	if err != nil {
		fmt.Println("Correct usage: go run main.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}

	in := os.Stdin
	out := os.Stdout
	switch len(flag.Args()) {
	case 1:
		in, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer in.Close()

	case 2:
		in, err = os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer in.Close()
		out, err = os.Create(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()

	case 0:

	default:
		fmt.Println("Too mush params\nUsage: go run task_1.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}

	strings, err := getStrings(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, str := range uniq.Uniq(strings, options) {
		_, err = out.WriteString(str + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
