package home

import (
	sharedTypes "app/pkg/shared_types"
	"fmt"
	"net/http"
)

const (
	DisplayNameKey sharedTypes.MetadataKey = "Home"
	UrlKey         sharedTypes.MetadataKey = "/"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Home!")
}
