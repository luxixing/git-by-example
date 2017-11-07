package main
import(
	"fmt"
)
func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func(){
		for{
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			}else{
				fmt.Println("received all jobs")
				done <- true
				return 
			}
		}
	}()
	for i:= 1; i<4;i++{
		jobs <- i
		fmt.Println("send job ", i)
	}
	close(jobs)
	fmt.Println("send all jobs")
	//阻塞 main 进程
	<-done
}
