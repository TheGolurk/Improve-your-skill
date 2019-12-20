// Difficult: easy

package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 8, 16, 20}, 28))
}

func twoSum(nums []int, target int) []int {

	mapHash := make(map[int]int)

	for index := 0; index <= len(nums)-1; index++ {
		complement := target - nums[index]

		if _, ok := mapHash[complement]; ok {
			return []int{mapHash[complement], index}
		}

		mapHash[nums[index]] = index
	}
	panic("No hay solucion")
}
