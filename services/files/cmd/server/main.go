package main

import (
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal"
)

func main() {
	r := internal.NewApiRouter()

	r.SetTrustedProxies([]string{"reverse-proxy"})

	r.GET("/healthcheck")

	r.Run()
}
