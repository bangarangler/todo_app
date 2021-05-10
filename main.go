package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

var port = goDotEnvVar("PORT")

func handleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, Golang ❤️")
}

func main() {
	if port == "" {
		port = strconv.Itoa(8080)
	}

	http.HandleFunc("/", handleFunc)

	println(port)
	log.Printf("todo_app running on http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
