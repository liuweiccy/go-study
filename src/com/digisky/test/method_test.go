package test

import (
	"fmt"
	"testing"
)

type TwoInts struct {
	a int
	b int
}

func (ti *TwoInts) AddThen() int {
	return ti.a + ti.b
}

func (ti *TwoInts) AddToParam(param int) int {
	return ti.AddThen() + param
}

func TestMethod1(t *testing.T) {
	two1 := new(TwoInts)
	two1.a = 10
	two1.b = 20

	fmt.Println(two1.AddThen())
	fmt.Println(two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Println(two2.AddThen())
}

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func TestMethod2(t *testing.T) {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}
