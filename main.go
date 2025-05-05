package main

import (
	"fmt"
	"time"
)

type emptyStruct struct{}

func main(){

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	
	ch1 <- "Hello"
	ch2 <- 1123
	ch3 <- emptyStruct{} 

	fmt.Println(mergeChannels(ch1, ch2, ch3))

}	

func mergeChannels (ch ... chan interface{}) <- chan interface{}{
	mergedCannel := make(<- chan interface{})

	go func(){
		select{
		case <- time.After(5*time.Second):
			fmt.Println("time out")
		}
	}()
	
	return mergedCannel
}