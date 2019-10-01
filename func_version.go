package main

import "fmt"

const (
	currentVersion   = "1.0.0"
	currentGoVersion = "go1.13.1"
)

func printVersion() {
	fmt.Println("Version   :", currentVersion)
	fmt.Println("Go version:", currentGoVersion)
}
