package main

import "fmt"


const age = 32
func main(){
	// const name = "golang"

	// name = "javascript" // can't re-declare once value its assign to value 

	const (
		port = 5000
		name = "Nitesh"
	)

	fmt.Println(port , name)

	fmt.Println(age)
}
