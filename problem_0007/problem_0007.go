package problem_0007

func Reverse(x int) int {
	return reverse(x)
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	for x%10 == 0 {
		x /= 10
	}
	y := 0
	for x != 0 {
		if y >= 214748364 || y <= -214748364 {
			if !((y == 214748364 && x%10 <= 7) || (y == -214748364 && x%10 >= -8)) {
				return 0
			}
		}
		y = y*10 + x%10
		x = x / 10
	}
	return y
}
