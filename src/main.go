package main

import "github.com/ThalesLoreto/ogavirt-backend/server"

func main() {
	server := server.NewServer(8000)
	server.Start()
}
