package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
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

func TestGoRoutine(t *testing.T) {
	f("direct")

	go f("goroutine")

	go func(s string) {
		fmt.Println(s)
	}("going")

	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TestChannel(t *testing.T) {
	s := time.Now()
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message

	fmt.Println("耗时：", (time.Now()).Sub(s))
	fmt.Println(msg)
}

func TestArray(t *testing.T) {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

func fv(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[2] = a[2] * 2
	fmt.Println(a)
}

func TestFvp(t *testing.T) {
	var ar = [3]int{1, 2, 3}
	fp(&ar)
	fv(ar)
}

func TestSlice(t *testing.T) {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Println("arr1 length:", len(arr1))
	fmt.Println("slice1 length:", len(slice1))
	fmt.Println("slice1 cap length:", cap(slice1))

	slice1 = slice1[0:4]

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}

	fmt.Println("slice1 length:", len(slice1))
	fmt.Println("slice1 cap length:", cap(slice1))

	fmt.Println(sum(arr1[:]))
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func TestRang(t *testing.T) {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	for ix, season := range seasons {
		fmt.Printf("Season %d is %s \n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Println(season)
	}

	for ix := range seasons {
		fmt.Println(ix)
	}
}

func TestReslice(t *testing.T) {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Println("slice1 len:", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
}

func TestCopyAppendSlice(t *testing.T) {
	s1_from := []int{1, 2, 3}
	s1_to := make([]int, 10)
	n := copy(s1_to, s1_from)
	fmt.Println(s1_to)
	fmt.Println("copy len：", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	fmt.Println(compare([]byte("ssss"),[]byte("2222222222")))
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
