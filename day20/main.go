package main

import (
	"fmt"
	"sync"
)

func task(id int , wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("task no :: ", id);
}

func main() {
	var wg sync.WaitGroup;


	for i:= range 1000000 {
		wg.Add(1)
		go task(i,&wg)
	}

	wg.Wait()
}