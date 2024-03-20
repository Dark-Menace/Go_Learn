package main
import("fmt")

func main(){
fmt.Println("Welcome to the matrix module....")
fmt.Println("1.Matrix Addition")
fmt.Println("2.Matrix Subtraction")
fmt.Println("3.Matrix Scalar Multiplication")
fmt.Println("4.Matrix Multiplication")
fmt.Println("5.Matrix Transpose")
fmt.Println("6.Matrix Inverse")
fmt.Println("Choose your choice:")
var choice int
fmt.Scan(&choice)
switch choice{
case 1:
	Mat_Add_Con()
case 2:
	Mat_Sub_Con()
case 3:
	Mat_Scalar_Mul_Con()
case 4:
	Mat_Mul_Con()
case 5:
	Mat_Trasnpose_Con()
case 6:
	Mat_Inverse()
default:
	fmt.Println("Invalid choice.")

}


}
