package main

import "fmt"

func main() {
	fmt.Println("Enter your name & age")
	var name string
	var age int
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("your name is %s & city is %d", name, age)
}
