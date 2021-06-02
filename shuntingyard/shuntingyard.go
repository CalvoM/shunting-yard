package shuntingyard

import (
	"errors"
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

var ErrRParensFound = errors.New("ErrRParensFound")

type OperatorDetails struct {
	Precedence int
	Assoc      OperatorAssociative
}

//ToPostFix Converts any Infix notation streamed in input variable to PostFix Notation
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
	fmt.Println(string(buf[0]), string(*opStack), string(*content))
	if operators[buf[0]].Assoc == Assoc_none {
		*content = append(*content, buf...) // Non-operators should be taken straight to postfix
	} else {
		if buf[0] == ')' {
			stackLen := len(*opStack)
			for (*opStack)[stackLen-1] != '(' {
				*content = append(*content, (*opStack)[stackLen-1])
				*opStack = (*opStack)[:stackLen-1]
				stackLen = len(*opStack)
			}
			*opStack = (*opStack)[:stackLen-1] //The LParens should not be in Postfix
			return ErrRParensFound
		} else if buf[0] == '(' { //Like starting new scope of parsing
			(*opStack) = append((*opStack), buf...)
			var parse_err error
			for parse_err != ErrRParensFound {
				parse_err = parseTokens(input, opStack, content, operators)
			}
		} else {
			stackLen := len(*opStack)
			if stackLen > 0 {
				curSymbol := operators[buf[0]]
				stackSymbol := operators[(*opStack)[stackLen-1]]
				if curSymbol.Precedence > stackSymbol.Precedence {
					(*opStack) = append((*opStack), buf...)
				} else if stackSymbol.Precedence > curSymbol.Precedence {
					if (*opStack)[stackLen-1] != '(' { // If the top of the stack is ( then no need to pop
						*content = append(*content, (*opStack)[stackLen-1])
						*opStack = (*opStack)[:stackLen-1]
						(*opStack) = append((*opStack), buf...)
					} else {
						(*opStack) = append((*opStack), buf...)
					}
				} else if stackSymbol.Precedence == curSymbol.Precedence { //Compare the associativity
					if curSymbol.Assoc == Assoc_rtl {
						(*opStack) = append((*opStack), buf...)
					} else if curSymbol.Assoc == Assoc_ltr {
						*content = append(*content, (*opStack)[stackLen-1])
						*opStack = (*opStack)[:stackLen-1]
						(*opStack) = append((*opStack), buf...)
					}
				}
			} else {
				(*opStack) = append((*opStack), buf...)
			}
		}
	}
	return nil

}
