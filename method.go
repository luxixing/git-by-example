package main
import(
	"fmt"
)
type rect struct{
	width, height int
}
func(r *rect) area()int{
	return r.width*r.height
}
func(r rect) perim()int{
	return 2*r.width+2*r.height
}
func main(){
	//使用指针接收者可以避免方法调用时的复制,也可以使得接收者调用方法时改变自身结构数据
	//接受者是值
	r := rect{width:10, height:5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	//接受者是指针
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}


