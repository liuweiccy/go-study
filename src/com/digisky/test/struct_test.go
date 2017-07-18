package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

type student struct {
	name string
	age  int8
}

var name student

func TestST1(t *testing.T) {
	name.name = "Eric"
	name.age = 24
	fmt.Println(name)
}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func TestST2(t *testing.T) {
	t1 := time.Now()
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 13.5
	ms.str = "ccy"
	fmt.Println(ms.i1)
	fmt.Println(ms.f1)
	fmt.Println(ms.str)
	fmt.Println(ms)

	m1 := &struct1{1, 2, "ew"}
	var m2 struct1
	m2 = struct1{1, 2, "ew"}

	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(time.Now().Sub(t1))

	m3 := struct1{1, 2, "ew"}
	m4 := struct1{i1: 1, f1: 2}

	fmt.Println(m3)
	fmt.Println(m4)
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func TestST3(t *testing.T) {
	var pers1 Person
	pers1.firstName = "liu"
	pers1.lastName = "wei"
	upPerson(&pers1)
	fmt.Println(pers1)

	pers2 := new(Person)
	pers2.firstName = "chen"
	(*pers2).lastName = "cy"
	upPerson(pers2)
	fmt.Println(pers2)

	pers3 := &Person{"eric", "liu"}
	upPerson(pers3)
	fmt.Println(pers3)
}

type TagType struct {
	f1 bool   "f1 aaa"
	f2 string "f2 sss"
	f3 int    "f3 mmm"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	filed := ttType.Field(ix)
	fmt.Println(filed)
	fmt.Println(filed.Tag)
	fmt.Println(filed.Name)
	fmt.Println(filed.Type)
}

func TestST4(t *testing.T) {
	tt := TagType{true, "VV", 27}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func TestST5(t *testing.T) {
	outer := new(outerS)
	outer.b = 6
	outer.c = 1.2
	outer.int = 100
	outer.in1 = 12
	outer.in2 = 21

	fmt.Println(outer)
	fmt.Println(outer.b)
	fmt.Println(outer.c)
	fmt.Println(outer.int)
	fmt.Println(outer.in1)
	fmt.Println(outer.in2)

	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println(outer2)
}

type A struct {
	a int
}

type B struct {
	a, b int
}

type C struct {
	A
	B
}

type D struct {
	B
	b float32
}

func TestST6(t *testing.T)  {
	c1 := new(C)
	d1 := new(D)

	fmt.Println(c1.A.a)
	fmt.Println(d1.B.b)
}