/*
 * @lc app=leetcode.cn id=2300 lang=golang
 *
 * [2300] 咒语和药水的成功对数
 */

// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	n := len(spells)

	answers := make([]int, n)

	for i := 0; i < n; i++ {
		index := binarySearch(potions, int(success-1)/spells[i]+1)

		if index == len(potions) {
			answers[i] = 0
		} else {
			answers[i] = len(potions)-index
		}
	}

	return answers
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left+(right-left)/2

		if nums[mid] >= target {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return left
}
// @lc code=end

