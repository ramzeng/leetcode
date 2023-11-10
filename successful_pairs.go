package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	var answers []int
	var answer int
	baseline := success / int64(potions[len(potions)-1])

	for _, spell := range spells {
		answer = 0

		if int64(spell) < baseline {
			answers = append(answers, answer)
			continue
		}

		index := sort.Search(len(potions), func(i int) bool {
			return int64(potions[i]*spell) >= success
		})

		answers = append(answers, len(potions)-index)
	}

	return answers
}
