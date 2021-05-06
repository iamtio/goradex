package main

import (
	"fmt"
	"log"
	"time"

	"github.com/iamtio/goradex/radexone"
	"github.com/tarm/serial"
)

func main() {
	fmt.Println("Starting...")
	drr := radexone.NewDataRequest(0)
	encoded := drr.Marshal()

	c := &serial.Config{Name: "COM9", Baud: 9600, ReadTimeout: time.Millisecond * 100}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">: % X\n", encoded)
	s.Write(encoded)
	buf := make([]byte, 1)
	var result []byte
	for {
		if n, err := s.Read(buf); err != nil || n == 0 {
			break
		}
		result = append(result, buf[0])
	}
	fmt.Printf("<: % X\n", result)
	resp := radexone.DataReadResponse{}
	resp.Unmarshal(result)
	fmt.Printf("CPM: %d, Ambient: %d, Accumulated: %d", resp.CPM, resp.Ambient, resp.Accumulated)
}
