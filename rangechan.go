package main
import(
	"fmt"
)
func main(){
	ch := make(chan map[string]int, 5)
	m1 := make(map[string]int)
	m1["a"] = 1
	m2 := make(map[string]int)
	m2["b"] = 2
	ch<-m1
	ch<-m2
	close(ch)
	for v := range ch{
		fmt.Println("range map channel k-v ", v)
	}
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	close(queue)
	for e := range queue{
		fmt.Println(e)
	}

}
