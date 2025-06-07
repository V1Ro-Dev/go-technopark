package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	uniq2 "go-technopark/HW1/task_1/uniq"
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

	if len(flag.Args()) > 2 {
		return nil, nil, errors.New("Too mush params\nUsage: go run task_1.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

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

	return in, out, nil
}

func output(out io.Writer, str string) error {
	_, err := out.Write([]byte(str + "\n"))
	if err != nil {
		return fmt.Errorf("ошибка при записи в выходной поток: %w", err)
	}

	return nil
}

func main() {
	options, err := uniq2.GetParsedFlags()
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

	uniqStrings, err := uniq2.Uniq(strings, options)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range uniqStrings {
		err = output(out, str)
		if err != nil {
			log.Fatal(err)
		}
	}

}
