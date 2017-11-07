package main
import(
	"fmt"
	"os"
	"encoding/json"
)
type(
	Response1 struct {
		Page int
		Fruits[]string
	}
	Response2 struct {
		Page int `json:"page"`
		Fruits []string `json:"fruits"`
	}
)
var p = fmt.Println
func main(){
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(1)
	p(string(intB))

	fltB, _ := json.Marshal(2.24)
	p(string(fltB))

	strB, _ := json.Marshal("gopher")
	p(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	p(string(slcB))

	mapD := map[string]int{"apple":5, "lettuce":7}
	mapB , _ := json.Marshal(mapD)
	p(string(mapB))

	res1D := &Response1{
		Page : 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))
	res2D := &Response2{
		Page : 2,
		Fruits: []string{"appl", "peac", "pea"},
	}
	res2B, _ := json.Marshal(res2D)
	p(string(res2B))

	byt := []byte(`{"num":6.12, "strs":["a", "b"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil{
		panic(err)
	}
	p(dat)

	num := dat["num"].(float64)
	p(num)

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	p(res)
	p(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d)
}
