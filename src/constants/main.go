package main

import (
	"fmt"
)

func main() {

	const port int = 3300
	fmt.Printf("Application is running in port %v\n", port)

	//Error while assigning value during runtime
	// const PI float32 = math.Sin(20) // throws an error

	const pi = 10
	var i int64 = 10
	sum := pi + i
	fmt.Printf(" %v , %T\n", sum, sum) // 20 , int64
}
