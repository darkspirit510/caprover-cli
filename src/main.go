package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello World"
}

func DoAction(params []string) string {
	//fmt.Println(len(params))

	return "Usage: "
}