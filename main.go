package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Processing request %+v\n", request)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(writer, "{\"greeting\":\"Hello World\"}")
}

func getEnv(name string, defaultPort string) string {
	envPort, exists := os.LookupEnv(name)
	if (!exists) {
			envPort = defaultPort
	}

	return envPort;
}