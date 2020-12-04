package pfunc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WorkerWithWG(id int, jobs <-chan int ,wg *sync.WaitGroup){
	defer wg.Done()
	for v:= range jobs{
		rand.Seed(time.Now().UnixNano())
		num := (rand.Int63n(10)+1)*300000000
		ran := time.Duration(num)
		fmt.Println("worker:",id," is working for :",v,ran)
		time.Sleep(ran)
		fmt.Println("worker:",id," finished: ",v)
	}
}
