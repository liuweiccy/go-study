package test

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

var errNotFound error = errors.New("Not found error!")

func Sqrt(f float64)(float64, error){
	if f < 0 {
		return 0, errors.New("Math - square root of negative number")
	}
	return f*f, nil
}

func TestErr1(t *testing.T) {
	fmt.Println(errNotFound)
}

func TestErr2(t *testing.T) {
	fmt.Println("Starting the program")
	//panic("fatal error")
	fmt.Println("Ending the program")
	log.Println("test error")
}

func TestErr3(t *testing.T) {
	fmt.Println("Starting the program")
	test()
	fmt.Println("Ending the program")
}
func badCall() {
	panic("Bad end")
}
func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

type ParseError struct {
	Index int
	Word  string
	Err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintln(e.Word)
}

func Parse(input string) (numbers []int, err error) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		var ok bool
	//		err, ok = r.(error)
	//		if !ok {
	//			err = fmt.Errorf("Err -- pkg:%v", r)
	//		}
	//	}
	//}()
	fields := strings.Fields(input)
	numbers = field2numbers(fields)
	return
}
func field2numbers(fileds []string) (numbers []int) {
	if len(fileds) == 0 {
		panic("Panic -- no words to parse")
	}
	for idx, filed := range fileds {
		num, err := strconv.Atoi(filed)
		if err != nil {
			panic(&ParseError{idx, filed, err})
		}
		numbers = append(numbers, num)
	}
	return
}

func TestErr4(t *testing.T) {
	var examples = []string{
		"100 50 25 12.5 6.25",
		//"1 2 3 4 5",
		"1st class",
		//"1 2 3 4 5",
		"",
	}
	for _, ex := range examples {
		time.Sleep(1e8)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					var ok bool
					if !ok {
						fmt.Errorf("Err -- pkg:%v", r)
					}
				}
			}()
			//fmt.Printf("Parsing %q:\n", ex)
			nums, err := Parse(ex)
			if err != nil {
				fmt.Println("失败：",err)
			}
			fmt.Println("成功：",nums)
		}()
	}
	time.Sleep(1e9)
}
