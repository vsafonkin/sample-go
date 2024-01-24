package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/vsafonkin/sample-go/gendata"
)

func main() {
	s := gendata.IntRandomSlice(100, 100)
	slices.Sort(s)
	fmt.Println(s)
	fmt.Println("--------")
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
