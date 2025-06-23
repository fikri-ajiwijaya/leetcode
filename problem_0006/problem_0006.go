package problem_0006

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}

func convert(s string, numRows int) string {
	if numRows < 1 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	r := []rune(s)
	new := make([]rune, 0, len(r))
	for i := 0; i < numRows; i++ {
		for j := 0; ; j++ {
			{
				k := j*(2*numRows-2) + i
				if k >= len(r) {
					break
				}
				new = append(new, r[k])
			}
			if i > 0 && i < numRows-1 {
				k := (j+1)*(2*numRows-2) - i
				if k < len(r) {
					new = append(new, r[k])
				}
			}
		}
	}
	return string(new)
}
