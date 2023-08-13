package main

import (
	"fmt"
)

func main() {

	DataTypes()
	Arrays()
}

func DataTypes() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}
func Arrays() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := []string{"rattapon", "rattana"}

	fmt.Println(arr1)
	fmt.Println(arr2[2])
	fmt.Println(len(arr3))

}
func Arrays2() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := []string{"rattapon", "rattana"}

	fmt.Println(arr1)
	fmt.Println(arr2[2])
	fmt.Println(len(arr3))

}
