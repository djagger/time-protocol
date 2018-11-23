package tests

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"time"

	"time-protocol/config"
	"time-protocol/entities"
)

func Test_GetTime(t *testing.T) {
	tt, err := entities.GetTime()
	if err != nil {
		fmt.Println(err.Error())
	}

	b := make([]byte, 4)
	now := time.Now().Unix() + config.TimeShift
	binary.BigEndian.PutUint32(b, uint32(now))

	if !bytes.Equal(tt, b) {
		t.Errorf("Incorrect time, got: %v, want: %v.", tt, b)
	}
}

// Method Listen tested in Test_NewClient test.
func Test_Listen(t *testing.T) {

}
