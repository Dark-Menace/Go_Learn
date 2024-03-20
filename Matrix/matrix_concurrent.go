package main

import("fmt";"sync";"time")

type datatype interface{
	int64|float64
}

func Mat_Add_Con(){
var wg =sync.WaitGroup{}
row,col:=Enter_RowCol()
A:=Make_Matrix(row,col)
B:=Make_Matrix(row,col)
var result=make([][]int64,row)
var t0=time.Now()
for i:=0;i<row;i++{
	result[i]=make([]int64,col)
	go func(i int){
		defer wg.Done()
		wg.Add(1)
	for j:=0;j<col;j++{	
		result[i][j]=A[i][j]+B[i][j]	
}
}(i)
}
wg.Wait()
fmt.Println("Time for Addition(using go-routines):",time.Since(t0))
Mat_Add(row,col,A,B)
//fmt.Println("The Matrix A:")
//Print_Mat(A,row,col)
//fmt.Println("The Matrix B:")
//Print_Mat(B,row,col)
//fmt.Println("The result Matrix :")
//Print_Mat(result,row,col)
}

func Mat_Sub_Con(){
var wg =sync.WaitGroup{}
row,col:=Enter_RowCol()
A:=Make_Matrix(row,col)
B:=Make_Matrix(row,col)
var result=make([][]int64,row)
var t0=time.Now()
for i:=0;i<row;i++{
	result[i]=make([]int64,col)
	go func(i int){
		defer wg.Done()
		wg.Add(1)
	for j:=0;j<col;j++{	
		result[i][j]=A[i][j]-B[i][j]	
}
}(i)
}
wg.Wait()
fmt.Println("Time for Subtraction(using go-routines):",time.Since(t0))
Mat_Sub(row,col,A,B)
}
func Mat_Scalar_Mul_Con(){

	var wg=sync.WaitGroup{}
	row,col:=Enter_RowCol()
	A:=Make_Matrix(row,col)
	var result=make([][]int64,row)
	var scalar_num int64
	fmt.Println("Enter an integer to perform Scalar Mutiplication")
	fmt.Scan(&scalar_num)
	var t0=time.Now()
	for i:=0;i<row;i++{
		result[i]=make([]int64,col)
		go func(i int){
			defer wg.Done()
			wg.Add(1)
			for j:=0;j<col;j++{
				result[i][j]=scalar_num*A[i][j]
			}
		}(i)
}
wg.Wait()
fmt.Println("Time for Scalar Multiplication(using go-routines):",time.Since(t0))
Mat_Scalar_Mul(row,col,scalar_num,A)

}
func Mat_Mul_Con(){
var wg =sync.WaitGroup{}
row1,col1:=Enter_RowCol()
row2,col2:=Enter_RowCol()
A:=Make_Matrix(row1,col1)
B:=Make_Matrix(row2,col2)
if(col1!=row2){
	panic("Multiplication condition cannot be satisfied:col1 != row2")
}
var result=make([][]int64,row1)
var t0=time.Now()
for i:=0;i<row1;i++{
	result[i]=make([]int64,col2)
	go func(i int){
		defer wg.Done()
		wg.Add(1)
	for j:=0;j<col2;j++{	
		for k:=0;k<col1;k++{
			result[i][j]+=A[i][k]*B[k][j]
		}	
}
}(i)
}
wg.Wait()
fmt.Println("Time for Matrix Multiplication(using go-routines):",time.Since(t0))
Mat_Mul(row1,col1,col2,A,B)
}
func Mat_Trasnpose_Con()  {
	var wg=sync.WaitGroup{}
	row,col:=Enter_RowCol()
	A:=Make_Matrix(row,col)
	var result=make([][]int64,col)
	var t0=time.Now()
	for i:=0;i<col;i++{
	result[i]=make([]int64,row)
	go func(i int){
		defer wg.Done()
		wg.Add(1)
	for j:=0;j<row;j++{	
		result[i][j]=A[j][i]	
}
}(i)
}
wg.Wait()
fmt.Println("Time for Transposing a Matrix(using go-routines):",time.Since(t0))
Mat_Trasnpose(row,col,A)

}
func Mat_Inverse()  {
	
}