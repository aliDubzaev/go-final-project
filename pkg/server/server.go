package server

import (
	"fmt"
	"net/http"

	"github.com/aliDubzaev/go-final-project/pkg/api"
)

func Run() error {
	port := 7540
	webDir := "./web"

	api.Init()
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
