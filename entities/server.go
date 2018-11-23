package entities

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"time-protocol/config"
)

type timeProtocolServer struct{}

func NewServer() (*timeProtocolServer, error) {
	return &timeProtocolServer{}, nil
}

// Listen - handle TCP requests and handle it with HandleRequest method.
func (s *timeProtocolServer) Listen(network, address string) {
	ln, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("Error listening:", err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Listening on " + address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Error on accept: ", err.Error())
			continue
		}

		go s.HandleRequest(conn)
	}

}

// HandleRequest - handle TCP request, get current time
// and if didn't get errors writes it to response and close connection.
func (s *timeProtocolServer) HandleRequest(conn net.Conn) {
	log.Println("Get request")

	t, err := GetTime()
	if err == nil {
		conn.Write(t)
	}

	conn.Close()
}

// GetTime - get time.Now() and add time shift with of seconds
// from 1 January 1900 to 1 January 1970.
// It responses with time as a 32 bit binary integer.
func GetTime() (out []byte, err error) {
	buf := new(bytes.Buffer)

	timeData := int32((time.Now().Unix()) + config.TimeShift)

	err = binary.Write(buf, binary.BigEndian, timeData)
	if err != nil {
		return nil, err
	}

	out = buf.Bytes()
	buf.Reset()

	return out, err
}
