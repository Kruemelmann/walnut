package main

import (
	"embed"
)

//go:embed static/ui/dist
var FrontendFS embed.FS

func main() {
	//TODO read with viper
	var (
		host = "localhost"
		port = "8000"
	)

	srv := NewServer(host, port)
	srv.Start()
}
