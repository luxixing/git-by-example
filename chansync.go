package main
import(
	"fmt"
	"time"
)
func worker(done chan bool){
	fmt.Print("working ... ")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	//去除 <-done go worker 可能还没有开始运行，此处用于阻塞
	<-done
}
