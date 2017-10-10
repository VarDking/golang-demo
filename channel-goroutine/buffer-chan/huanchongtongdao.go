package main

import (
	"sync"
	"fmt"
	"strconv"
)

const (
	Worker = 4
	WorkCounts = 20
)

var wg = sync.WaitGroup{}

func main()  {

	tasks := make(chan string, WorkCounts)
	wg.Add(Worker)
	for i:=0; i<Worker; i++ {
		go work(i, tasks)
	}
	//通道中装载任务数据，装载完后关闭通道
	for i:=0; i<WorkCounts; i++ {
		tasks <- "{\"taskNo\":"+strconv.Itoa(i)+"}"
	}
	close(tasks)
	wg.Wait()
}

func work(workNo int,tasks chan string)  {
	defer wg.Done()
	for{
		task, ok := <-tasks
		//带缓冲通道，如果取不到数据的话，显示false
		if !ok{
			fmt.Println("工人[",workNo,"]失业了，退出")
			return
		}
		fmt.Println("工人[",workNo,"]消费了任务:"+task)
	}
}
