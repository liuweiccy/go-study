package main

import "fmt"
import (
	"com/digisky/hello"
	"com/digisky/trans"
)

func main()  {
	hello.Name()

	fmt.Println("Hello:", 23, 45)

	fmt.Println(trans.Pi)
}
