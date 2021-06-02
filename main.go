package main

import (
	"strings"

	"github.com/CalvoM/shunting-yard/shuntingyard"
)

func main() {
	r := strings.NewReader("A+B*C")
	shuntingyard.ToPostFix(r, nil)
}
