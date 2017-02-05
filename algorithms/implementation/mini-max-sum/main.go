package main

import "fmt"
import "sort"

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var nums int64arr

	nums = make(int64arr, 5)

	for i, _ := range nums {
		fmt.Scan(&nums[i])
	}

	min := int64(0)
	max := int64(0)

	sort.Sort(nums)

	for i := 0; i < 4; i++ {
		min += nums[i]
		max += nums[4-i]
	}

	fmt.Println(min, max)
}
