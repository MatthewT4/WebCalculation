package main

import "awesomeProject2/internal/http"

func main() {

	router := http.NewRouter()
	router.Start()
}
