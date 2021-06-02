package shuntingyard

import (
	"fmt"
	"io"
	"reflect"
)

type OperatorAssociative uint8

const (
	Assoc_none OperatorAssociative = iota
	Assoc_ltr
	Assoc_rtl
)

type OperatorDetails struct {
	Precedence int
	Assoc      OperatorAssociative
}

func ToPostFix(input io.Reader, operators map[byte]OperatorDetails) {
	content := make([]byte, 0)
	opStack := make([]byte, 0)
	if operators == nil {
		operators = getDefaultOperatorMap()
	}
	for {
		err := parseTokens(input, &opStack, &content, operators)
		if err == io.EOF {
			break
		}
	}
	if len(opStack) != 0 {
		reverseSlice(opStack)
		content = append(content, opStack...)
	}
	fmt.Println(string(content))

}

func getDefaultOperatorMap() map[byte]OperatorDetails {
	operators := make(map[byte]OperatorDetails)
	LBracketOp := OperatorDetails{Precedence: 15, Assoc: Assoc_ltr}
	RBracketOp := OperatorDetails{Precedence: 15, Assoc: Assoc_ltr}
	MultOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	DivOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	ModOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	AddOp := OperatorDetails{Precedence: 13, Assoc: Assoc_ltr}
	SubOp := OperatorDetails{Precedence: 13, Assoc: Assoc_ltr}
	operators['('] = LBracketOp
	operators[')'] = RBracketOp
	operators['*'] = MultOp
	operators['/'] = DivOp
	operators['%'] = ModOp
	operators['+'] = AddOp
	operators['-'] = SubOp
	return operators
}

func reverseSlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func parseTokens(input io.Reader, opStack *[]byte, content *[]byte, operators map[byte]OperatorDetails) error {
	buf := make([]byte, 1)
	_, err := input.Read(buf)
	if err == io.EOF {
		return io.EOF
	}
	if operators[buf[0]].Assoc == Assoc_none {
		(*content) = append((*content), buf...)
	} else {
		if buf[0] == ')' {
			reverseSlice((*opStack))
			var i uint8
			i = 0
			fmt.Println(string(*opStack))
			for (*opStack)[i] != '(' {
				(*content) = append((*content), (*opStack)[0])
				i += 1
			}
			fmt.Println(string(*opStack))
			(*opStack) = append((*opStack), (*opStack)[1:]...)
			return nil
		}
		(*opStack) = append((*opStack), buf...)
		if buf[0] == '(' {
			parseTokens(input, opStack, content, operators)
		}
	}
	return nil

}
