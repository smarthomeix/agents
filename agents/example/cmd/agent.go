package main

import (
	"github.com/smarthomeix/agents/agents/example/internal/service"
	"github.com/smarthomeix/agents/pkg/server"
)

func main() {
	server.NewServer(service.NewExampleService())
}
