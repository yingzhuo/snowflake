package main

import "fmt"

const (
	currentVersion   = "1.0.0"
	currentGoVersion = "1.13.1"
)

func printVersion() {
	fmt.Println("version:", currentVersion)
	fmt.Println("golang :", currentGoVersion)
}
