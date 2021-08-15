
// Go program to illustrate the
// use of keywords
package main
import "fmt"
  
// Here, package, import, func,
// var are keywords
func main() {
  
// Here, a is a valid identifier
var a = "GeeksforGeeks" 
  
fmt.Println(a)
  
// Here, the default is an
// illegal identifier and
// compiler will throw an error
// var default = "GFG"
}


25 KEYWORDS IN GO:

	break
	case
	chan
	const
	continue
	default
	defer
	else
	fallthrough
	for
	func
	go
	goto
	if
	import
	interface
	map
	package
	range
	return
	select
	struct
	switch
	type
	var


CODE:

	package main

	import "fmt"
	var name = "Goku"
	func main(){
		fmt.Printf("1+1=%d", 1+1)
		fmt.Printf("name is %s", name)
	}