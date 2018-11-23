package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"time-protocol/config"
	"time-protocol/entities"
)

func main() {
	var host, port string

	flag.Parse()

	// Get params
	if flag.Arg(0) == "" {
		host = config.Host
	} else {
		host = flag.Arg(0)
	}

	if flag.Arg(1) == "" {
		port = strconv.Itoa(config.DefaultPort)
	} else {
		port = flag.Arg(1)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	c, err := entities.NewClient(address)
	if err != nil {
		log.Fatalln("NewClient failed:", err.Error())
	}

	// Get request
	tm, err := c.RequestTime()
	if err != nil {
		log.Fatalln("RequestTime failed:", err.Error())
	}

	fmt.Println(tm)
}
