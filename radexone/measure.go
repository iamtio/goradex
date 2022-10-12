package radexone

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

type Measure struct {
	CPM         int
	Ambient     float64
	Accumulated float64
}
type MeasureHandler struct {
	SerialPort string
	SerialBaud int
}

func (m *MeasureHandler) GetValues() Measure {
	drr := NewDataRequest(0)
	encoded := drr.Marshal()

	c := &serial.Config{Name: m.SerialPort, Baud: m.SerialBaud, ReadTimeout: time.Millisecond * 100}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	// fmt.Printf(">: % X\n", encoded)
	s.Write(encoded)
	buf := make([]byte, 1)
	var result []byte
	for {
		if n, err := s.Read(buf); err != nil || n == 0 {
			break
		}
		result = append(result, buf[0])
	}
	// fmt.Printf("<: % X\n", result)
	resp := DataReadResponse{}
	resp.Unmarshal(result)
	return Measure{
		CPM:         int(resp.CPM),
		Ambient:     float64(resp.Ambient) / 100,
		Accumulated: float64(resp.Accumulated) / 100,
	}
}
