package main

import "fmt"

const (
	a = iota
	// resource
	RESOURCE = "https://management.azure.com/"
	// LOGIN_API is a thing
	LoginAPI = "https://login.microsoftonline.com/"
	B
	c
	d
)

func main() {
	fmt.Println("Hello, 世界")

}

type s struct {
	a int
}

func (a s) doathing() {
}

func (a *s) doanotherthing() {

}
