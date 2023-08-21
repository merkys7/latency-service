//go:build linux
// +build linux

package main

import (
	"github.com/g-portal/metadata-server/pkg/server"
)

func main() {
	server.StartServer()
}
