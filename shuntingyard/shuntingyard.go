package shuntingyard

import (
	"fmt"
	"io"
)

type OperatorAssociative uint8

const (
	Assoc_none OperatorAssociative = iota
	Assoc_left
	Assoc_right
)

type OperatorDetails struct {
	Precedence int
	Assoc      OperatorAssociative
}

func ToPostFix(input io.Reader) {
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
