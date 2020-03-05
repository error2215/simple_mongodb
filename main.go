package main

import (
	"github.com/error2215/simple_mongodb/server"
	_ "github.com/error2215/simple_mongodb/server/config"
)

func main() {
	server.Start()
}
