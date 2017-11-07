package main
import(
	"os"
	"fmt"
)
func main(){
	argWithProg := os.Args
	argWithOutProg := argWithProg[1:]
	fmt.Println("with :", argWithProg)
	fmt.Println("withOut :", argWithOutProg)

	arg := os.Args[3]
	fmt.Println(arg)
}
