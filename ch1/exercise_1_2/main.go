package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(arg + " " + strconv.Itoa(index))
	}
}
