package stack

import (
	"fmt"
	"strconv"
)

/*
Infix Notation: <operand><operator><operand>

Binary operator: An operator that requires exact two operands
Follow the precendence rules
1. Parentheses
2. Exponents                   (right to left) - right associative
3. Division and Multiplication (left to right) - left  associative
4. Addition and Subtraction    (left to right) - left  associative
----------------------------------------------------------------------
Prefix Notation: <operator><operand><operand>
(Polish)

(Infix -> Prefix)
Example: 2 + 3     -> + 2 3
Example: P - Q 		 -> - P Q
Example: a + b * c -> + a * b c
----------------------------------------------------------------------
Postfix Notation: <operand><operand><operator>
(Reverse Polish Notation)

It's easiest to parse and cost less in terms of complexity and memory
Algorithms to parse is really straightforward and intuitive (it's prefered for computation using machines)

(Prefix -> Postfix)
Example: + 2 3     -> 2 3 +
Example: + P Q     -> P Q -
Example: + a * b c -> a b c * +
----------------------------------------------------------------------
To convert Infix to Postfix or Prefix, we need to go step by step
Example: a + b * c -> Prefix

1. Convert the part that should be evaluated first (order of precedence)
a + b * c   -> a + (* b c)
a + (* b c) -> + a (* b c)

2. We can use parenthesis in intermediate steps, and after, remove it
+ a (* b c) -> + a * b c

3. Parenthesis just adds readabilty to human, for machine, it doesn't matter
So we can remove to save some memory

Example: a + b * c -> Postfix
a + b * c   -> a + (b c *)
a + (b c *) -> a (b c *) +
a (b c *) + -> a b c * +
*/

/*
Algorithms to evaluate prefix and postfix are similar

### Infix: a * b + c * d - e
			1. {(a * b) + (c * d)} - e
				-> {(a b *) + (c d *)} - e
			2. {(a b *) + (c d *)} - e
				-> {(a b *)(c d *) +}  - e
			3. {(a b *)(c d *) +}  - e
				-> {(a b *)(c d *) +}  e -
			4. {(a b *)(c d *) +}  e -
				-> a b * c d * + e -
### Postfix: a b * c d * + e -

Evaluating: a b * c d * + e -
a = 2, b = 3, c = 5, d = 4, e = 9 -> 2 3 * 5 4 * + 9 -
1. Scan for left to right
2. find the fist ocurrence operator
3. for the first operator, the two preceding entities will always be operands
-> 6 5 4 * + 9 -
-> 6 20 + 9 -
-> 26 9 -
-> 17
*/
func isOperator(token string) bool {
	return token == "*" || token == "/" || token == "+" || token == "-"
}

func performOperation(operation string, operand1 string, operand2 string) (string, error) {
	op1, _ := strconv.Atoi(operand1)
	op2, _ := strconv.Atoi(operand2)

	if operation == "*" {
		return fmt.Sprint(op1 * op2), nil
	} else if operation == "/" {
		return fmt.Sprint(op1 / op2), nil
	} else if operation == "+" {
		return fmt.Sprint(op1 + op2), nil
	} else if operation == "-" {
		return fmt.Sprint(op1 - op2), nil
	}

	return "", fmt.Errorf("operation not recognized")
}

// TODO: Handle multiple digits numbers (https://gist.github.com/mycodeschool/7702441)
func PostfixToInt(e string) int {
	var operands Stack

	for _, c := range e {
		token := string(c)

		if token == " " {
			continue
		}

		if isOperator(token) {
			operand2, _ := operands.Pop()
			operand1, _ := operands.Pop()
			res, _ := performOperation(token, (operand1).(string), (operand2).(string))
			operands.Push(res)
		} else {
			operands.Push(token)
		}
	}

	res, _ := strconv.Atoi((operands.Top()).(string))
	return res
}

// The same, but scanning from right to left
func PrefixToInt(e string) int {
	var operands Stack

	for i := len(e) - 1; i >= 0; i-- {
		token := string(e[i])
		if token == " " {
			continue
		}

		if isOperator(token) {
			// The first operand os the first popped
			operand1, _ := operands.Pop()
			operand2, _ := operands.Pop()
			res, _ := performOperation(token, (operand1).(string), (operand2).(string))
			operands.Push(res)
		} else {
			operands.Push(token)
		}
	}

	res, _ := strconv.Atoi((operands.Top()).(string))
	return res
}
