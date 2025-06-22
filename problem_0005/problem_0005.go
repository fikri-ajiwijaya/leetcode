package problem_0005

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func get_palindrome_length_at(r []rune, i int) int {
	length_odd := 1
	for j, k := i-1, i+1; j >= 0 && k < len(r); j, k = j-1, k+1 {
		if r[j] != r[k] {
			break
		}
		length_odd += 2
	}
	length_even := 0
	for j, k := i-1, i; j >= 0 && k < len(r); j, k = j-1, k+1 {
		if r[j] != r[k] {
			break
		}
		length_even += 2
	}
	if length_odd > length_even {
		return length_odd
	} else {
		return length_even
	}
}

func get_palindrome_at(r []rune, i int, length int) string {
	if length%2 == 0 {
		return string(r[i-length/2 : i+length/2])
	} else {
		return string(r[i-length/2 : i+length/2+1])
	}
}

func longestPalindrome(s string) string {
	longest := 0
	pos := -1
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		length := get_palindrome_length_at(r, i)
		if length > longest {
			longest = length
			pos = i
		}
	}
	return get_palindrome_at(r, pos, longest)
}
