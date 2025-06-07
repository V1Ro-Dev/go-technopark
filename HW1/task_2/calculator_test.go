package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestSuccessfulCases(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		out  *big.Float
		err  error
	}{
		{
			"Testing 1 operator",
			"1+1.5",
			big.NewFloat(2.5).SetPrec(256),
			nil,
		},
		{
			"Testing 1 operator with brackets",
			"(1.2+3.6)",
			big.NewFloat(4.8).SetPrec(256),
			nil,
		},
		{
			"Testing 2 operators with brackets",
			"(1+3.3)*4.2",
			big.NewFloat(18.06).SetPrec(256),
			nil,
		},
		{
			"Testing priority of operations",
			"1+5*6.2",
			big.NewFloat(32).SetPrec(256),
			nil,
		},
		{
			"Testing multiple operators with brackets and priority of operations",
			"-(-11-(1*20/2)-11/2*3)",
			big.NewFloat(37.5).SetPrec(256),
			nil,
		},
		{
			"Testing multiple operators with brackets and priority of operations",
			"(6-4/2)*(7+6.5*3+2)",
			big.NewFloat(114).SetPrec(256),
			nil,
		},
		{
			"Testing single number",
			"123",
			big.NewFloat(123).SetPrec(256),
			nil,
		},
		{
			"Testing wrong parenthesis sequence",
			"(1+1.5))",
			big.NewFloat(0).SetPrec(256),
			errors.New("incorrect parenthesis sequence was given"),
		},
		{
			"Testing wrong usage of operators",
			"1+*1.5",
			big.NewFloat(0).SetPrec(256),
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong usage of operators",
			"*1+1.5",
			big.NewFloat(0).SetPrec(256),
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong usage of operators",
			"1+1..5",
			big.NewFloat(0).SetPrec(256),
			errors.New("incorrect usage of operators"),
		},
		{
			"Testing wrong chars in the expression",
			"5%2",
			big.NewFloat(0).SetPrec(256),
			errors.New("incorrect chars were used in the expression"),
		},
		{
			"Testing empty expression",
			"",
			big.NewFloat(0).SetPrec(256),
			errors.New("empty expression was given"),
		},
	}
	for _, test := range tests {
		actual, err := calculate(test.in)
		require.Equal(t, test.err, err, test.name)
		require.Zero(t, test.out.Cmp(actual), test.name)
	}
}
