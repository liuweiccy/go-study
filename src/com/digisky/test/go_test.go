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

func TestGo7(t *testing.T) {
	stream := pump2()
	go suck(stream)
}

func source(ch chan<- int) {
	for {
		ch <- 1
	}
}

func sink(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func TestGo8(t *testing.T) {
	var c = make(chan int)
	defer close(c)
	go source(c)
	go sink(c)
	time.Sleep(1)
}

func TestGo9(t *testing.T) {
	ch := make(chan int)
	go generate(ch)
	time.Sleep(1)
	for {
		prime := <-ch
		fmt.Println(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch, prime)
		ch = ch1
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func TestGo10(t *testing.T) {
	sendChan := make(chan int)
	recvChan := make(chan string)
	go processChan(sendChan, recvChan)
	go func() {
		for i := 0; i < 10; i++ {
			sendChan <- i
		}
	}()
	go func() {
		for {
			fmt.Println("收到值", <-recvChan)
		}
	}()
	time.Sleep(1e4)
}

func processChan(in <-chan int, out chan<- string) {
	for inValue := range in {
		fmt.Println(inValue)
		result := "OK"
		out <- result
	}
}

//14.9
func TestGo11(t *testing.T) {
	ch := make(chan string)
	go sendData1(ch)
	getData1(ch)
}

func sendData1(ch chan string) {
	ch <- "Beijing"
	ch <- "ShangHai"
	ch <- "ChengDu"
	ch <- "ShenZhen"
	ch <- "ChongQing"
	close(ch)
}

func getData1(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			fmt.Println("通道已经关闭...")
			break
		}
		fmt.Println(input)
	}
}

// 14.10
func TestG12(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump3(ch2)
	go suck1(ch1, ch2)
	time.Sleep(1)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump3(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck1(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("接受到来自channel—1的值：", v)
			break
		case v := <-ch2:
			fmt.Println("接受到来自channel—2的值：", v)
			break
		}
	}
}

func TestGo13(t *testing.T)  {
	tick := time.Tick(1e8)
	boom := time.After(5e8)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(5e7)
		}
	}
}
