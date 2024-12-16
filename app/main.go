package main

import (
	"fmt"
	"net/http"
)

func main() {
    
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
