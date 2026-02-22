package main

import "fmt"

//Example 1:
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int)   // A mapping to store numbers and their indices
	for i, n := range nums { // i = 0, n = 2; i = 1, n = 7; i = 2, n = 11;
		complement := target - n                  // Find the required number to reach the target t = 9, n = 2; So, c = 7; t = 9, n = 7, c = 2;
		if index, found := m[complement]; found { // m[7] = not found, condition not met; m[2] = not found;
			return []int{index, i}
		}

		m[n] = i // m[7] = 0; m[2] = 1;
	}

	return nil
}
