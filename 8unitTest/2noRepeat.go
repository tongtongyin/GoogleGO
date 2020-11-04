package main
var lastOccurred = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int {

	//记录每个字符最后出现的位置
	// stores last occured pos +1.
	// 0 means not seen
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI > start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

