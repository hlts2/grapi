package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	return grapiserver.New().
		AddServers(
		// TODO
		).
		Serve()
}
