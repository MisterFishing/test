package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", main)
}

func main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go!");
}
