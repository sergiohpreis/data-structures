package stack

// Should check if an expression is balanced
// True or False
// Examples:
// (a+b) 					 -> true
// {(a+b) + (a+b)} -> true
// {(x+y) * (z) 	 -> false
// [2*3] + (a)] 	 -> false
// {a_z) 					 -> false

func IsBalanced(expression string) bool {
	var stack Stack

	for _, c := range expression {
		char := string(c)

		if char == "(" || char == "[" || char == "{" {
			stack.Push(char)
		} else if stack.IsEmpty() {
			return false
		} else if char == ")" && stack.Top() != "(" {
			return false
		} else if char == "]" && stack.Top() != "[" {
			return false
		} else if char == "}" && stack.Top() != "{" {
			return false
		} else {
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}
