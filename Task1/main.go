package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "{}[]()"
	var s2 string = "{}[()]"
	var s3 string = "{}[(])"
	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
}
func isValid(s string) bool {
	for {
		if strings.Contains(s, "{}") {
			s = strings.Replace(s, "{}", "", -1)
		} else if strings.Contains(s, "[]") {
			s = strings.Replace(s, "[]", "", -1)
		} else if strings.Contains(s, "()") {
			s = strings.Replace(s, "()", "", -1)
		} else {
			return len(s) == 0
		}
	}

}
