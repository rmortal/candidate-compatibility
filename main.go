package main

import (
	"candidate-compatibility/api"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := "9000"
	fmt.Printf("Starting up on http://localhost:%s\n", port)

	// Define Routes
	r := chi.NewRouter()
	r.Mount("/analyze-compatibility", api.AnalyeController{}.AnalyzeRoutes())
	fmt.Println("Mounted routes")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
