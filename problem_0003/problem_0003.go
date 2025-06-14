package problem_0003

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	max_length := 0
	pos := map[rune]int{}
	beg := 0
	for i, r := range s {
		loc, found := pos[r]
		if found && loc >= beg {
			beg = loc + 1
		}
		pos[r] = i
		length := i + 1 - beg
		if length > max_length {
			max_length = length
		}
	}
	return max_length
}
