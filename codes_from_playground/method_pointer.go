package main

import (
	"fmt"
)

// here author represents struct -- data structure
type author struct {
	name string
	branch string
}

// here *a.branch = res.branch 
// so *a = res
// if *a = res then a = &res
// so *author refers to &res ???


func (a *author) show_1 (abranch string) {
	(*a).branch = abranch
}



func main() {
	
	// initialising values for data structure author
	
	res := author{
		name : "ABC",
		branch : "MECH",
	}
	fmt.Println("Branch name before:", res.branch)
	
	// calling show_1 -- string "ECE" from is passed to the parameter and *a points to the original res variable 
	res.show_1("ECE")
	
	// changes are made permanently
	fmt.Println("Branch in res after res.show_1() is:", res.branch)
	
	//using &res
	p:= &res
	p.show_1("CIVIL")
	fmt.Println("Branch in res after res.show_1() and using &res is:", res.branch)
	
	
}
