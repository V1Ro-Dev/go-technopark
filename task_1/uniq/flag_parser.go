package uniq

import (
	"errors"
	"flag"
)

type Options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

func GetParsedFlags() (Options, error) {
	cFlag := flag.Bool("c", false, "counts the number of occurrences of strings in the input data")
	dFlag := flag.Bool("d", false, "displays repeated strings in the input data")
	uFlag := flag.Bool("u", false, "displays task_1 strings in the input data")
	fFlag := flag.Int("f", 0, "doesn't count first num fields in the input data")
	sFlag := flag.Int("s", 0, "doesn't count first num chars in the input data")
	iFlag := flag.Bool("i", false, "ignores case of letters")

	flag.Parse()

	options := Options{
		c: *cFlag,
		d: *dFlag,
		u: *uFlag,
		f: *fFlag,
		s: *sFlag,
		i: *iFlag,
	}

	if options.c && options.d || options.c && options.u || options.u && options.d || options.d && options.i && options.c {
		err := errors.New("wrong flag sequence provided")
		return options, err
	}

	return options, nil
}
