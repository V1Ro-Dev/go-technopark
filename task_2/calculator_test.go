package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccessfulCases(t *testing.T) {
	var tests = []struct {
		in  string
		out float64
	}{
		{
			"1+1.5",
			2.5,
		},
		{
			"(1.2+3.6)",
			4.8,
		},
		{
			"(1+3.3)*4.2",
			18.06,
		},
		{
			"1+5*6.2",
			32,
		},
		{
			"-(-11-(1*20/2)-11/2*3)",
			37.5,
		},
		{
			"(6-4/2)*(7+6.5*3+2)",
			114,
		},
	}
	for _, test := range tests {
		require.Equal(t, test.out, calculate(parseExpression(test.in)))
	}
}

func TestNegativeCases(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{
			"(1+1.5))",
			false,
		},
		{
			"1+*1.5",
			false,
		},
		{
			"*1+1.5",
			false,
		},
		{
			"1+1..5",
			false,
		},
		{
			"",
			false,
		},
	}
	for _, test := range tests {
		require.Equal(t, test.out, isValid(test.in))
	}
}
