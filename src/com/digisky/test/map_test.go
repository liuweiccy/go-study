package test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestMap1(t *testing.T) {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Println("one:", mapLit["one"])
	fmt.Println("key2:", mapCreated["key2"])
	fmt.Println("two:", mapLit["two"])
	fmt.Println("ten:", mapLit["ten"])
	fmt.Println(mapLit)
	fmt.Println(mapAssigned)
	fmt.Println(mapCreated)
}

func TestMap2(t *testing.T) {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}

	fmt.Println(mf)
	fmt.Println(mf[1]())
}

func TestMap3(t *testing.T) {

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25

	if value, ok := map1["Beijing"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no found Beijing")
	}

	delete(map1, "Washington")

	if value, ok := map1["Washington"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no found Washington")
	}
}

func TestMap4(t *testing.T) {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0
	for key, value := range map1 {
		fmt.Println(key, "->", value)
	}
}

func TestMap5(t *testing.T) {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][2] = rand.Int()
	}

	items2 := make([]map[int]int, 5)
	for _, value := range items2 {
		// 只是一个copy
		value = make(map[int]int, 1)
		value[3] = 3
	}

	fmt.Println(items)
	fmt.Println(items2)
}

var barVal = map[string]int{
	"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87}

func TestMap6(t *testing.T) {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key : %v, value : %v\n", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for key := range barVal {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("key : %v, value : %v\n", k, barVal[k])
	}
}
