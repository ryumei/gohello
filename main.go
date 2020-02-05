package main

import "fmt"

func greeting(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	fmt.Printf("%s\n", greeting("world"))
}
