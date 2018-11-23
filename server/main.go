package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"time-protocol/config"
	"time-protocol/entities"
)

// RFC 868 Time Protocol implementation
// https://tools.ietf.org/html/rfc868
func main() {
	portFlag := flag.Int("p", config.DefaultPort, "listen port")
	flag.Parse()

	s, err := entities.NewServer()
	if err != nil {
		log.Fatalln("New server init failed:", err.Error())
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%d", config.Host, *portFlag)

	s.Listen(config.NetType, address)
}
