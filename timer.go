package main

import(
	"time"
	"fmt"
)

func main(){
	timer1 := time.NewTimer(time.Second * 4)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTimer(time.Second * 8)
	go func(){
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stoped")
	}
}
