package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

type QuantaRadexOneRequest struct {
	Prefix          uint16
	Command         uint16
	ExtensionLength uint16
	PacketNumber    uint16
	// Reserved0       uint16
	CheckSum0 uint16
	// SubCommand uint16
	// Reserved1  uint16
	// CheckSum1  uint16
}

func calcChecksum(fields ...uint16) uint16 {
	var sum uint16
	for _, n := range fields {
		sum += n
	}
	return (0xffff - sum) % 0xffff
}

func MakeQuantaRadexOneRequest(prefix, command, el, pn uint16) *QuantaRadexOneRequest {
	return &QuantaRadexOneRequest{
		prefix, command, el, pn, 0,
	}
}
func (qrr *QuantaRadexOneRequest) CalcChecksum0() {
	qrr.CheckSum0 = calcChecksum(qrr.Prefix, qrr.Command, qrr.ExtensionLength, qrr.PacketNumber)
}
func (qrr QuantaRadexOneRequest) Marshal() ([]byte, error) {
	result := make([]byte, 18)
	bo := binary.BigEndian
	bo.PutUint16(result[0:], qrr.Prefix)
	bo.PutUint16(result[2:], qrr.Command)
	bo.PutUint16(result[4:], qrr.ExtensionLength)
	bo.PutUint16(result[6:], qrr.PacketNumber)
	bo.PutUint16(result[10:], qrr.CheckSum0)
	bo.PutUint16(result[12:], 0x0100) // type
	bo.PutUint16(result[14:], 0x0c00) // reserved
	bo.PutUint16(result[16:], calcChecksum(0x0100, 0x0c00))
	return result, nil
}
func main() {
	fmt.Println("Starting...")
	// req := []byte{0x7B, 0xFF, 0x20, 0x00, 0x06, 0x00, 0x18, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x08, 0x0C, 0x00, 0xF3, 0xF7}
	// fmt.Printf("raw: %x\n", req)
	qrr := QuantaRadexOneRequest{
		0x7bff, 0x2000, 0x0600, 0x0000, 0x0000,
	}
	qrr.CalcChecksum0()
	encoded, err := qrr.Marshal()
	if err != nil {
		log.Fatalf("Can't encode qrr: %s", err)
	}
	c := &serial.Config{Name: "COM9", Baud: 9600, ReadTimeout: time.Millisecond * 100}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">: % X\n", encoded)
	s.Write(encoded)
	// scanner := bufio.NewScanner(s)
	// for scanner.Scan() {
	// 	fmt.Printf("% X", scanner.Bytes())
	// }
	buf := make([]byte, 1)
	var result []byte
	for {
		if n, err := s.Read(buf); err != nil || n == 0 {
			break
		}
		result = append(result, buf[0])
	}

	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("<: % X\n", result)
	// ctx := gousb.NewContext()
	// defer ctx.Close()

	// // Open any device with a given VID/PID using a convenience function.
	// dev, err := ctx.OpenDeviceWithVIDPID(0xabba, 0xa011) // Radex One
	// if err != nil {
	// 	log.Fatalf("Could not open a device: %v", err)
	// }
	// defer dev.Close()
	// // Claim the default interface using a convenience function.
	// // The default interface is always #0 alt #0 in the currently active
	// // config.
	// intf, done, err := dev.DefaultInterface()
	// if err != nil {
	// 	log.Fatalf("%s.DefaultInterface(): %v", dev, err)
	// }
	// defer done()
	// // Opening for read
	// usbIn, err := intf.InEndpoint(1)
	// if err != nil {
	// 	log.Fatalf("%s.InEndpoint(1): %v", intf, err)
	// }

	// // Requesting data
	// usbOut, err := intf.OutEndpoint(1)
	// if err != nil {
	// 	log.Fatalf("%s.OuteEndpoint(1): %v", intf, err)
	// }
	// fmt.Printf("qrr: % X\n", encoded)
	// written, err := usbOut.Write(encoded)
	// if err != nil {
	// 	log.Fatalf("Can't write: %s", err)
	// }
	// fmt.Printf("Written: %d of %d\n", written, len(encoded))
	// // Reading
	// buf := make([]byte, usbIn.Desc.MaxPacketSize)
	// readed, err := usbIn.Read(buf)
	// if err != nil {
	// 	log.Fatalf("Can't read. %s", err)
	// }
	// fmt.Printf("Readed: % X\n", readed)

	// Generate some data to write.
	// data := make([]byte, 5)
	// for i := range data {
	// 	data[i] = byte(i)
	// }

	// // Write data to the USB device.
	// numBytes, err := ep.Write(data)
	// if numBytes != 5 {
	// 	log.Fatalf("%s.Write([5]): only %d bytes written, returned error is %v", ep, numBytes, err)
	// }
	// fmt.Println("5 bytes successfully sent to the endpoint")
}
