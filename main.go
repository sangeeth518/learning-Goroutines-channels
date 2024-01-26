package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go routines")
	go new()
	new()
	go func() {
		fmt.Println("hello from anonymous goroutine")
	}()
	time.Sleep(1 * time.Second)

	ch := make(chan int)
	go multiplywithchannel(ch)
	ch <- 10

	ch1 := make(chan int)
	close(ch1)
	k, ok := <-ch1
	fmt.Println(ok, k)

	c := make(chan string)
	go initstring(c)
	time.Sleep(2 * time.Second)
	for {
		m, ok := <-c
		if ok == false {
			fmt.Println("channel close", ok)
			break
		}
		fmt.Println("from main channel", m, ok)
	}

	// bidirectional channels

	chnl1 := make(chan string)
	chnl2 := make(chan string)
	go send(chnl1)
	valuefromchannel := <-chnl1
	fmt.Println("chnl1", valuefromchannel)

	go recieve(chnl2)
	chnl2 <- valuefromchannel

}

func new() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("studing goroutines and channels")
	}
}

func multiplywithchannel(s chan int) {
	fmt.Println("value is :", 100*<-s)
}
func initstring(chnll chan string) {
	for v := 0; v < 3; v++ {
		chnll <- "hi from instring"
	}
	close(chnll)

}

func send(s chan string) {
	s <- "bidirectional channels"
}
func recieve(s chan string) {
	fmt.Println(<-s)

}
