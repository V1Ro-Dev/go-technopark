package uniq

import (
	"errors"
	"fmt"
	"strings"
)

func skipFields(str string, opt Options) string {

	fields := strings.Split(str, " ")

	if len(fields) < opt.skipFields {
		return ""
	}

	return strings.Join(fields[opt.skipFields:], " ")
}

func skipChars(str string, opt Options) string {

	if len(str) < opt.skipChars {
		return ""
	}

	if opt.caseInsensitive {
		str = strings.ToLower(str)
	}

	return str[opt.skipChars:]
}

func updateAns(res []string, opt Options, occurrences int, str string) []string {

	switch {
	case opt.uniqStrings && occurrences == 1,
		opt.repeatedStrings && occurrences > 1,
		!opt.repeatedStrings && !opt.countOccurrences && !opt.uniqStrings:

		res = append(res, str)

	case opt.countOccurrences:

		res = append(res, fmt.Sprintf("%d %s", occurrences, str))
	}

	return res
}

func isValidOptions(options Options) bool {

	switch {
	case options.countOccurrences && options.repeatedStrings,
		options.countOccurrences && options.uniqStrings,
		options.uniqStrings && options.repeatedStrings,
		options.repeatedStrings && options.uniqStrings && options.caseInsensitive:

		return false
	}

	return true
}

func Uniq(stringSlc []string, opt Options) ([]string, error) {

	if !isValidOptions(opt) {
		return []string{""}, errors.New("'c', 'd', 'u' flags should be used separately")
	}

	var (
		res         []string
		occurrences = 1
		prev        = skipChars(skipFields(stringSlc[0], opt), opt)
		prevPos     = 0
	)

	for i, str := range stringSlc {
		if i == 0 {
			continue
		}

		now := skipChars(skipFields(str, opt), opt)

		if now == prev {
			occurrences++
			continue
		}

		res = updateAns(res, opt, occurrences, stringSlc[prevPos])
		prev = now
		prevPos = i
		occurrences = 1
	}

	res = updateAns(res, opt, occurrences, stringSlc[prevPos])

	return res, nil
}
