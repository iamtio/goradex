package radexone

import (
	"log"
)

type Request interface {
	Marshal() []byte
}
type PacketHeader struct {
	Prefix          uint16 `pos:"0"`
	Command         uint16 `pos:"2"`
	ExtensionLength uint16 `pos:"4" le:"1"`
	PacketNumber    uint16 `pos:"6" le:"1"`
	Reserved0       uint16 `pos:"8"`
	CheckSum0       uint16 `pos:"10"`
}
type DataReadRequest struct {
	PacketHeader
	RequestType uint16 `pos:"0"`
	Reserved1   uint16 `pos:"2"`
	CheckSum1   uint16 `pos:"4"`
}

type DataReadResponse struct {
	PacketHeader
	RequestType uint16 `pos:"0"`
	Reserved1   uint16 `pos:"4"`
	Ambient     uint16 `pos:"8" le:"1"`
	Accumulated uint16 `pos:"12" le:"1"`
	CPM         uint16 `pos:"16" le:"1"`
	CheckSum1   uint16 `pos:"20"`
}

func NewDataRequest(packetNum uint16) DataReadRequest {
	drr := DataReadRequest{
		PacketHeader: PacketHeader{
			Prefix:          0x7bff,
			Command:         0x2000,
			ExtensionLength: 0x0006,
			PacketNumber:    packetNum,
			Reserved0:       0,
			CheckSum0:       0,
		},
		RequestType: 0x0008,
		Reserved1:   0x0c00,
		CheckSum1:   0,
	}
	drr.CheckSum0 = CalcChecksum(drr.Prefix, drr.Command, LEWord(drr.ExtensionLength), LEWord(drr.PacketNumber), drr.Reserved0)
	drr.CheckSum1 = CalcChecksum(drr.RequestType, drr.Reserved1)
	return drr
}

func (drr DataReadRequest) Marshal() []byte {
	var buf []byte
	buf = append(marshalStruct(drr.PacketHeader)[:], marshalStruct(drr)[:]...)
	return buf
}

func (drr *DataReadResponse) Unmarshal(packet []byte) {
	unmarshalStruct(packet, &drr.PacketHeader)
	if cs := CalcChecksum(drr.Prefix, drr.Command, LEWord(drr.ExtensionLength), LEWord(drr.PacketNumber), drr.Reserved0); cs != drr.CheckSum0 {
		log.Fatalf("Checksum0 mismatch: %x != %x\n", cs, drr.CheckSum0)
	}
	unmarshalStruct(packet[12:], drr)
	if cs := CalcChecksum(drr.RequestType, drr.Reserved1, LEWord(drr.Ambient), LEWord(drr.Accumulated), LEWord(drr.CPM)); cs != drr.CheckSum1 {
		log.Fatalf("Checksum1 mismatch: %x != %x\n", cs, drr.CheckSum1)
	}
}
