package main

import "learnDemo/src/kafka"

func main() {
	kafka.InitProducer("120.79.223.58:9092")
	for {
		kafka.Send("test1", "haha")
	}
}
