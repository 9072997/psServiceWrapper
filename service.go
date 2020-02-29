package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kardianos/service"
)

// to event log
var logger service.Logger

func fatalErr(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go runScript(false)
	return nil
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	// load service name and description from service.json file
	configJSON, err := customFolder.Find("service.json")
	if err != nil {
		// logger not ready
		panic(err)
	}
	svcConfig := new(service.Config)
	err = json.Unmarshal(configJSON, svcConfig)
	if err != nil {
		// logger not ready
		panic(err)
	}

	// create service object
	prg := new(program)
	s, err := service.New(prg, svcConfig)
	if err != nil {
		// logger still not ready
		panic(err)
	}

	// "test" is a custom action implimented by us
	if len(os.Args) == 2 && os.Args[1] == "test" {
		// run with output enabled
		runScript(true)
		return
	}

	// allow the user to install the service using this binary
	if len(os.Args) >= 2 {
		err := service.Control(s, os.Args[1])
		if err != nil {
			fmt.Println("Valid actions:")
			for i := 0; i < len(service.ControlAction); i++ {
				fmt.Println(" ", service.ControlAction[i])
			}
			// "test" is implimented by us
			fmt.Println(" ", "test")

			fmt.Println(err)
		}
		return
	}

	// start the service
	logger, err = s.Logger(nil)
	if err != nil {
		panic(err)
	}
	err = s.Run()
	fatalErr(err)
}
