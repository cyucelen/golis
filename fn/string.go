package fn

import "strings"

func TrimEach(cutset string) func([]string) []string {
	return func(ss []string) []string {
		return MapString(ss, func(s string) string {
			return strings.Trim(s, cutset)
		})
	}
}

var TrimCommasFromEach = TrimEach(",")

var TrimSpacesFromEach = TrimEach(" ")

func IsStartsWith(s string, c byte) bool {
	return s[0] == c
}

func IsEndsWith(s string, c byte) bool {
	return s[len(s)-1] == c
}

func IsSingleChar(s string) bool {
	return len(s) == 1
}
