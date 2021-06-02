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
	buf := make([]byte, 10)
	content := make([]byte, 1)
	for {
		_, err := input.Read(buf)
		if err == nil {
			content = append(content, buf...)
		} else if err == io.EOF {
			break
		}
	}
	fmt.Println(string(content))

}

func getDefaultOperatorMap() map[byte]OperatorDetails {
	operatorMap := make(map[byte]OperatorDetails)

	return operatorMap
}
