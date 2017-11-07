package main
import(
	"fmt"
)
func main(){
	nums := []int{2,3,5}
	sum := 0 
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index ", i, " num ", num)
		}
	}

	kvs := map[string]string {"a":"apple", "b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//i c 打印的是asci 编码，需要格式
	for i, c := range "go" {
		fmt.Printf("\n%d , %c\n", i, c)
	}
}
