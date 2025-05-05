package main

import (
	"fmt"
	"sync"
)

func main(){
	count := 10
	wg := sync.WaitGroup{}

	wg.Add(count)
	go func(){
	
		for i := range count{
			defer wg.Done()
			fmt.Println("hello world", i)
		}
	}()


	wg.Wait()

	fmt.Println("end")
	


}	