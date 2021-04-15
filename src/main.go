package main

import (
	"learnDemo/src/channel"
	"time"
)

func main() {
	ch := make(chan int, 64)
	go channel.ProducerFunc(1, ch)
	go channel.ProducerFunc(2, ch)

	go channel.ConsumerFunc(ch)
	time.Sleep(5 * time.Second)
}
