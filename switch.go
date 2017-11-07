package main

import(
	"fmt"
	"time"
)

func main(){
	i := 2
	fmt.Println("Write ", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("下午")
	}
	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("布尔")
		case int:
			fmt.Println("整形")
		default:
			fmt.Printf("不知道是个啥 type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
