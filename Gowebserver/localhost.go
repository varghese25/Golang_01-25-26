package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Go Server, localhost!")
		fmt.Fprintln(w, "Step 1 Initialize the Go module")
		fmt.Fprintln(w, "Step 2 Open your terminal in E:\\Golang_01-25-26\\Gowebserver: go mod init gowebserver")
		fmt.Fprintln(w, "Step 3 This will create a go.mod file, which might look like:")
		fmt.Fprintln(w, "module gowebserver")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "go 1.21")
		fmt.Fprintln(w, "go run .")
		fmt.Fprintln(w, "Step 4 Open your browser and navigate to http://localhost:8080 to see the output.")
		fmt.Fprintln(w, "Step 5 To stop the server, go back to your terminal and press Ctrl + C.")

    })

    http.ListenAndServe(":8080", nil)
}