package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		rest := append([]int{}, nums[:i]...)
		rest = append(rest, nums[i+1:]...)
		for _, perm := range permute(rest) {
			result = append(result, append([]int{num}, perm...))
		}
	}

	return result
}

func main() {

	fmt.Println(permute([]int{4, 5, 6}))

}
