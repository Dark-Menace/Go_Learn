package main

import (
	"fmt"
	"time"
)

func Mat_Add(row,col int ,A[][]int64,B[][]int64){

var result=make([][]int64,row)
var t0=time.Now()
for i:=0;i<row;i++{
	result[i]=make([]int64,col)
	for j:=0;j<col;j++{	
		result[i][j]=A[i][j]+B[i][j]	
}
}
fmt.Println("Time for Addition(No go-routines):",time.Since(t0))

}
func Mat_Sub(row,col int ,A[][]int64,B[][]int64){

var result=make([][]int64,row)
var t0=time.Now()
for i:=0;i<row;i++{
	result[i]=make([]int64,col)
	for j:=0;j<col;j++{	
		result[i][j]=A[i][j]+B[i][j]	
}
}
fmt.Println("Time for Subtraction(No go-routines):",time.Since(t0))
}
func Mat_Scalar_Mul(row,col int,scalar_num int64,A [][]int64){
	var result=make([][]int64,row)
	var t0=time.Now()
	for i:=0;i<row;i++{
		result[i]=make([]int64,col)
			for j:=0;j<col;j++{
				result[i][j]=scalar_num*A[i][j]
			}
}

fmt.Println("Time for Scalar Multiplication(No go-routines):",time.Since(t0))
}
func Mat_Mul(row1,col1,col2 int,A [][]int64,B[][]int64){
	var result=make([][]int64,row1)
	var t0=time.Now()
	for i:=0;i<row1;i++{
		result[i]=make([]int64,col2)
			for j:=0;j<col2;j++{
				for k:=0;k<col1;k++{
				result[i][j]+=A[i][k]*B[k][j]
			}
		}
}

fmt.Println("Time for Matrix Multiplication(No go-routines):",time.Since(t0))
}
func Mat_Trasnpose(row ,col int,A[][]int64){
var result=make([][]int64,col)
var t0=time.Now()
for i:=0;i<col;i++{
	result[i]=make([]int64,row)
	for j:=0;j<row;j++{	
		result[i][j]=A[j][i]	
}
}
fmt.Println("Time for Transposing a Matrix(No go-routines):",time.Since(t0))
}