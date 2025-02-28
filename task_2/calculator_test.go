package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccessfulCases(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		out  float64
		err  error
	}{
		{
			"Testing 1 operator",
			"1+1.5",
			2.5,
			nil,
		},
		{
			"Testing 1 operator with brackets",
			"(1.2+3.6)",
			4.8,
			nil,
		},
		{
			"Testing 2 operators with brackets",
			"(1+3.3)*4.2",
			18.06,
			nil,
		},
		{
			"Testing priority of operations",
			"1+5*6.2",
			32,
			nil,
		},
		{
			"Testing multiple operators with brackets and priority of operations",
			"-(-11-(1*20/2)-11/2*3)",
			37.5,
			nil,
		},
		{
			"Testing multiple operators with brackets and priority of operations",
			"(6-4/2)*(7+6.5*3+2)",
			114,
			nil,
		},
		{
			"Testing single number",
			"123",
			123,
			nil,
		},
		{
			"Testing wrong parenthesis sequence",
			"(1+1.5))",
			0,
			errors.New("incorrect parenthesis sequence was given"),
		},
		{
			"Testing wrong usage of operators",
			"1+*1.5",
			0,
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong usage of operators",
			"*1+1.5",
			0,
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong usage of operators",
			"1+1..5",
			0,
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong chars in the expression",
			"5%2",
			0,
			errors.New("incorrect chars were used in the expression"),
		},
		{
			"Testing empty expression",
			"",
			0,
			errors.New("empty expression was given"),
		},
	}
	for _, test := range tests {
		actual, err := calculate(test.in)
		require.Equal(t, test.err, err, test.name)
		require.Equal(t, test.out, actual, test.name)
	}
}
