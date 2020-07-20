package main

import (
	"log"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	// io.WriteString(response, "Hello, world!\n")
	helloBytes := []byte("Hello, world!\n")
	response.Write(helloBytes)
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	// TODO: check out what http.Handle does???

	log.Print("Starting the server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// If you put any code here, it will never run
	//
	// that's because the ListenAndServer line above
	// will run forever, and if it fails and stops running
	// for some reason, that log.Fatal function that "wraps"
	// ListenAndServe will crash the program
}
