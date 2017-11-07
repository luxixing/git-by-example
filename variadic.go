package main
import(
	"fmt"
)

//自动参数
func sum(nums ...int){
	fmt.Print(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main(){
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	//自动展开
	sum(nums...)
}
