package main

import (
    "encoding/json"
    "log"
    "net/http"
    "context"
	// "fmt"

    "github.com/gorilla/mux"

	"oras.land/oras-go/v2/content"

	// "oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/retry"
)

type Data struct {
    Message string `json:"message"`
}

func main() {
    // Create a new router
    router := mux.NewRouter()

    // Define your API routes
    router.HandleFunc("/api/data", getData).Methods("GET")

    // Start the HTTP server
    log.Fatal(http.ListenAndServe(":8080", router))
}

func getData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Fetch data from your data source or perform any other operations
    // data := Data{
    //     Message: "Hello from the API!",
    // }

    reg := "docker.io"
	repo, err := remote.NewRepository(reg + "/library/hello-world")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	repo.Client = &auth.Client{
		Client: retry.DefaultClient,
		Cache:  auth.DefaultCache,
		Credential: auth.StaticCredential(reg, auth.Credential{
			Username: "vasudevrani",
			Password: "5@bYk._(twRx7BW",
		}),
	}
	tag := "latest"
	descriptor, err := repo.Resolve(ctx, tag)
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

	// fmt.Println(string(pulledBlob))

    // Convert the data to JSON
    jsonData, err := json.Marshal(string(pulledBlob))
    // fmt.Print(jsonData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set the response content type and write the JSON data
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
