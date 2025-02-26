package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccessfulCases(t *testing.T) {
	var tests = []struct {
		name  string
		in    []string
		out   []string
		flags Options
		err   error
	}{
		{
			"test",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{},
			nil,
		},
		{
			"test",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
			Options{countOccurrences: true},
			nil,
		},
		{
			"test",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
			Options{repeatedStrings: true},
			nil,
		},
		{
			"test",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"", "Thanks."},
			Options{uniqStrings: true},
			nil,
		},
		{
			"test",
			[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik."},
			[]string{"I LOVE MUSIC.", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{caseInsensitive: true},
			nil,
		},
		{
			"test",
			[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"We love music.", "", "I love music of Kartik.", "Thanks."},
			Options{skipFields: 1},
			nil,
		},
		{
			"test",
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			Options{skipChars: 1},
			nil,
		},
	}
	for _, test := range tests {
		require.Equal(t, test.out, Uniq(test.in, test.flags))
	}
}

func TestNegativeCases(t *testing.T) {
	var tests = []struct {
		in    []string
		out   []string
		flags Options
	}{
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{countOccurrences: true, repeatedStrings: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{uniqStrings: true, repeatedStrings: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{countOccurrences: true, uniqStrings: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{countOccurrences: true, repeatedStrings: true, uniqStrings: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{countOccurrences: true, repeatedStrings: true},
		},
		{
			[]string{""},
			[]string{""},
			Options{},
		},
	}
	for _, test := range tests {
		require.Equal(t, test.out, Uniq(test.in, test.flags))
	}
}
