package main
import(
	"fmt"
)
type person struct{
	name string
	age int
}
func main(){
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"alice", age:30})
	fmt.Println(person{name:"fred"})
	fmt.Println(&person{name :"an", age :40})
	s:=person{name:"sean", age:30}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println("sp point value:", &sp)
}
