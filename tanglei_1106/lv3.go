package main

import (
	"fmt"
)
var over = make(chan bool)

func get() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 9 {
			over <- true
		}
	}
}
//func main(){
//	go get()
//	<- over
//	println("over")
//}
//func main() {
//	over := make(chan bool)
//	for i := 0; i < 10; i++ {
//		多协程出问题了
//		go func() {
//			fmt.Println(i)
//		}()
//		if i == 9 {
//			over <- true
//		}
//	}
//	<-over
// 通道堵塞
//	fmt.Println("over!!!")
//}
//这个问题我也不好说，但我解决了>\*^*/<