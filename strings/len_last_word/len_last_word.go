package lenlastword

/*
	input: string
	output: int of len of last word
	con: don't count whitespaces

	find the index for final char of last word
	find the index for first char of last word
	final char index - first char index = last word len
*/

func lenLastWord(s string) int {
	//Find index for final char
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// copy final char index
	end := i
	// find the first char of the last word
	for i > 0 && s[i] != ' ' {
		i--
	}
	// final char index - first char index = last word length
	return end - i
}
