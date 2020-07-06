package main

import (
	"github.com/jmckee46/myAppTest/docker"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	err := docker.Build()
	if err != nil {
		logger.Panic("myAppTest build err", err.String())
	}
}
