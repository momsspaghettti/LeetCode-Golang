package tasks

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/expression-add-operators/

type iExpressionNode interface {
	eval() int
}

type numericNode struct {
	value int
}

func (node *numericNode) eval() int {
	return node.value
}

type iBinaryOperationNode interface {
	setRight(node iExpressionNode)
}

type binaryOperationNode struct {
	leftValue, rightValue iExpressionNode
}

func (binOpNode *binaryOperationNode) setRight(node iExpressionNode) {
	if binOpNode.rightValue == nil {
		binOpNode.rightValue = node
	} else {
		(binOpNode.rightValue).(iBinaryOperationNode).setRight(node)
	}
}

type additionNode struct {
	*binaryOperationNode
}

func (node *additionNode) eval() int {
	return node.leftValue.eval() + node.rightValue.eval()
}

type subtractionNode struct {
	*binaryOperationNode
}

func (node *subtractionNode) eval() int {
	return node.leftValue.eval() - node.rightValue.eval()
}

type multiplicationNode struct {
	*binaryOperationNode
}

func (node *multiplicationNode) eval() int {
	return node.leftValue.eval() * node.rightValue.eval()
}

type stack []iExpressionNode

func (s *stack) push(node iExpressionNode) {
	*s = append(*s, node)
}

func (s *stack) pop() iExpressionNode {
	var node iExpressionNode
	node, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return node
}

func (s *stack) length() int {
	return len(*s)
}

func parseExpressionForAddOperators(expressionStr string) (iExpressionNode, error) {
	stack := stack(make([]iExpressionNode, 0, len(expressionStr)))

	buff := strings.Builder{}
	leftMost := true
	for i := 0; i < len(expressionStr); i++ {
		if expressionStr[i] == ' ' {
			continue
		}

		c := expressionStr[i]
		if c != '+' && c != '-' && c != '*' {
			buff.WriteByte(c)
			continue
		}

		value, err := strconv.Atoi(buff.String())
		if err != nil {
			return nil, err
		}

		buff.Reset()

		operationNodeInner := &binaryOperationNode{}
		var operationNode iExpressionNode
		switch c {
		case '+':
			operationNode = &additionNode{operationNodeInner}
		case '-':
			operationNode = &subtractionNode{operationNodeInner}
		case '*':
			operationNode = &multiplicationNode{operationNodeInner}
		default:
			return nil, fmt.Errorf("invalid operation symbol: %s", string(c))
		}

		valueNode := &numericNode{value}
		if leftMost {
			operationNodeInner.leftValue = valueNode
			stack.push(operationNode)
			leftMost = false
		} else {
			leftNode := stack.pop()
			binOpLeftNode, ok := (leftNode).(iBinaryOperationNode)
			if !ok {
				return nil, fmt.Errorf("cast to iBinaryOperationNode failed")
			}
			if c == '*' {
				operationNodeInner.leftValue = valueNode
				binOpLeftNode.setRight(operationNode)
				stack.push(leftNode)
			} else {
				binOpLeftNode.setRight(valueNode)
				operationNodeInner.leftValue = leftNode
				stack.push(operationNode)
			}
		}
	}

	if stack.length() > 2 {
		return nil, fmt.Errorf("stack.length() > 2")
	}

	s := buff.String()
	if len(s) > 0 {
		value, err := strconv.Atoi(buff.String())
		if err != nil {
			return nil, err
		}

		valueNode := &numericNode{value}
		var node iExpressionNode
		if stack.length() > 0 {
			node = stack.pop()
			binOpNode, ok := (node).(iBinaryOperationNode)
			if !ok {
				return nil, fmt.Errorf("cast to iBinaryOperationNode failed")
			}
			binOpNode.setRight(valueNode)
		} else {
			node = valueNode
		}

		stack.push(node)
	}

	if stack.length() != 1 {
		return nil, fmt.Errorf("stack.length() != 1")
	}

	return stack.pop(), nil
}

func splitToStringsArray(str string) []string {
	res := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		res = append(res, str[i:i+1])
	}
	return res
}

type item struct {
	expr    []string
	nextInd int
}

func strArrayToStr(strArray []string) string {
	return strings.Join(strArray, "")
}

func AddOperators(num string, target int) []string {
	l := len(num)
	n := l * l * l * l

	stack := make([]item, 0, n)
	res := make([]string, 0, n)

	numArr := splitToStringsArray(num)

	stack = append(stack, item{[]string{numArr[0]}, 1})

	for len(stack) > 0 {
		var i item
		i, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if i.nextInd == len(numArr) {
			expressionStr := strArrayToStr(i.expr)
			expression, err := parseExpressionForAddOperators(expressionStr)
			if err != nil {
				panic(err)
			}

			if expression.eval() == target {
				res = append(res, expressionStr)
			}
		} else {
			var expr []string
			if i.expr[len(i.expr)-1][0] != '0' {
				expr = make([]string, 0, len(i.expr))
				expr = append(expr, i.expr...)
				expr[len(expr)-1] += numArr[i.nextInd]
				stack = append(stack, item{expr, i.nextInd + 1})
			}

			expr = make([]string, 0, len(i.expr)+2)
			expr = append(expr, i.expr...)
			expr = append(expr, "+")
			expr = append(expr, numArr[i.nextInd])
			stack = append(stack, item{expr, i.nextInd + 1})

			expr = make([]string, 0, len(i.expr)+2)
			expr = append(expr, i.expr...)
			expr = append(expr, "-")
			expr = append(expr, numArr[i.nextInd])
			stack = append(stack, item{expr, i.nextInd + 1})

			expr = make([]string, 0, len(i.expr)+2)
			expr = append(expr, i.expr...)
			expr = append(expr, "*")
			expr = append(expr, numArr[i.nextInd])
			stack = append(stack, item{expr, i.nextInd + 1})
		}
	}

	return res
}
