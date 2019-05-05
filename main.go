package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server listening")

	http.HandleFunc("/", handlerHere)
	http.ListenAndServe(":8080", nil)
}

func handlerHere(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>This is the page</h1>"))
	w.Write([]byte("<h2>Read about this content</h2>"))
	w.Write([]byte("<p>venenatis tellus in metus vulputate eu scelerisque felis imperdiet proin fermentum leo vel orci porta non pulvinar neque laoreet suspendisse</p>"))
}
