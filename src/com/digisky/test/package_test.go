package test

import (
	"container/list"
	"fmt"
	"go/types"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"testing"
	"unsafe"
)

func TestPavkage1(t *testing.T) {
	L1 := list.New()

	L1.PushBack(101)
	L1.PushBack(102)
	L1.PushBack(103)

	L1.PushFront(101)
	L1.PushFront(102)
	L1.PushFront(103)

	for e := L1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestPavkage2(t *testing.T) {
	fmt.Println(unsafe.Sizeof(types.Int64))
}

func TestPackage3(t *testing.T) {
	searchIn := "John: 2578.32 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		fmt.Println("number:", v)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

func TestPackage4(t *testing.T) {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1990)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Println(ip)
}
