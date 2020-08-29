package interpreter

import (
	"container/list"
	"log"
	"strconv"
	"strings"
)

type Expression interface {
	interpret() int64
}

type SubtractionExpression struct {
	exp1 Expression
	exp2 Expression
}

func NewSubtractionExpression(exp1 Expression, exp2 Expression) SubtractionExpression {
	return SubtractionExpression{
		exp1: exp1,
		exp2: exp2,
	}
}

func (exp SubtractionExpression) interpret() int64 {
	return exp.exp1.interpret() - exp.exp2.interpret()
}

type AddExpression struct {
	exp1 Expression
	exp2 Expression
}

func NewAddExpression(exp1 Expression, exp2 Expression) AddExpression {
	return AddExpression{
		exp1: exp1,
		exp2: exp2,
	}
}

func (exp AddExpression) interpret() int64 {
	return exp.exp1.interpret() + exp.exp2.interpret()
}

type NumberExpression struct {
	number int64
}

func NewNumberExpressionFromInt64(number int64) NumberExpression {
	return NumberExpression{
		number: number,
	}
}

func NewNumberExpressionFromString(number string) NumberExpression {
	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Fatal("Invalid expression for num", number)
	}
	return NumberExpression{
		number: num,
	}
}

func (exp NumberExpression) interpret() int64 {
	return exp.number
}

// Demo example how to use

type ExpressionInterpreter struct {
	nums *list.List
}

func NewExpressionInterpreter() ExpressionInterpreter {
	return ExpressionInterpreter{
		nums: list.New(),
	}
}

func (demoInterpreter ExpressionInterpreter) Interpret(str string) int64 {
	elements := strings.Fields(str)
	length := len(elements)
	nums := demoInterpreter.nums
	//1 2 3 + -
	for i := 0; i < (length+1)/2; i++ {
		nums.PushBack(NewNumberExpressionFromString(elements[i]))
	}
	for i := (length + 1) / 2; i < length; i++ {
		num1Ele := nums.Front()
		num2Ele := num1Ele.Next()
		num1Exp := num1Ele.Value.(NumberExpression)
		num2Exp := num2Ele.Value.(NumberExpression)
		nums.Remove(num1Ele)
		nums.Remove(num2Ele)
		var expression Expression
		switch elements[i] {
		case "+":
			expression = NewAddExpression(num1Exp, num2Exp)
		case "-":
			expression = NewSubtractionExpression(num1Exp, num2Exp)
		default:
			log.Fatal("Unrecoginzed expression: ", elements[i])
		}
		nums.PushFront(NewNumberExpressionFromInt64(expression.interpret()))
	}
	return nums.Front().Value.(NumberExpression).interpret()
}
