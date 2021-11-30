/*
Question:
Detect palindrome products in a given range.

A palindromic number is a number that remains the same when its digits are reversed. For example, 121 is a palindromic number but 112 is not.

Given a range of numbers, find the largest and smallest palindromes which are products of two numbers within that range.

Your solution should return the largest and smallest palindromes, along with the factors of each within the range. If the largest or smallest palindrome has more than one pair of factors within the range, then return all the pairs.

Solution:

1. Sort the array, then insert the values of the array into a channel (raw_chan).
2. Create a function that determines if an int is a palindrome
3. Create a function that adds two elements of the raw_chan into a sum_chan.
4. Create two structs of the same type, smallest_pal, largest_pal.
	Struct will have value, slice of slices (all factorials)
5. Create worker pool A: Adds 2 values to the sum_chan.
6. Create worker pool B: Checks to see if value from sum_chan is a palindrome.
	If it is, is it smaller than the smallest_pal or larger than the largest_pal?
	Keep track of the factors

Should I have a struct keeping track?

*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	if len(numStr) == 1 {
		return true

	}

	var firstHalf string
	var secondHalf string

	if len(numStr)%2 == 0 {
		firstHalf = numStr[:(len(numStr) / 2)]
		secondHalf = numStr[(len(numStr) / 2):]
	} else {
		firstHalf = numStr[:(len(numStr) / 2)]
		secondHalf = numStr[(len(numStr)/2)+1:]
	}

	reverseSecondHalf := ""
	for i := (len(secondHalf) - 1); i >= 0; i-- {
		reverseSecondHalf += string(secondHalf[i])
	}

	if reverseSecondHalf == firstHalf {
		return true
	}

	return false
}

type Palindrome struct {
	Value     int
	Factorial [][]int
}

func (pal *Palindrome) update(val int, fact []int) {
	if pal.Value == val {
		pal.Factorial = append(pal.Factorial, fact)
	} else {
		pal.Value = val
		pal.Factorial = nil
		pal.Factorial = [][]int{fact}
	}

}

// create a clean method that trims duplicates in the factorial

func main() {
	fmt.Println(isPalindrome(10101))

	inputVal := make([]int, 2)
	for i := 10; i <= 99; i++ {
		inputVal = append(inputVal, i)
	}
	sort.Ints(inputVal)
	maxPalindromeValue := Palindrome{Value: 0}
	maxPalindromeValue.Factorial = [][]int{}

	minPalindromeValue := Palindrome{Value: (inputVal[len(inputVal)-1] * inputVal[len(inputVal)-1])}
	minPalindromeValue.Factorial = [][]int{}

	for first := 0; first < len(inputVal); first++ {
		for second := 0; second < len(inputVal); second++ {
			cur := inputVal[first] * inputVal[second]
			cur_fact := []int{inputVal[first], inputVal[second]}
			if isPalindrome(cur) {
				if cur <= minPalindromeValue.Value && cur != 0 {
					minPalindromeValue.update(cur, cur_fact)
				} else if cur >= maxPalindromeValue.Value {
					maxPalindromeValue.update(cur, cur_fact)
				}
			}
		}
	}

	fmt.Println(minPalindromeValue)
	fmt.Println(maxPalindromeValue)

}
