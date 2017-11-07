package main
import(
	"time"
	"fmt"
)
func main(){
	requests := make(chan int, 5)
	for i:= 1; i<=5; i++ {
		requests <- i
	}
	close(requests)
	//每2秒处理一次请求
	limiter := time.Tick(time.Second * 2)
	for req := range requests {
		<- limiter
		fmt.Println("requst ", req, time.Now())
	}
	//初始化一个限流器,限制最高并发数3，并且使用当前时间填充
	burstyLimiter := make(chan time.Time, 3)
	for i:= 0; i< 3; i++{
		burstyLimiter <- time.Now()
	} 
	//每两秒允许最大请求数量3
	go func(){
		for t:= range time.Tick(time.Second * 2){
			burstyLimiter<- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i<=5; i++{
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("requst",req, time.Now())
	}
}
