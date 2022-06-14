package main

import(
	"fmt"
	"reflect"
)

func main(){
	
	/*
	var num1 int = 1234567890
	var num2 float64 = 1.23456789

	fmt.Println(num1)
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(num2)
	fmt.Println(reflect.TypeOf(num2))

	var string_a string = "Hello,World!"
	string_b := "Hello,World!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))
	*/



	var a int
	a = 1

	var b int = 10

	ans_a_b := a < b

	fmt.Println(ans_a_b)
	fmt.Println(reflect.TypeOf(ans_a_b))


}