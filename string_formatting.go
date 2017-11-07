package main
import(
	"fmt"
	"os"
)
var  fp = fmt.Printf
type point struct {
	x, y int
}
func main(){
	p := point{1, 2}
	fp("%v\n",p)
	fp("%+v\n", p)
	fp("%#v\n", p)
	fp("%T\n", p)
	fp("%t\n", true)
	fp("%d\n", 123)
	fp("%b\n", 14)
	fp("%c\n", 33)
	fp("%x\n",456) 
	fp("%f\n", 78.09) 
	fp("%e\n", 123400000.0) 
	fp("%E\n", 123400000.0) 
	fp("%s\n", "\"string\"") 
	fp("%q\n", "\"string\"") 
	fp("%x\n", "hex this") 
	fp("%p\n", &p) 
	fp("|%6d|%6d|\n", 12, 345) 
	fp("|%6.2f|%6.2f|\n", 1.2, 3.45) 
	fp("|%-6.2f|%-6.2f|\n", 1.2, 3.45) 
	fp("|%6s|%6s|\n", "foo", "b") 
	fp("|%-6s|%-6s|\n", "foo", "b") 
	s := fmt.Sprintf("a %s", "string")
	fp(s)
	fmt.Fprintf(os.Stderr, "an %s \n", "error")

}
