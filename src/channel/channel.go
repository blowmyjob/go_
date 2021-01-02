package channel

import "fmt"

func Reveiver(c chan int) {
	for {
		data := <-c
		if data == 0 {
			fmt.Println("我完事了")
			break
		}
		fmt.Println(data)
	}
	c <- 0
}

func Producer() {
	c := make(chan int)
	// 并发执行printer, 传入channel
	go Reveiver(c)
	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		c <- i
	}
	// 通知并发的printer结束循环(没数据啦!)
	c <- 0
	// 等待printer结束(搞定喊我!)
	<-c
}
