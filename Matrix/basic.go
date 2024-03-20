package main
import("fmt";"math/rand")
func Enter_RowCol()(int,int){
fmt.Println("Enter the number of rows:")
var row int
fmt.Scan(&row)
fmt.Println("Enter the number of columns:")
var col int
fmt.Scan(&col)
if (row<0 ||col<0){
	panic("Negative value for row or column!")
}
return row,col
}
func Make_Matrix(row,col int )([][]int64){
var array=make([][]int64, row)
for i:=0;i<row;i++{
	array[i]=make([]int64,col)
	for j:=0;j<col;j++{
		array[i][j]=rand.Int63n(101)
	}
}
return array
}
func Print_Mat[T int64|float64](mat [][]T,row,col int){
	for i:=0;i<row;i++{
	for j:=0;j<col;j++{
		fmt.Print(mat[i][j],"  ")
	}
	fmt.Println()
}
}
