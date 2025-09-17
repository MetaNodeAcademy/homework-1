package main

import "fmt"

func main() {
	par := 10
	add(&par)
	fmt.Println(par)

	a := 5
	b := 6
	c := 7
	inls := []*int{&a, &b, &c}
	fmt.Println(inls)
	multls(inls)
	fmt.Println(*inls[0], *inls[1], *inls[2])
	fmt.Println(a, b, c)

	inls2 := []int{2, 3, 4}
	fmt.Println(inls2)
	multls2(&inls2)
	fmt.Println(inls2)
}

func add(intPara *int) {
	*intPara += 10
}

func multls(nums []*int) {
	for i := 0; i < len(nums); i++ {
		*nums[i] = *nums[i] * 2
	}
}

func multls2(nums *[]int) {
	numsls := *nums
	for i := 0; i < len(numsls); i++ {
		numsls[i] = numsls[i] * 2
	}
}
