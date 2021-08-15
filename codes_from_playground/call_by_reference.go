package main

import (
	"fmt"
	"strings"
)

func joinstr(elements...string) string{
	return strings.Join(elements, "-")
}

func main(){
	fmt.Printf("Zero argument --- %s", joinstr())
	
	//pass a slice to a function
	
	element := []string{"Apollo", "Hercules", "Anaxagoras"}
	// not that element is used not elements 
	// element is argument
	// elements are paramaters
	fmt.Println(joinstr(element...))
}