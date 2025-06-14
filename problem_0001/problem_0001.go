package problem_0001

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	complement := map[int]int{}
	for i, n := range nums {
		j, exists := complement[n]
		if exists {
			return []int{j, i}
		}
		complement[target-n] = i
	}
	return nil
}
