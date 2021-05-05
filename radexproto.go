package main

type RadexOneRequest interface {
	Marshal() []byte
}
type BaseRequest struct {
	Prefix          uint16 `pos:"0"`
	Command         uint16 `pos:"2"`
	ExtensionLength uint16 `pos:"4"`
	PacketNumber    uint16 `pos:"6"`
	Reserved0       uint16 `pos:"8"`
	CheckSum0       uint16 `pos:"10"`
}
type DataReadRequest struct {
	BaseRequest
	RequestType uint16 `pos:"0"`
	Reserved1   uint16 `pos:"2"`
	CheckSum1   uint16 `pos:"4"`
}

func NewDataRequest(packetNum uint16) DataReadRequest {
	drr := DataReadRequest{
		BaseRequest: BaseRequest{
			Prefix:          0x7bff,
			Command:         0x2000,
			ExtensionLength: LEWord(0x0006),
			PacketNumber:    LEWord(packetNum),
			Reserved0:       0,
			CheckSum0:       0,
		},
		RequestType: 0x0008,
		Reserved1:   0x0c00,
		CheckSum1:   0,
	}
	drr.CheckSum0 = CalcChecksum(drr.Prefix, drr.Command, drr.ExtensionLength, drr.PacketNumber, drr.Reserved0)
	drr.CheckSum1 = CalcChecksum(drr.RequestType, drr.Reserved1)
	return drr
}

func (drr DataReadRequest) Marshal() []byte {
	var buf []byte
	buf = append(marshalStruct(drr.BaseRequest)[:], marshalStruct(drr)[:]...)

	return buf
}
