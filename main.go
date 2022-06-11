package main

import (
	server "podcasts/server"
)

func main() {

	useDummyData := true

	server.StartServer(useDummyData)
}
