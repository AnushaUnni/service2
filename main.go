// client/main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AnushaUnni/service1/lib/calculate" // Importing the library from microservice
)

func main() {
	// Use the library directly
	a, b := 3, 5
	sum := calculate.Sum(a, b)
	fmt.Printf("Using library: The sum of %d and %d is %d\n", a, b, sum)

	// Call the microservice API
	url := fmt.Sprintf("http://localhost:8080/sum?a=%d&b=%d", a, b)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error calling microservice API: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response from the API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	fmt.Printf("Using API: %s\n", string(body))
}
