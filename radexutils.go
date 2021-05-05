package main

import (
	"encoding/binary"
	"log"
	"reflect"
	"strconv"
)

func CalcChecksum(fields ...uint16) uint16 {
	var sum uint16
	for _, n := range fields {
		sum += n
	}
	return (0xffff - sum) % 0xffff
}
func LEWord(value uint16) uint16 {
	return value<<8 | value>>8 //0xAABB => 0xBBAA
}
func marshalStruct(i interface{}) []byte {
	drrV := reflect.ValueOf(i)
	drrT := drrV.Type()

	maxPos := 0
	for i := 0; i < drrV.NumField(); i++ {
		pos := drrT.Field(i).Tag.Get("pos")
		if pos == "" {
			continue
		}
		posInt, _ := strconv.Atoi(pos)
		if posInt > maxPos {
			maxPos = posInt
		}
	}
	buf := make([]byte, maxPos+2)
	for i := 0; i < drrV.NumField(); i++ {
		typeField := drrT.Field(i)
		pos := typeField.Tag.Get("pos")
		if pos == "" {
			continue
		}
		posInt, err := strconv.Atoi(pos)

		if err != nil {
			log.Printf("Can't parse pos for: %s\n", typeField.Name)
			continue
		}
		binary.BigEndian.PutUint16(buf[posInt:], uint16(drrV.Field(i).Uint()))
	}
	return buf
}
func unmarshalStruct(packet []byte, i interface{}) {
	drrV := reflect.ValueOf(i).Elem()
	drrT := drrV.Type()
	for i := 0; i < drrV.NumField(); i++ {
		pos := drrT.Field(i).Tag.Get("pos")
		le := drrT.Field(i).Tag.Get("le")
		if pos == "" {
			continue
		}
		posInt, err := strconv.Atoi(pos)
		if err != nil {
			log.Printf("Can't parse pos for: %s\n", drrT.Field(i).Name)
			continue
		}
		//TODO: Check buf len
		var value uint16
		if le != "" {
			value = binary.BigEndian.Uint16(packet[posInt : posInt+2])
		} else {
			value = binary.LittleEndian.Uint16(packet[posInt : posInt+2])
		}

		drrV.Field(i).SetUint(uint64(value))
	}
}
