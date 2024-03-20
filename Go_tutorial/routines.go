package main

import (
	"fmt"
	"sync"
	"time"
)

func routine_implement(){
	var wg=sync.WaitGroup{}//tasks will be added to this group
	t1:=time.Now()//current time before starting
	for i:=0;i<10;i++{
		wg.Add(1)//adds task to the group
		go main_concurrency(i*100,(i+1)*100,&wg)//calls the go routine
	}
	wg.Wait()//waits until all the tasks in the group are completed
	fmt.Println()
	fmt.Println("Time_taken(with go-routines):",time.Since(t1))
}
func main_concurrency(start int,end int, wg *sync.WaitGroup){
	for j:=start;j<end;j++{
		time.Sleep(2000)//sleeps for 2000ns
	}
	wg.Done()//removes the task from the group when the task is done
}

func no_concurrency(){
	t0:=time.Now()
	for i:=0;i<100;i++{
		time.Sleep(2000)//sleeps for 2000ns
	}
	fmt.Println()
	fmt.Println("Time_taken(No concurrency):",time.Since(t0))
}

func no_lock_used(){ //no mutex used to change the contents of arr1
	arr1:=[]int16{1,2,3}
	var wg=sync.WaitGroup{}
	var i int16
	for  i=4;i<=20;i++{
		wg.Add(1)
		go func(arr1 *[]int16,j int16){
			*arr1=append(*arr1,j)
			wg.Done()
	}(&arr1,i)
	}
	wg.Wait()
	fmt.Println("len of the array without using locks:",len(arr1))
	fmt.Println(arr1[:])
	
}

func lock_used(){
	arr1:=[]int16{1,2,3}
	var m=sync.Mutex{}//mutex used
	var wg=sync.WaitGroup{}
	var i int16
	for i=4; i <=20; i++ {
		wg.Add(1)
		go func(arr1 *[]int16,j int16){
			m.Lock()//used to lock other routines to access the shared memory
			*arr1=append(*arr1,j)//Lock() works like a Writer's lock
			m.Unlock()//unlocks the shared memory space
			wg.Done()
	}(&arr1,i)
	//can use Rlock() & Runlock() for implemnting reader-writer problem
	}
	wg.Wait()
	fmt.Println("len of the array using locks:",len(arr1))
	fmt.Println(arr1[:])
}

func channel_implementation(){
	var ch= make(chan int)//creates  and unbuffered channel
	go process(ch)//go-routine call
	fmt.Print("Channel content: ")
	for i:=range ch{
		fmt.Print(i)//reads the channel content
	}//Here the loop waits to read the content even though
	// process finishes executing.hence close the channel in process()
	fmt.Println()
}
func process(c chan int){// adds elements into the channel
	defer close(c)
	for i:=0;i<5;i++{
		c<-i
	}
	//close the channel so a deadlock wont be raised
}
func generic_func_call(){
var intSlice=[]int{1,2,3}
fmt.Println("Int sum:",generic(intSlice))
var float32Slice=[]float32{1.1,2.2,3.3}
fmt.Printf("Int sum:%.2f\n",generic(float32Slice))
}
func generic[T int | float32 | int16](slice []T)T{
var sum T//helps to avoid writing functions for diff datatype,though the function does the same thing
	for _,v :=range slice{
	sum+=v
} 
return sum
}
