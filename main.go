package main

import (
	"strings"

	"github.com/CalvoM/shunting-yard/shuntingyard"
)

func main() {
	r := strings.NewReader("A+B")
	shuntingyard.ToPostFix(r, nil)
}
