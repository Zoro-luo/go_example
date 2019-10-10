package main

import (
	"fmt"
	"time"
)

//向放入intChan 1-8000
func putNum (intChan chan int) {
	for i := 0; i < 80; i++ {
		intChan <- i
	}
	close(intChan)
}

//开启4个协程 从intChan取出数据 并判断是否是素数 如果是就放入primeChan
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	//var num int
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)	//10毫秒
		num,ok := <- intChan
		if !ok {	//intChan 取不到
			break
		}
		flag = true		//假设 为true 那么这个数就是素数
		//判断num是不是素数
		for i :=2;i < num;i++{
			if num % i == 0{	//说明num 不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个素数放入到primeChan 
			primeChan <- num
		}
	}
	fmt.Println("有一个协程取不到数据,退出")
	//暂时还不能关闭 primeChan 
	//向 exitChan 写入 true
	exitChan <- true
}

func main() {
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)	//放入结果
	//退出的管道
	exitChan := make(chan bool,4)	//4

	// 开启一个协程 放入intChan 1-8000个数
	go putNum(intChan)
	// 开启4个协程 从intChan取出数据 并判断是否是素数 如果是就放入primeChan
	//放入到primeChan管道
	for i := 0; i < 4; i++ {
		go primeNum(intChan,primeChan,exitChan)
	}
	//主线程要进行处理 从exitChan管道取4个true
	go func(){
		for i := 0; i < 4; i++ {
			<- exitChan
		}
		//当从exitChan 管道取出4个为true的结果,就可以关闭primeNum 协程
		close(primeChan)
	}()
	
	for {
		res, ok := <- primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n",res)
	}
	fmt.Println("main线程退出")
}