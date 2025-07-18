package main

import (
	"net/http"
	"fmt"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🎉 CI/CD 成功部署! 版本: %s", os.Getenv("VERSION"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
