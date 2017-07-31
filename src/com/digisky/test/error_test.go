package test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var errNotError error =  errors.New("Not found error!!")

func TestError1(t *testing.T)  {
	t1:=time.Now()
	fmt.Println(errNotError)
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(time.Now().Sub(t1))
}

func Sqrt(f float64)(float64, error){
	if f < 0 {
		return 0, errors.New("Math - square root of negative number")
	}
	return f*f, nil
}