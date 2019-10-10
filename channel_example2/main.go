package main
//使用*select 解决从管道取数据的阻塞问题
import (
	"fmt"
	//"time"
)



func main() {
	//定义一个管道 10个数据 int
	intChan := make(chan int,10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//定义一个管道 5个数据string
	stringChan := make(chan string,5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello_" + fmt.Sprintf("%d",i)
	}

	//传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock
	//问题是 实际开发中,我们不好确定在哪关闭该管道 可以使用select 方式解决
	for {
		select {
			//注意：这里,如果intChan一直没有关闭,不会一直阻塞而deadlock
			//会自动到下一个case匹配
			case v := <- intChan :
				fmt.Printf("从intChan读取数据%d\n",v)
			case v := <- stringChan :
				fmt.Printf("从stringChan读取数据%s\n",v)
			default:
				fmt.Println("都找不到了! 开发者可以加入自己的逻辑")
				return  
		}
	}
}	