package main

import (
	"fmt"
)

func calculateWeights(s string) map[int]bool {
	weights := make(map[int]bool)
	n := len(s)
	for i := 0; i < n; {
		char := s[i]
		count := 0
		for i+count < n && s[i+count] == char {
			count++
		}
		charWeight := int(char - 'a' + 1)
		for j := 1; j <= count; j++ {
			weights[charWeight*j] = true
		}
		i += count
	}
	return weights
}

func processQueries(s string, queries []int) []string {
	weights := calculateWeights(s)
	results := make([]string, len(queries))
	for i, q := range queries {
		if weights[q] {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}
	return results
}

func isBalanced(s string) string {
	// Stack to keep track of opening brackets
	var stack []rune

	// Map to hold matching pairs of brackets
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// Push opening brackets onto the stack
			stack = append(stack, char)
		case ')', '}', ']':
			// Check if stack is empty or top of stack is not the matching opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If stack is empty, all brackets were matched
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func highestPalindrome(s string, k int) string {
	if k < 0 {
		return "-1"
	}

	n := len(s)
	bytes := []byte(s)

	// Recursive function to create the highest palindrome
	var makePalindrome func(l, r, k int) bool
	makePalindrome = func(l, r, k int) bool {
		if l >= r {
			return true
		}

		if bytes[l] != bytes[r] {
			if k == 0 {
				return false
			}
			if bytes[l] > bytes[r] {
				bytes[r] = bytes[l]
			} else {
				bytes[l] = bytes[r]
			}
			k--
		}

		return makePalindrome(l+1, r-1, k)
	}

	// Step 1: Make the string a palindrome with minimal changes
	if !makePalindrome(0, n-1, k) {
		return "-1"
	}

	// Step 2: Maximize the palindrome with remaining changes
	var maximizePalindrome func(l, r, k int)
	maximizePalindrome = func(l, r, k int) {
		if l >= r || k <= 0 {
			return
		}

		if bytes[l] == bytes[r] {
			if bytes[l] != '9' && k >= 2 {
				bytes[l], bytes[r] = '9', '9'
				maximizePalindrome(l+1, r-1, k-2)
			} else {
				maximizePalindrome(l+1, r-1, k)
			}
		} else {
			if bytes[l] != '9' && bytes[r] != '9' && k >= 1 {
				bytes[l], bytes[r] = '9', '9'
				maximizePalindrome(l+1, r-1, k-1)
			} else {
				maximizePalindrome(l+1, r-1, k)
			}
		}
	}

	maximizePalindrome(0, n-1, k)
	return string(bytes)
}

func main() {
	exampleString := "abbcccd"
	queries := []int{1, 3, 9, 8}
	results := processQueries(exampleString, queries)
	for _, result := range results {
		fmt.Println(result)
	}

	exampleInputs := []string{
		"{[()]}",
		"{[(])}",
		"{(([])[[]])[]}",
	}

	for _, input := range exampleInputs {
		fmt.Printf("Input: %s, Output: %s\n", input, isBalanced(input))
	}

	input1 := "3943"
	k1 := 1
	fmt.Printf("Input: %s, k: %d, Output: %s\n", input1, k1, highestPalindrome(input1, k1))

	input2 := "932239"
	k2 := 2
	fmt.Printf("Input: %s, k: %d, Output: %s\n", input2, k2, highestPalindrome(input2, k2))
}
