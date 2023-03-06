package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("RUNNING!")
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
