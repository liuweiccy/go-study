package test

import (
	"fmt"
	"reflect"
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

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(lener Lener) bool {
	return lener.Len()*10 > 42
}

func TestI6(t *testing.T) {
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

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func TestI7(t *testing.T) {
	data := IntArray{74, 59, 238, -748, 9845, 0, 0, 42}
	sorter := Sorter(data)
	Sort(sorter)
	fmt.Println(data)
}

func TestI8(t *testing.T) {
	var x float64 = 34.23
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
}

func TestI9(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("v.type:", v.Type())
	fmt.Println("v.kind:", v.Kind())
	fmt.Println("v.Float:", v.Float())
	fmt.Println(v.Interface())
	y := v.Interface().(float64)
	fmt.Println("y:", y)
	fmt.Println("y.type:", reflect.TypeOf(y))
}

func TestI10(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("setability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("setability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("The Elem of v is:", v)
	fmt.Println("setability of v:", v.CanSet())
	v.SetFloat(3.45)
	fmt.Println(v.Interface())
	fmt.Println(v)
}

type NotKnowType struct {
	s1, s2, s3 string
}

func (n NotKnowType) String() string {
	return n.s1 + "." + n.s2 + "." + n.s3
}

func TestI11(t *testing.T) {
	var secret interface{} = NotKnowType{"VV", "YY", "BB"}
	value := reflect.ValueOf(secret)
	fmt.Println("type:", value.Type())
	fmt.Println("kind:", value.Kind())

	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)
}

type Ts struct {
	A int
	B string
}

func TestI12(t1 *testing.T) {
	t := Ts{23, "VV"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(i," ", typeOfT.Field(i).Name, " ", f.Type(), " ", f.Interface())
	}
	s.Field(0).SetInt(21)
	s.Field(1).SetString("YY")
	fmt.Println(s)
}
