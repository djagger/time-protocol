package tests

import (
	"fmt"
	"testing"

	"time-protocol/config"
	"time-protocol/entities"
)

func Test_NewClient(t *testing.T) {
	s, err := entities.NewServer()
	if err != nil {
		t.Error("NewServer failed:", err.Error())
	}

	address := fmt.Sprintf("%s:%d", config.Host, config.DefaultPort)

	go s.Listen(config.NetType, address)

	c, err := entities.NewClient(address)
	if err != nil {
		t.Error("NewClient failed:", err.Error())
	}

	_, err = c.RequestTime()
	if err != nil {
		t.Error("RequestTime failed:", err.Error())
	}
}
