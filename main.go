// client/main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AnushaUnni/service1/lib/calculate" // Importing the library from microservice
)

func main() {
	// Use the library directly
	a, b := 3, 5
	sum := calculate.Sum(a, b)
	fmt.Printf("Using library: The sum of %d and %d is %d\n", a, b, sum)

	// Call the microservice API
	service1API := os.Getenv("SERVICE1_API_URL")
	if service1API == "" {
		log.Fatal("SERVICE1_API_URL environment variable is not set")
	}
	url := fmt.Sprintf("%s/sum?a=%d&b=%d", service1API, a, b)
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
