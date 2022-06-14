package main

import(
	"fmt"
	"reflect"
)


func main(){
	// 変数 := [要素数]データ型{データ1, データ2, ・・・}
	a:= [...]string{"A","B","C"}

	//b:= [3]string{"D","E","F"}

	ab:=[2][3]int{{1,2,3},{4,5,6}}

	fmt.Println(reflect.TypeOf(ab))

	for i:=0; i < len(a); i++{
		fmt.Println(a[i])
	}
}