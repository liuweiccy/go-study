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
