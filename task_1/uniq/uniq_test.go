package uniq

import (
	"errors"
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
			"Flags: none",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{},
			nil,
		},
		{
			"Flags: 'c'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
			Options{countOccurrences: true},
			nil,
		},
		{
			"Flags: 'd'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
			Options{repeatedStrings: true},
			nil,
		},
		{
			"Flags: 'u'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"", "Thanks."},
			Options{uniqStrings: true},
			nil,
		},
		{
			"Flags: 'i'",
			[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik."},
			[]string{"I LOVE MUSIC.", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{caseInsensitive: true},
			nil,
		},
		{
			"Flags: 'f': 1",
			[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"We love music.", "", "I love music of Kartik.", "Thanks."},
			Options{skipFields: 1},
			nil,
		},
		{
			"Flags: 's': 1",
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			Options{skipChars: 1},
			nil,
		},
		{
			"Flags: 'c'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{""},
			Options{countOccurrences: true, repeatedStrings: true},
			errors.New("'c', 'd', 'u' flags should be used separately"),
		},
		{
			"Flags: 'd', 'u'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{""},
			Options{uniqStrings: true, repeatedStrings: true},
			errors.New("'c', 'd', 'u' flags should be used separately"),
		},
		{
			"Flags: 'c', 'u'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{""},
			Options{countOccurrences: true, uniqStrings: true},
			errors.New("'c', 'd', 'u' flags should be used separately"),
		},
		{
			"Flags: 'c', 'd', 'u'",
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{""},
			Options{countOccurrences: true, repeatedStrings: true, uniqStrings: true},
			errors.New("'c', 'd', 'u' flags should be used separately"),
		},
		{
			"Flags: none, input: none",
			[]string{""},
			[]string{""},
			Options{},
			nil,
		},
	}
	for _, test := range tests {
		actual, err := Uniq(test.in, test.flags)
		require.Equal(t, test.err, err)
		require.Equal(t, test.out, actual)
	}
}
