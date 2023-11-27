package retos

import (
	"errors"
	"fmt"
	"math"
)

/**
Writes a function in Go that finds and returns
the largest element in a slice of integers.

Input:

nums []int: A slice of integers. The slice can
contain any number of elements, and the elements
can be any integer value (positive, negative, or zero)

Output:
The function should return two values:
	- The largest integer in the slice.
	- An error object, which should be non-nil if the
		slice is empty.

Example 1:
	Input: nums = [3, 5, 4, 4, 6, 4, 9, 5, 2]
	Output: 9, nil

Example 2:
	Input: nums = [-2, -5, -4, -4, -6, -4, -9, 7, -4]
	Output: 7, nil

Example 3:
	Input: numbs = []
	Output: 0, error
*/

// FindLargestBasic Solución poco eficiente al problema.
// Esta solución tiene una complejidad algorítmica es de O(n^2).
func FindLargestBasic(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

loopI:
	for i := 0; i < len(nums); i++ {
		numInTurn := nums[i]
		for j := 0; j < len(nums); j++ {
			if numInTurn < nums[j] {
				continue loopI
			}
		}
		return numInTurn, nil
	}
	return 0, nil
}

// FindLargest Solución al problema con una complejidad algorítmica O(n).
func FindLargest(nums []int) (int, error) {
	length := len(nums)
	if length == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max, nil
}

// FindLargestPro Solución más eficiente al problema.
// La complejidad algorítmica es de O(n/2).
func FindLargestPro(nums []int) (int, error) {
	length := len(nums)
	if length == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	var value int
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		r := nums[i] - nums[j]
		if r >= 0 && nums[i] >= value {
			value = nums[i]
			continue
		}
		if r < 0 && nums[j] >= value {
			value = nums[j]
			continue
		}
	}
	if middle := math.Floor(float64(length / 2)); length%2 > 0 && nums[int(middle)] > value {
		value = nums[int(middle)]
	}
	return value, nil
}
