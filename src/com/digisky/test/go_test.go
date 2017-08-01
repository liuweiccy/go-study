package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGo1(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

func TestGo2(t *testing.T) {
	fmt.Println("in main")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main")
	time.Sleep(10000)
	fmt.Println("end main")
}

func longWait() {
	fmt.Println("Beginning longWait")
	time.Sleep(5000)
	fmt.Println("end longWait")
}
func shortWait() {
	fmt.Println("Beginning shortWait")
	time.Sleep(2000)
	fmt.Println("end shortWait")
}

func TestGo3(t *testing.T) {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "VV"
	ch <- "YY"
	ch <- "BB"
}

func getData(ch chan string) {
	for {
		name := <-ch
		fmt.Println(name)
	}
}

func TestGo4(t *testing.T) {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	fmt.Println(<-ch)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func f1(in chan int) {
	fmt.Println(<-in)
}
func TestGo5(t *testing.T) {
	out := make(chan int)
	go func() { out <- 2 }()
	go f1(out)
}

func compute(ch chan int) {
	fmt.Println("计算中...")
	time.Sleep(1e9)
	ch <- 1
}
func TestGo6(t *testing.T) {
	ch := make(chan int)
	go compute(ch)
	fmt.Println("等待完成")
	<-ch
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func TestGo7(t *testing.T)  {
	stream := pump2()
	go suck(stream)
}

func source(ch chan <- int)  {
	for {
		ch <- 1
	}
}

func sink(ch <- chan int)  {
	for {
		fmt.Println(<-ch)
	}
}

func TestGo8(t *testing.T)  {
	var c = make(chan int)
	defer close(c)
	go source(c)
	go sink(c)
	time.Sleep(1)
}
