package test

import (
	"fmt"
	"testing"
)

type Shaper interface {
	Namer() string
}

type Square struct {
	name string
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (s *Square) Namer() string {
	return s.name
}

func TestI1(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 10
	sq1.name = "VV"
	fmt.Println(sq1.Area())
	fmt.Println(sq1.Namer())

	sh := Shaper(sq1)
	fmt.Println(sh.Namer())
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s *stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getValue() float32 {
	return c.price
}

func (c *car) Namer() string {
	return c.make
}

type valueable interface {
	getValue() float32
}

func showValue(value valueable) {
	fmt.Println(value.getValue())
}

func TestI2(t *testing.T) {
	v1 := valueable(&stockPosition{"VV", 10.2, 11.1})
	v2 := valueable(&car{"JiLi", "low", 120000})
	showValue(v1)
	showValue(v2)
}

func TestI3(t *testing.T) {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	sq1.name = "VV"

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok {
		fmt.Println(t)
	}

	if u, ok := areaIntf.(*Square); ok {
		fmt.Println(u)
	} else {
		fmt.Println("Error")
	}
}

func TestI4(t *testing.T) {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	sq1.name = "VV"

	areaIntf = sq1
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Println("square", t)
	case *car:
		fmt.Println("car", t)
	default:
		fmt.Println("unkown")
	}
}

type Stringer interface {
	String() string
}

type S1 struct {
	name string
}

func (s *S1) String() string {
	return s.name
}

type S2 struct {
	name string
}

func (s *S2) String2() string {
	return s.name
}

func TestI5(t *testing.T) {
	s1 := new(S1)
	s2 := new(S2)

	s1.name = "VV"
	s2.name = "CCY"

	var v Stringer

	v = s1

	if sv, ok := v.(Stringer); ok {
		fmt.Println("s1 is Stringer", sv)
	}
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int)  {
	for i:=start;i<=end;i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(lener Lener) bool {
	return lener.Len()*10 > 42
}

func TestI6(t *testing.T)  {
	var list List
	//CountInto(list, 1, 10)
	if LongEnough(list) {
		fmt.Println("list is long enough!")
	}

	plist := new(List)
	CountInto(plist, 1, 10)

	if LongEnough(plist) {
		fmt.Println("plist is long enough!")
	}
}
