package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	_, _ = os.Create("proverb.txt")
	file,_:=os.OpenFile("proverb.txt",os.O_WRONLY|os.O_CREATE,0777)
	_, _ = file.WriteString("don't communicate by sharing memory share memory by communicating")
	f, err := os.Open("proverb.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//关闭文件
	defer f.Close()
	//Read方法从f中读取最多len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。文件终止标志是读取0个字节且返回值err为io.EOF
	//定义切片保存读取的数据，要指定容量
	var b []byte = make([]byte, 2*1024)
	n, errR := f.Read(b)
	//出错，同时没有到末尾
	if errR != nil && errR != io.EOF {
		fmt.Println(errR.Error())
	}
	fmt.Println(n, string(b))
}