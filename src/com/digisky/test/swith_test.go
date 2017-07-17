package test

import (
	"testing"
	"fmt"
	"time"
)

func TestSwitch(t *testing.T)  {
	i := 2
	fmt.Print("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	s := time.Now()
	switch {
	case s.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch s := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't konw type ", s)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func TestGoRoutine(t *testing.T)  {
	f("direct")

	go f("goroutine")

	go func(s string) {
		fmt.Println(s)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TestChannel(t *testing.T)  {
	s := time.Now()
	message := make(chan string)

	go func() {message <- "ping"}()

	msg := <-message

	fmt.Println("耗时：", (time.Now()).Sub(s))
	fmt.Println(msg)
}

func TestArray(t *testing.T)  {
	var arr1[5] int

	for i:=0;i<len(arr1);i++ {
		arr1[i] = i*2
	}

	for i:=0; i<len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

func fv(a [3]int)  {
	fmt.Println(a)
}

func fp(a *[3]int)  {
	a[2] = a[2] * 2
	fmt.Println(a)
}

func TestFvp(t *testing.T)  {
	var ar = [3]int{1,2,3}
	fp(&ar)
	fv(ar)
}