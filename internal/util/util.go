package util

import "fmt"

import (
	"bufio"
	"os"
)

func Loop(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
}

type Employee struct {
	id int
	name string
	salary int
}

func Name() {
	buf := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name:")
	buf.Scan()
	name := buf.Text()
	fmt.Println(name)

}