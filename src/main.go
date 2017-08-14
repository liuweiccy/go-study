package main

import "fmt"
import (
	"com/digisky/hello"
	"com/digisky/trans"
	"com/digisky/net"
)

func main()  {
	hello.Name()

	fmt.Println("Hello:", 23, "\n", 45)

	fmt.Println(trans.Pi)

	i := 1/0
	fmt.Println(i)
	test.Client()
}