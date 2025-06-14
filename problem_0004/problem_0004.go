package problem_0004

// import "slices"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	nums := append(nums1, nums2...)
// 	slices.Sort(nums)
// 	n := len(nums)
// 	if n%2 == 0 {
// 		return float64(nums[n/2-1]+nums[n/2]) / 2.0
// 	} else {
// 		return float64(nums[(n-1)/2])
// 	}
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	var loc int
	if n%2 == 0 {
		loc = n / 2
	} else {
		loc = (n - 1) / 2
	}
	nums := make([]int, loc+1)
	i := 0
	j := 0
	for k := 0; k <= loc; k++ {
		if j >= len(nums2) || i < len(nums1) && nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}
	if n%2 == 0 {
		return float64(nums[loc-1]+nums[loc]) / 2.0
	} else {
		return float64(nums[loc])
	}
}
