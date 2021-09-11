package stack

import (
	"strings"
)

// The string it self work as a stack
// I can access the top of it and pop as showed above
// Time Complexity: Big O(n)
// Space Complexity: Big O(n) -> directly proportional to n
func ReverseString(s string) string {
	var reverse []string

	for len(s) > 0 {
		n := len(s) - 1
		top := s[n]                            // Top Item
		reverse = append(reverse, string(top)) // byte to string

		s = s[:n] // Pop
	}

	return strings.Join(reverse, "")
}

// Time Complexity: O(n)
// Space Complexity: Big O(1)
// (In this case, still uses O(n), cause strings in go are imutable I still need to create a copy
// of the string (N size), how to do this?)
// TODO: Implement using Space Complexity of O(1)
func ReverseStringUsingSwap(s string) string {
	var reverse []byte

	i := 0
	j := len(s) - 1

	reverse = append(reverse, []byte(s)...)

	for i < j {
		reverse[i] = s[j]
		reverse[j] = s[i]
		i++
		j--
	}

	return string(reverse)
}
