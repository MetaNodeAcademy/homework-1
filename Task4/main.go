package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 2}
	i := removeDuplicates(nums1)
	fmt.Println(i)
}
func removeDuplicates(nums []int) int {
	// rst := []int{}
	// rst = append(rst, nums[0])
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		//两个值相同的加一
	// 	} else {
	// 		rst = append(rst, nums[i+1])
	// 	}
	// }
	// fmt.Println(rst)
	// return len(rst)

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {

		} else {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums)
	return j
}
