package main
import(
	"time"
	"fmt"
)
func main(){
	ticker := time.NewTicker(time.Second * 2)
	go func(){

		for t := range ticker.C{
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second *6)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
