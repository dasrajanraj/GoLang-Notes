package main

import (
	"fmt"
)

func main() {
	// declaring syntax
	// variableName := [<size>]<type>{<initialValues>}
	salesEachMonth := [12]int{1, 12, 11, 31, 12, 54, 7, 1, 14, 86, 8, 8}
	fmt.Printf("Sales in each month %v", salesEachMonth)
}
