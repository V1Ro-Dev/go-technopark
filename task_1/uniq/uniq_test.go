package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccessfulCases(t *testing.T) {
	var tests = []struct {
		in    []string
		out   []string
		flags Options
	}{
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
			Options{c: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
			Options{d: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"", "Thanks."},
			Options{u: true},
		},
		{
			[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik."},
			[]string{"I LOVE MUSIC.", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
			Options{i: true},
		},
		{
			[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"We love music.", "", "I love music of Kartik.", "Thanks."},
			Options{f: 1},
		},
		{
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			Options{s: 1},
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
			Options{c: true, d: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{u: true, d: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{c: true, u: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{c: true, d: true, u: true},
		},
		{
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"wrong flag sequence provided"},
			Options{c: true, d: true},
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
