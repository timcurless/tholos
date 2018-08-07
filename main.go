package main

import (
	"github.com/timcurless/tholos/service"
)

// TODO: Add an auth method. K8s secret? AppRole?

func main() {
	service.StartWebServer("8081")
}
