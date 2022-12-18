package main

import (
	"fmt"
	c "github.com/0xRyuk/test-release/internal/core"
)

func main() {
	fmt.Println("Test-release")
	fmt.Println("Square of 5 is", c.Square(5))
	name("Ryuk")
}

func name(str string) {
	fmt.Println("Hello ,", str)
}