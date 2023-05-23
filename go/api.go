package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"oras.land/oras-go/v2/content"

	// "oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
	// "oras.land/oras-go/v2/registry/remote/auth"
	// "oras.land/oras-go/v2/registry/remote/retry"
)

type Data struct {
	Message string `json:"message"`
}

type oci struct {
	Registry string `json:"registry"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define your API routes
	router.HandleFunc("/api/data", getData).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"POST"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:1313"})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Print("method")
	fmt.Print(r.Method, "method")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var image oci
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&image)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	fmt.Print(image)

	Registry := image.Registry
	Repo := image.Repo
	Tag := image.Tag

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// // Fetch data from your data source or perform any other operations
	// // data := Data{
	// //     Message: "Hello from the API!",
	// // }

	// // reg := "gcr.io"
	// // reg := "docker.io" or "gcr.io"
	// // repository can be /library/nginx  or  /distroless/static

	repo, err := remote.NewRepository(Registry + Repo)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// // repo.Client = &auth.Client{
	// // 	Client: retry.DefaultClient,
	// // 	Cache:  auth.DefaultCache,
	// // 	Credential: auth.StaticCredential(reg, auth.Credential{
	// // 		Username: "izureki",
	// // 		Password: "5@bYk._(twRx7BW",
	// // 	}),
	// // }
	// // tag := "latest"
	descriptor, err := repo.Resolve(ctx, Tag)
	if err != nil {
		panic(err)
	}
	rc, err := repo.Fetch(ctx, descriptor)
	if err != nil {
		panic(err)
	}
	defer rc.Close() // don't forget to close
	pulledBlob, err := content.ReadAll(rc, descriptor)
	if err != nil {
		panic(err)
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(string(pulledBlob))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response content type and write the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
