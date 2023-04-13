package tasks

import "sort"

// https://leetcode.com/problems/brace-expansion-ii/

// Expression /*Base interface for expressions*/
type Expression interface {
	Evaluate() []string
}

type CharExpression struct {
	Char byte
}

func (ce *CharExpression) Evaluate() []string {
	return []string{string(ce.Char)}
}

type BinaryExpression struct {
	Left, Right Expression
}

type OrExpression struct {
	BinaryExpression
}

func (oe *OrExpression) Evaluate() []string {
	set := make(map[string]struct{})
	for _, l := range oe.Left.Evaluate() {
		set[l] = struct{}{}
	}
	for _, r := range oe.Right.Evaluate() {
		set[r] = struct{}{}
	}

	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}

	return res
}

type AndExpression struct {
	BinaryExpression
}

func pop[T any](list *[]T) T {
	var res T
	res, *list = (*list)[len(*list)-1], (*list)[:len(*list)-1]
	return res
}

func (ae *AndExpression) Evaluate() []string {
	left := ae.Left.Evaluate()
	right := ae.Right.Evaluate()

	res := make([]string, 0, len(left)*len(right))
	for _, l := range left {
		for _, r := range right {
			res = append(res, l+r)
		}
	}

	return res
}

func parseExpressionForBraceExpansionII(expression string) Expression {
	postfixNotation := makePostfixNotation(expression)
	stack := make([]Expression, 0, len(postfixNotation))

	for _, token := range postfixNotation {
		if token != '*' && token != '+' {
			stack = append(stack, &CharExpression{Char: token})
			continue
		}

		var res Expression
		if token == '+' {
			res = &OrExpression{BinaryExpression{Right: pop(&stack), Left: pop(&stack)}}
		} else {
			res = &AndExpression{BinaryExpression{Right: pop(&stack), Left: pop(&stack)}}
		}

		stack = append(stack, res)
	}

	return stack[len(stack)-1]
}

func makePostfixNotation(expression string) []byte {
	res := make([]byte, 0, 2*len(expression))
	stack := make([]byte, 0, 2*len(expression))

	for i := range expression {
		c := expression[i]
		if c == '{' {
			stack = append(stack, c)
			continue
		}

		var op byte
		if c == '}' {
			for stack[len(stack)-1] != '{' {
				res = append(res, pop(&stack))
			}
			pop(&stack)

			if i+1 < len(expression) {
				nextC := expression[i+1]
				if nextC != ',' && nextC != '}' {
					op = '*'
				}
			}
		} else if c == ',' {
			op = '+'
		} else {
			res = append(res, c)
			if i+1 < len(expression) {
				nextC := expression[i+1]
				if nextC != ',' && nextC != '}' {
					op = '*'
				}
			}
		}

		if op != '+' && op != '*' {
			continue
		}

		for len(stack) > 0 &&
			(op == '+' && (stack[len(stack)-1] == '*' || stack[len(stack)-1] == '+') ||
				op == '*' && stack[len(stack)-1] == '*') {
			res = append(res, pop(&stack))
		}

		stack = append(stack, op)
	}

	for len(stack) > 0 {
		res = append(res, pop(&stack))
	}

	return res
}

func braceExpansionII(expression string) []string {
	exp := parseExpressionForBraceExpansionII(expression)
	res := exp.Evaluate()
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
