package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("-----")
	s := []int{1, 3, 5, 6, 9, 12, 13, 17, 24, 87, 99}
	if len(os.Args) > 1 {
		target, err := strconv.Atoi(os.Args[1])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(searchInsert(s, target))
	}
}

func searchInsert(nums []int, target int) int {
	i := len(nums) / 2
	switch {
	case i == 0:
		if nums[0] < target {
			return 1
		}
		return 0
	case nums[i] > target:
		out := searchInsert(nums[:i], target)
		return out
	case nums[i] < target:
		return searchInsert(nums[i:], target) + i
	}
	return i
}
