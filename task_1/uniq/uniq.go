package uniq

import (
	"strconv"
	"strings"
)

func skipFields(str string, opt Options) string {
	fields := strings.Split(str, " ")
	if len(fields) < opt.f {
		return ""
	}
	return strings.Join(fields[opt.f:], " ")
}

func skipChars(str string, opt Options) string {
	if len(str) < opt.s {
		return ""
	}
	return str[opt.s:]
}

func updateAns(res []string, opt Options, occurrences int, str string) []string {
	if opt.c {
		res = append(res, strconv.Itoa(occurrences)+" "+str)
	}
	if opt.u && occurrences == 1 || opt.d && occurrences > 1 || !opt.d && !opt.c && !opt.u {
		res = append(res, str)
	}
	return res
}

func isValidOptions(options Options) bool {
	if options.c && options.d || options.c && options.u || options.u && options.d || options.d && options.i && options.c {
		return false
	}
	return true
}

func Uniq(stringSlc []string, opt Options) []string {
	if !isValidOptions(opt) {
		return []string{"wrong flag sequence provided"}
	}
	res := []string{}
	occurrences := 1
	prev := skipChars(skipFields(stringSlc[0], opt), opt) // сначала скипаем поля, потом символы
	prevPos := 0
	for i := 1; i < len(stringSlc); i++ {
		now := skipChars(skipFields(stringSlc[i], opt), opt)

		if opt.i {
			prev, now = strings.ToLower(prev), strings.ToLower(now)
		}

		if now == prev {
			occurrences++
		} else {
			res = updateAns(res, opt, occurrences, stringSlc[prevPos])
			prev = now
			prevPos = i
			occurrences = 1
		}
	}
	res = updateAns(res, opt, occurrences, stringSlc[prevPos])
	return res
}
