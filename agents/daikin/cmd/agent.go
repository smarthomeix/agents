package main

import (
	"github.com/smarthomeix/agents/agents/daikin/internal/service"
	"github.com/smarthomeix/agents/pkg/server"
)

func main() {
	server.NewServer(service.NewDaikinService())
}
