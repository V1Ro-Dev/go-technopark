package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-technopark/task_1/uniq"
	"io"
	"log"
	"os"
)

func getStrings(in io.Reader) ([]string, error) {

	scanner := bufio.NewScanner(in)
	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return res, fmt.Errorf("ошибка при чтении строк: %w", err)
	}

	return res, nil
}

func input() (io.ReadCloser, io.WriteCloser, error) {

	in := os.Stdin
	out := os.Stdout
	var err error

	if len(flag.Args()) == 1 {
		in, err = os.Open(flag.Args()[0])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка при открытии файла: %w", err)
		}
	}

	if len(flag.Args()) == 2 {
		in, err = os.Open(flag.Args()[0])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка при открытии файла: %w", err)
		}
		out, err = os.Create(flag.Args()[1])
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка при создании файла: %w", err)
		}
	}

	if len(flag.Args()) > 2 {
		return nil, nil, fmt.Errorf("Too mush params\nUsage: go run task_1.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	return in, out, nil
}

func output(out io.Writer, str string) {
	var err error

	_, err = out.Write([]byte(str + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	options, err := uniq.GetParsedFlags()
	if err != nil {
		log.Fatal(err, "\nCorrect usage: go run main.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	in, out, err := input()
	if err != nil {
		log.Fatal(err)
	}

	if in != os.Stdout {
		defer in.Close()
	}
	if out != os.Stdout {
		defer out.Close()
	}

	strings, err := getStrings(in)
	if err != nil {
		log.Fatal(err)
	}

	uniqStrings, err := uniq.Uniq(strings, options)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range uniqStrings {
		output(out, str)
	}

}
