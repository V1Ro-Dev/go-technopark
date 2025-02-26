package uniq

import (
	"errors"
	"flag"
)

type Options struct {
	countOccurrences bool
	repeatedStrings  bool
	uniqStrings      bool
	skipFields       int
	skipChars        int
	caseInsensitive  bool
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
		countOccurrences: *cFlag,
		repeatedStrings:  *dFlag,
		uniqStrings:      *uFlag,
		skipFields:       *fFlag,
		skipChars:        *sFlag,
		caseInsensitive:  *iFlag,
	}

	switch {
	case options.countOccurrences && options.repeatedStrings,
		options.countOccurrences && options.uniqStrings,
		options.uniqStrings && options.repeatedStrings,
		options.repeatedStrings && options.uniqStrings && options.caseInsensitive:

		return options, errors.New("wrong flag sequence provided")
	}

	return options, nil
}
