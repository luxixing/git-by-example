package main
import(
	"fmt"
	"time"
	"sync/atomic"
)
func main(){
	var ops uint64 = 0
	for i := 0; i<50;i++{
		go func(){
			//主协程退出则主协程创建的协程也退出
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
