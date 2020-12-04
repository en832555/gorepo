package main

import (
	"fmt"
	"gorepo/src/goByExample/pfunc"
	"gorepo/src/goByExample/pstruct"
	"time"
)

func main (){

	fmt.Println("start...")
	p:= pstruct.Person{Name: "李四",Age:10}
	p.Eat()
	fmt.Println(p.Age)
	a,e := pfunc.Divided(20,2)
	if e!=nil{
		fmt.Println("出错：",e.Error())
	}else {
		fmt.Println("结果：",a)
	}

	//分割：---------------------------协程------------------------------

	go pfunc.Loop(30,"A")

	go pfunc.Loop(40,"B")

	//time.Sleep(5000000*1000)
	time.Sleep(time.Second*5)
	fmt.Println("over")

	// 分割 ---------------------------通道------------------------------

	msgChan :=make(chan int,5)
	overChan := make(chan string)

	go func(){
		for i:=0;i<100;{
			msgChan <- i
			i++
		}
	}()

	go func(){
		for i:=0;i<100;{
		fmt.Println("出来的是：",<-msgChan, len(msgChan))
		i++
		if i==10 {
			overChan <- "over"
		}
	}
	}()

	fmt.Println("无缓冲结束",<-overChan)
	//time.Sleep(time.Second*3)

	/*chan10 := make (chan int,1)
	go func (){
		for i:=0;i<10;{
			chan10 <- i
			fmt.Println("写入：",i)
			i++
		}
	}()
	for i:=0;i<10;{
		i++
		//fmt.Println("通道长度", len(chan10))
		fmt.Println("读取：",<-chan10)

	}
	fmt.Println("缓冲结束")*/

	//-----------------------------------非阻塞通道t-----------------------------
	chan1 := make(chan string ,1)
	func(){

		chan1 <- "we in"
	}()
	time.Sleep(3*time.Second)
	fmt.Println(" 出来的是：",<-chan1)

	//-------------------------------阻塞造成死锁，用defaultc解决-----------------
	/*chan4117 := make(chan int ,2)
	for i:=0;i<3;i++{
		fmt.Println("进入循环：。。。")
		select {
		case chan4117 <- i:
			fmt.Println("b")
		default :
			fmt.Println("wrong")
		}
	}
	for i:=0;i<3;i++{
		fmt.Println("出值循环：。。。")
		select {
		case  v:=<- chan4117:
			fmt.Println(v)
		default :
			fmt.Println("wrong")
		}
	}*/

	//-------------------------关闭通道，遍历通道------------------------------
	/*chan9187 := make(chan string,5)
	chan9187 <- "a"
	chan9187 <- "b"
	chan9187 <- "c"
	chan9187 <- "d"

	close(chan9187)

	for elem := range chan9187{
		fmt.Println("出值：",elem)
	}*/

	//------------------------------Timer定时器----------------------------------
	/*fmt.Println(time.Now())
	timer1 := time.NewTimer(time.Second)
	timer1.Stop()
	timer1.Reset(time.Second)
	v := <-timer1.C
	fmt.Println(v)*/

	//------------------------------Ticker打点器----------------------------------
	/*ticker1 := time.NewTicker(time.Second)
	for i:=0;;i++{
		fmt.Println("进入循环:",i)
		select {
		case t:= <-ticker1.C:
			fmt.Println(<-ticker1.C,t)
		case <-time.After(time.Second):
			fmt.Println("default:",i)
		}

		if i==5{
			ticker1.Stop()
			break
		}
	}*/

	//-----------------------------工作池-------------------------------
	//3个工人，需要完成8个件
	/*workers := 4
	nums := 8
	jobs:=make(chan int,nums)
	results:=make(chan int,nums)

	//工人就绪
	for i:=1;i<=workers;i++{
		go pfunc.Worker(i,jobs,results)
	}
	//8个件就绪
	for i:=1;i<=nums;i++{
		jobs<- i*10000
	}
	close(jobs)
	for i:=1;i<=nums;i++{
		<-results
	}*/

	//-----------------------------WaitGroup-------------------------------

/*	workers := 4
	nums := 8

	//8个件就绪
	jobs:=make(chan int,nums)
	for i:=1;i<=nums;i++{
		jobs<- i*10000
	}
	close(jobs)

	//4名工人开始
	var wg sync.WaitGroup
	for i:=1;i<=workers;i++{
		wg.Add(1)
		go pfunc.WorkerWithWG(i,jobs,&wg)
	}
	wg.Wait()*/

   //--------------------------------互斥锁--------------------------------------
/*   states := make(map[int] string)
   var key int
   mutes := &sync.Mutex{}
   go func(){
		for i:=0;i<100;i++{
			mutes.Lock()
			key++
			states[key]="go routine 1--> writing: "+strconv.Itoa(i)
			mutes.Unlock()
		}
	}()
	go func(){
		for i:=0;i<100;i++{
			mutes.Lock()
			key++
			states[key]="go routine 2--> writing: "+strconv.Itoa(i)
			mutes.Unlock()
		}
	}()

   time.Sleep(time.Second*5)
   for k,v:=range states{
   		fmt.Println("key : ",k," ---- value : ",v)
   }*/

	//------------------------------------ sort ------------------------------------------
	/*arr_float64 := []float64{1.1,2.2,4.4,5.5,3.3}
	ints := []int{2,3,5,1,4}
	strs := []string{"2xc","aba","rd","8gb","qe"}
	sort.Float64s(arr_float64)
	fmt.Println(arr_float64)
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Strings(strs)
	fmt.Println(strs)

	//自定义类型pstruct.Students 实现了sort接口
	stu1 := pstruct.Student{"high03117","李四",12}
	stu2 := pstruct.Student{"high02017","史蒂夫",12}
	stu3 := pstruct.Student{"mid044123"," 刚好",13}
	stu4 := pstruct.Student{"low013242","风割",12}
	stu5 := pstruct.Student{"mid11117","饭是钢",14}

	stus := pstruct.Students{stu1,stu2,stu3,stu4,stu5}
	sort.Sort(stus)
	fmt.Println(stus)

	//search
	x:="high02017"
	pos :=sort.Search(len(stus),func (i int) bool{
		return  stus[i].ID > x
	})
	fmt.Println(stus[pos-1])*/

	//-------------------------------- panic defer recover()--------------------------------
	/*end := make(chan bool)
	go func(){
		defer func() {
			err :=recover()
			if err!=nil {
				end <- false
				fmt.Println("recover打印错误",err)
			}
		}()
		_, err := os.Create("/temp/file")
		if err!=nil {
			panic(err)
		}
	}()
	status := <-end
	fmt.Println(status)

	var a *int
	var b int
	fmt.Println(a,b)*/
	//--------------------------------- json --------------------------------------
	/*mapA :=make(map[string]interface{})
	mapA["apple"] = 5
	mapA["pear"]  =7.7
	fmt.Println(mapA)
	mapB,_ := json.Marshal(mapA)
	fmt.Println(string(mapB))

	byt := []byte("{\"a\":[\"john\",\"lisa\"] , \"b\":2.2}")
	fmt.Println(string(byt))
	var dat map[string]interface{}
	if err := json.Unmarshal(byt,&dat);err!=nil{
		panic(err)
	}
	fmt.Println(dat)
	num := dat["b"].(float64)
	fmt.Println(num+2.3)

	strs := dat["a"].([]interface{})
	fmt.Println(strs[1].(string))

	str := "{\"id\": 1002 , \"params\":[\"mch\",\"account\",\"sign\"]}"
	ord := order{}
	if err1 := json.Unmarshal([]byte(str),&ord);err1 !=nil{
		panic(err1)
	}
	fmt.Println(ord.Id,ord.Param)*/

	//------------------------------------ xml ----------------------------------
	/*coffee :=&pxml.Plant{Id:1001 ,Name:"coffee" ,Origin:[]string{"Eth","Blz"}}
	out ,_ :=xml.MarshalIndent(coffee," ","	")
	fmt.Println(xml.Header+string(out))

	var p pxml.Plant
	if err := xml.Unmarshal(out ,&p); err !=nil{
		panic(err)
	}
	fmt.Println(p)


	nesting := &pxml.Nesting{}
	nesting.Child = []*pxml.Plant{coffee,coffee}
	nesting.Parent = []*pxml.Plant{coffee}

	out2 ,_ :=xml.MarshalIndent(nesting, " ","  ")
	fmt.Println(string(out2))*/
}

/*type order struct {
	Id int "json:\"id\""
	Param []string "json:\"params\""
}*/


