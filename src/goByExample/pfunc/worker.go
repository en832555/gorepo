package pfunc

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int ,finished chan<- int){
	for v:= range jobs{
		fmt.Println("worker:",id," is working for :",v)
		time.Sleep(time.Second)
		fmt.Println("worker:",id," finished: ",v)
		finished<-v
	}
}
