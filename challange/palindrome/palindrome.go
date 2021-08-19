package palindrome

import (
	"fmt"
	"strconv"
)

func Run() {
	fmt.Println("start")

	res := Solution(123)
	fmt.Println(res)

	fmt.Println("end")
}

func Solution(x int) bool {
	str := strconv.Itoa(x)
	j := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}
