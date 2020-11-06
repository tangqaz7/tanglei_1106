package main

import (
	"fmt"
)

func newTask() {
	i := 0
	for i <100{
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
	}
}

func main() {
	//创建一个 goroutine，启动另外一个任务
	go newTask()

	i := 0
	//main goroutine 循环打印
	for i < 100{
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
	}
}