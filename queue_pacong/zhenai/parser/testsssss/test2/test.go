package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	//done := make(chan bool)

	go func() {
		for {
			res,bools:= <-jobs
			fmt.Println(res,bools);
			//j, more := <-jobs
			//fmt.Println(more);
			//if more {
			//	fmt.Println("received job", j)
			//} else {
			//	fmt.Println("received all jobs")
			//	done <- true
			//	return
			//}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		//fmt.Println("sent job", j)
	}
	//close(jobs)

	for{
		time.Sleep(time.Second)
	}
	//<-done
}