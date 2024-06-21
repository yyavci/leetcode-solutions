package main

import "fmt"

func main() {
	res := twoSum([]int{1, 2, 3, 4}, 5)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for y := 0; y < len(nums); y++ {
			if i == y {
				continue
			}
			s := nums[i] + nums[y]
			if s == target {
				return []int{i, y}
			}
		}
	}
	return nil
}
