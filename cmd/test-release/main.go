package main

import (
	"fmt"
	c "github.com/0xRyuk/test-release/internal/core"
	u "github.com/0xRyuk/test-release/internal/util"
)

func main() {
	fmt.Println("Test-release")
	fmt.Println("Square of 5 is", c.Square(5))
	name("Ryuk")
	u.Loop(10)
	u.Name()
	fmt.Println("added something else")

}

func name(str string) {
	fmt.Println("Hello ,", str)
}