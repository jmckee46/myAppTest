package main

import (
	"github.com/jmckee46/deployer/logger"
)

func main() {
	err := dockerCompose.Test()
	if err != nil {
		logger.Panic("myAppTest test err", err.String())
	}
}
