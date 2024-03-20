package main

import (
	"fmt"
)
type petrolEngine struct{//struct pattern
	car string;
	petrol uint8;
	fuel string;
	owenerinfo owner
}
type owner struct{
	name string
}
type dieselEngine struct{
	car string;
	diesel uint8;
	fuel string
	owenerinfo owner//struct within a struct
}
type combine interface{
	car_owner() string//interface to genarlize two struct with simialr fields
}
func basic(){//fundamentals
   //variable declaration and type testing
	var intNum int = 69
	fmt.Println(intNum)
	const rune_ele rune ='a'//rune->int32->unicode
    fmt.Println(rune_ele)//prints byte of the character(rune)
	var floatnum float32 = 69.69
	fmt.Println(floatnum + float32(intNum))
	result:=printme(intNum,floatnum)//shorthand rep
	//func call
	fmt.Printf("%v+%v=%v\n", floatnum, intNum,result)
	//switch call
	result_bool:=switch_case(1)
	fmt.Println("This switch result:",result_bool)

//array
var arr[3]int32//arr declared with 3 elements:[0,0,0] 
fmt.Println(&arr[0])//memory add of arr[0]
var arr1[3]int32=[3]int32{1,2,3}
fmt.Println(arr1)
arr2:=[...]int32{1,2,3}//capacity is 3 due to . operator
fmt.Println(arr2)
//slicing of array
test:=[]int32{1,2,3}
fmt.Println("The length:",len(test),"\tthe cap:",cap(test))
test=append(test,4)//appending
fmt.Println("The length:",len(test),"\tthe cap:",cap(test))
//capacity is memory allocated but cant be accessed without initialising entries 
fmt.Println(test[:])
pre_slice:=[...]int32{5,6}
test=append(test,pre_slice[:]...)//appending multiple items
fmt.Println(test)
}
func loop_and_maps(){//map usage and for loop
	// basic declaration of empty myMap->
//var myMap map[string]uint16=make(map[string]uint16)

var myMap=map[string]uint8{"Abhi":10,"Duke":12}

for name:= range myMap{//iterate over map (key values only)
	println("Name:",name)
}

fmt.Println(myMap["Abhi"])
var marks,ok=myMap["Dolly"]// marks: default value of unint8(key not found)
println(marks,"->",ok)// -->ok is false stating key-not found
delete(myMap,"Duke")//deletes the key-value pair
marks,ok=myMap["Duke"]// marks: default value of unint8(key not found)
println(marks,"->",ok)


//for-loop.....no while loop
for i,v:=range myMap{//access both key and vlaue in map
println(i,"->",v)
}
test:=[]int32{1,2,3,4,5,6}
for i,v:=range test{
	println("index:",i,"->value:",v)//i takes index,v takes value in  the array/slice
}
it:=0
for it<5{
	println("Hi")
	it+=1
}

for i:=1;i<3;i++{
	println(i)
}

word:="résume"
for i,v:=range word{
	println(i,"->",v)//here each char which is in utf-16(default is utf-8 for ascii value)
	//two indexes are occupied to store the byte value of the char
	//len of the string is legth of the byte and not char
}
}
func printme(a int, b float32) float32 {//if else-if ladder
	if a == 6 {
		return 0.00
	}else if(a==10){
		return 1.00
	}else{
	return (b + float32(a))
	}
}
func switch_case(a int)bool{//switch case
	switch(a){
	case 0: return true
	case 1:return true
	default:return false
	}
}
func struct_handle(){
	var car1 petrolEngine=petrolEngine{"Creta",16,"petrol",owner{"Arjun"}}
	car1.car="Audi"
	println(car1.car,car1.petrol,car1.owenerinfo.name)
	var car2 dieselEngine=dieselEngine{"Innova",12,"diesel",owner{"BK"}}
	println(car2.car,car2.diesel,car2.owenerinfo.name)
	println("Owner_info:")
	info(car1)
	println("Owner_info:")
	info(car2)
}
func (e petrolEngine) car_owner() string{
	return(e.car+"->"+e.owenerinfo.name)
}
func (d dieselEngine) car_owner() string{
	return(d.car+"->"+d.owenerinfo.name)
}
func info(c combine){
	println(c.car_owner())
}
func pointer_action(){
var p *int32=new(int32)//pointer declaration
fmt.Println("add_pointed:",p,"val_in_pointed:",*p)
//*p=10
var i int32=10
p=&i //p pointing to address of i
*p=5 //changing val of variable i thru p
println("value_pointed:",*p)

j:=10
k:=j
fmt.Println(&k,&j)//here copies val of j in k's address
//but if you create copies of slices,changes in one of the slice 
//affects other as they are essentially referring to pointers
}

func main() {
basic()
loop_and_maps()
struct_handle()
pointer_action()
no_concurrency()
routine_implement()
no_lock_used()
lock_used()
channel_implementation()
generic_func_call()
}
