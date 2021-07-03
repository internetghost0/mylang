package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func REPL() {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		result, _ := reader.ReadString('\n')

		l := NewLexer(strings.Trim(result, "\n"))
		err, tokens := l.MakeTokens()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Print("[ ")
			for _, token := range tokens {
				fmt.Print((token.value))
				fmt.Print(" ")
			}
			fmt.Println("]")

		}
	}
}
