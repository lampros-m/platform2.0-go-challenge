package main

import (
	server "gwi/platform2.0-go-challenge/api"
)

func main() {
	server := server.NewApplicationServer()
	server.Setup()
	defer server.App.Client.Close()

	server.Run()
}
