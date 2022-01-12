package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}
	go sayHello("Ajung")
	go sayHello("Setia")
	go sayHello("Fourth")
	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)

	var messages2 = make(chan string)
	for _, each := range []string{"Ajung", "Setia", "Fourth"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages2 <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessages(messages2)
	}

	go func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages2 <- data
	}("Name")
	var name = <-messages2
	fmt.Println(name)

	// buffered
	messages3 := make(chan int, 2)
	go func() {
		for {
			i := <-messages3
			fmt.Println("receive ", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send", i)
		messages3 <- i
	}

	// select
	var numbers = []int{1, 2, 3}
	fmt.Println("number ", numbers)
	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)
	var ch2 = make(chan int)
	go getMax(numbers, ch2)
	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg: %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("Max: %d\n", max)
		}
	}

	// range & close
	go sendMessage(messages2)
	printsMessage(messages2)

	// timeout
	rand.Seed(time.Now().Unix())
	var messages4 = make(chan int)
	go sendData(messages4)
	retrieveData(messages4)
}

func printMessages(what chan string) {
	fmt.Println(<-what)
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printsMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive `, data, "\n")
		case <-time.After(time.Second * 3):
			fmt.Println("timeout. no activities under 3 seconds")
			break loop
		}
	}
}
