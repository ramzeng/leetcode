package main

// f(n+1)=f(n)+f(n−1)
// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof
// 动态规划 O(n)
func fib(n int) int {
	const mod int = 1e9 + 7

	if n < 2 {
		return n
	}

	sum, a, b := 1, 0, 0

	for i := 2; i <= n; i++ {
		b = a
		a = sum
		sum = (a + b) % mod
	}

	return sum
}
