package shuntingyard

import (
	"fmt"
	"io"
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

func ToPostFix(input io.Reader, operatorMap map[byte]OperatorDetails) {
	content := make([]byte, 0)
	opStack := make([]byte, 0)
	if operatorMap == nil {
		operatorMap = getDefaultOperatorMap()
	}
	buf := make([]byte, 1)
	for {
		_, err := input.Read(buf)
		if err == io.EOF {
			break
		}
		if operatorMap[buf[0]].Assoc == Assoc_none {
			content = append(content, buf...)
		} else {
			opStack = append(opStack, buf...)
		}

	}
	if len(opStack) != 0 {
		content = append(content, opStack...)
	}
	fmt.Println(string(content))

}

func getDefaultOperatorMap() map[byte]OperatorDetails {
	operatorMap := make(map[byte]OperatorDetails)
	LBracketOp := OperatorDetails{Precedence: 15, Assoc: Assoc_ltr}
	RBracketOp := OperatorDetails{Precedence: 15, Assoc: Assoc_ltr}
	MultOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	DivOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	ModOp := OperatorDetails{Precedence: 14, Assoc: Assoc_ltr}
	AddOp := OperatorDetails{Precedence: 13, Assoc: Assoc_ltr}
	SubOp := OperatorDetails{Precedence: 13, Assoc: Assoc_ltr}
	operatorMap['('] = LBracketOp
	operatorMap[')'] = RBracketOp
	operatorMap['*'] = MultOp
	operatorMap['/'] = DivOp
	operatorMap['%'] = ModOp
	operatorMap['+'] = AddOp
	operatorMap['-'] = SubOp
	return operatorMap
}
