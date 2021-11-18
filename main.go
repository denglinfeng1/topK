package main

import (
	"fmt"
)

func topK(s []int, k int) int {
	if len(s) < k {
		return 0
	}
	base := (len(s)-1+0)/2
	i, j := 0, len(s) - 1
	for {
		if s[i] > s[base] {
			s[i], s[base] = s[base], s[i]
		}
		i++
		if s[j] < s[base] {
			s[j], s[base] = s[base], s[j]
		}
		j--
		if i>=j {
			break
		}
	}
	if len(s[:i])-1 >= k-1 {
		return s[k-1]
	}
	return topK(s[i:], k-len(s[:i]))

}


func main()  {
	a := []int{1,3,2,5,4,8,7,10,6,9}
	fmt.Println(topK(a, 9))
}