package main

import (
	"fmt"
	"strconv"
)

func main() {
	digits := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	//转成int值
	s := ""
	for _, value := range digits {
		s += strconv.Itoa(value)
	}
	//int值加一
	int1, _ := strconv.ParseInt(s, 10, 64)
	//转为字符串后变数组
	afts := strconv.FormatInt(int1+1, 10)

	var intls []int = []int{}
	for i := 0; i < len(afts); i++ {
		intsp, _ := strconv.Atoi(afts[i:(i + 1)])
		intls = append(intls, intsp)
	}
	return intls

}
