package channel

import "fmt"

func ProducerFunc(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func ConsumerFunc(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
