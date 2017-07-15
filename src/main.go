package main

import "fmt"
import (
	"com/digisky/hello"
	"com/digisky/trans"
)

func main()  {
	hello.Name()

	fmt.Println("Hello:", 23, "\n", 45)

	fmt.Println(trans.Pi)
}