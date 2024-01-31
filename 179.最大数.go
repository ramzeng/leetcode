/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Slice(strs, func(current, next int) bool {
		return strs[current]+strs[next] > strs[next]+strs[current]
	})

	// var quickSort func(nums []string, left, right int)
	// quickSort = func(nums []string, left, right int) {
	// 	if left >= right {
	// 		return
	// 	}

	// 	mid := nums[left+(right-left)/2]
	// 	i, j := left, right

	// 	for i <= j {
	// 		for nums[i]+mid > mid+nums[i] {
	// 			i++
	// 		}

	// 		for nums[j]+mid < mid+nums[j] {
	// 			j--
	// 		}

	// 		if i <= j {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 			i++
	// 			j--
	// 		}
	// 	}

	// 	quickSort(nums, left, j)
	// 	quickSort(nums, i, right)
	// }
	// quickSort(strs, 0, len(strs)-1)

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}

// @lc code=end

