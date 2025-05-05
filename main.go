package main

import (
	"fmt"
	"sync"
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

func mergeChannels (ch... chan interface{}) chan interface{}{
	out := make(chan interface{})

	wg := sync.WaitGroup{}

	for i := range ch{
		wg.Add(1)
		
		go func(){
			defer wg.Done()
			for v := range i{
				out <- v
			}
		}()
	}

	go func(){
		wg.Wait()
		close(out)
	}()
	
	
	return out
}