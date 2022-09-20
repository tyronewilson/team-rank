package util

func SplitOnLastSpace(str string) (string, string) {
	if len(str) == 0 {
		return "", ""
	}
	ptr := len(str) - 1
	for ptr >= 0 && str[ptr] != ' ' {
		ptr--
	}
	if ptr == -1 {
		return "", ""
	}
	return str[:ptr], str[ptr+1:]
}
