/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	diff := 0
	minDiff := math.MaxInt
	answer := 0

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff = abs(sum - target)

			if diff < minDiff {
				answer = sum
				minDiff = diff
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

