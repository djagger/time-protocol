package entities

import (
	"encoding/binary"
	"net"
	"time"

	"time-protocol/config"
)

var epoch = parseTime("1900-01-01 00:00:00")

func parseTime(t string) time.Time {
	tm, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		panic(err)
	}

	return tm
}

type timeProtocolClient struct {
	conn *net.TCPConn
}

// NewClient - establishes a connection by address
// and returns timeProtocolClient object with TCP connection.
func NewClient(address string) (*timeProtocolClient, error) {
	tcpAddr, err := net.ResolveTCPAddr(config.NetType, address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP(config.NetType, nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	return &timeProtocolClient{conn}, nil
}

// RequestTime - write response with standard Unix timestamp.
func (c *timeProtocolClient) RequestTime() (int64, error) {
	_, err := c.conn.Write([]byte{})
	if err != nil {
		return 0, err
	}

	var n uint32
	err = binary.Read(c.conn, binary.BigEndian, &n)
	if err != nil {
		return 0, err
	}

	tm := epoch.Add(time.Duration(n) * time.Second).Unix()

	return tm, nil
}
