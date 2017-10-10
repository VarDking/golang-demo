package main

import (
	"sync"
	"fmt"
)

var wg = sync.WaitGroup{}
var nQuan = 4

func main(){
	peopleId := make(chan int)
	wg.Add(1)
	go runner(peopleId)
	peopleId <- 1
	wg.Wait()
}

func runner(peopleId chan int)  {
	per, _ := <- peopleId
	if per == nQuan{
		fmt.Println("接力跑结束")
		wg.Done()
		return
	}
	fmt.Println("第",per,"棒已经开跑")
	fmt.Println("...🏃中...")
	per++
	if per < nQuan {
		fmt.Println("第",per,"棒已经正在准备")
	}
	go runner(peopleId)
	peopleId <- per
}
