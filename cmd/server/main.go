package main

import (
	"fmt"
	"github.com/AdrianPedriza/img-manager/pkg/rest"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/images", rest.UploadImageRequest)
	http.HandleFunc("/images/", rest.GetImageByIDRequest)

	fmt.Println("Server running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
