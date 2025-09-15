package main

import "fmt"

func main() {
	nums := []int{0, 3, -3, 4}
	target := 8
	rst := twoSum(nums, target)
	fmt.Println(rst)
}

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	rst := []int{}
	for i := 0; i < len(nums); i++ {
		intparam := target - nums[i]
		value, ok := mp[intparam]

		if ok {
			if nums[value]+nums[i] == target {
				rst = append(rst, value, i)
				break
			}
		} else {
			mp[intparam] = i
			mp[nums[i]] = i
		}

	}
	return rst
}
