package pfunc

import "fmt"

func Loop(i int,msg string){
	for i>0{
		fmt.Println(msg,"循环打印i：",i)
		i--
	}
}
