package main

import (
	"testing"
)

// GOTTA BE UPPERCASE T
//
// GET THAT TEEEEEEE
func Test_helloHandler(tea *testing.T) {
	// TODO: what is recorder? what is req?
	//
	// Check out this library to find out!
	//
	// https://pkg.go.dev/net/http/httptest?tab=doc
	helloHandler(recorder, req)
	tea.Fatal("Gotta Have That Uppercase T")
}
