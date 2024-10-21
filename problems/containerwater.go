package main

import (
	"fmt"
	"math"
	"sort"
)

func calcualte(a []int) int {

	left := 0
	right := len(a) - 1

	mx := 0
	for left < right {
		height := int(math.Min(float64(a[left]), float64(a[right])))
		width := right - left
		area := height * width

		mx = int(math.Max(float64(mx), float64(area)))

		if a[left] < a[right] {
			left++
		} else {
			right--
		}
	}

	return mx
}

func main() {

	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//fmt.Println(calcualte(a))

	sort.Ints(a)
	fmt.Println(a)
}
