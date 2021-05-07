package main

import (
	"fmt"
	"net/http"

	"github.com/sgreaves1/Agent-server/routes"
)

func main() {

	http.HandleFunc("/liveness", routes.LivenessRoute)

	err := http.ListenAndServe(fmt.Sprintf(":%d",7621), nil)

	if err != nil {
		fmt.Printf("Server failed to start: %d\n", err)
	}

	fmt.Printf("The application has started. Listening on port: %d\n", 7621)
	fmt.Println("Ctrl+C to exit.")

	fmt.Println("Closing server")
}


