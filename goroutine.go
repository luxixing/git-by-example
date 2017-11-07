package main
import(
	"fmt"
)

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from, ":", i)
	}
}
func main(){
	f("direct")
	go f("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")
	//阻塞主协程，不让其退出有时间执行子协程
	//注释下面则仅仅输出 main 里面第一行f("direct")的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

