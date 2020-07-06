package main

import (
	"github.com/jmckee46/myAppTest/docker-compose"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	err := dockerCompose.Stop()
	if err != nil {
		logger.Panic("myAppTest stop err", err.String())
	}
}
