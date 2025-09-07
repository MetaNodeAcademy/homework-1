package main

import (
	"fmt"
)

func main() {
	var strs []string = []string{"flwoess", "flow3", "flowru"}
	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 0; i < len(strs[0]); i++ {
		allPrefix := prefix[0:(i + 1)]
		count := 0
		flag := false
		for index := range strs {

			if strs[index][0:i+1] == allPrefix {
				count++
			}
			if len(strs[index]) == (i + 1) {
				//达到最短字符串的边界时
				flag = true
			}
		}

		if len(strs) != count {
			if i != 0 {
				return prefix[:(i)]
			} else {
				return ""
			}

		} else {
			if flag {
				return allPrefix
			}

		}
	}
	return ""
}
