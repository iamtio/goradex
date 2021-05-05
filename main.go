package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	fmt.Println("Starting...")
	drr := NewDataRequest(0)
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
}
