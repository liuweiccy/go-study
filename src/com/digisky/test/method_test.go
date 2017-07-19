package test

import (
	"fmt"
	"math"
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

type B1 struct {
	thing int
}

func (b *B1) change() {
	b.thing = 1
}

func (b B1) write() string {
	return fmt.Sprint(b)
}

func TestMethod3(t *testing.T) {
	var b1 B1
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B1)
	b2.change()
	fmt.Println(b2.write())
}

type List []int

func (l List) len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func TestMethod4(t *testing.T) {
	var lst List
	lst.Append(1)
	fmt.Println(lst, lst.len())

	plst := new(List)
	plst.Append(2)
	fmt.Println(plst, plst.len())
}

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	Point
	name string
}

func (point *NamePoint) Abs() float64 {
	return point.Point.Abs() * 100
}

func TestMethod5(t *testing.T) {
	n := &NamePoint{Point{3, 4}, "三角形"}
	fmt.Println(n.Abs())
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (log *Log) Add(s string) {
	log.msg += "\n" + s
}

func (log *Log) String() string {
	return log.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func TestMethod6(t *testing.T) {
	c := new(Customer)
	c.Name = "VV"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	c.Log().Add("2 - Hello,Hello")
	fmt.Println(c.Log())

	c2 := &Customer{"CCY", &Log{"1 - Yes we can!"}}
	c2.Log().Add("2 - Hello,Hello")
	fmt.Println(c2.Log())
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}
func (base Base) MoreMagic() {
	base.Magic()
	base.Magic()
}

type VooDoo struct {
	Base
}

func (VooDoo) Magic() {
	fmt.Println("voodoo magic")
}

func TestMethod7(t *testing.T) {
	v := new(VooDoo)
	v.Magic()
	v.MoreMagic()
}
