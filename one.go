package main

import "fmt"

func calc(){
	fmt.Println("This is another function")
}

func function3(){
	fmt.Println("This is function3")
}
func function4(){
	fmt.Println("This is function4")
}

func main(){
	name := "Max"
	fmt.Println("Hello", name)
	fmt.Println("Welcome")
	calc()
	function3()
	function4()
}
