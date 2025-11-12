package letterCombo

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	backtrack(&res, digits, "", 0)
	return res
}

var letterMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func backtrack(res *[]string, digits, current string, index int) {
	if index == len(digits) {
		*res = append(*res, current)
		return
	}
	digit := digits[index]
	letters := letterMap[digit]
	for _, letter := range letters {
		backtrack(res, digits, current+letter, index+1)
	}
}
