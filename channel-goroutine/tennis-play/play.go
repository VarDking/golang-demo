package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

//创建一场球赛，两个选手，打印出第几拍时谁输了
var wg = sync.WaitGroup{}

//获取随机数因子
func init()  {
	fmt.Println("获取随机数因子")
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	fmt.Println("比赛开始")
	pai := make(chan int)
	go player("张三", pai)
	go player("李四", pai)
	pai <- 1
	wg.Wait()
}

func player(name string, pai chan int){
	fmt.Println("[",name,"]:已就位")
	wg.Add(1)
	for{
		//等待球被击过来
		ball, err := <- pai
		if !err{
			fmt.Println("[",name,"]:已胜出")
			wg.Done()
			return
		}
		//获取随机数，判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("[",name,"]:击球失误")
			close(pai)
			wg.Done()
			return
		}
		//击球详情
		fmt.Println("[",name,"]:第",ball,"拍" )
		ball++
		//将球放回通道
		pai <- ball
	}
}

