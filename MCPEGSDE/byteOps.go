package mcpegsde

import (
	"encoding/binary"
)

// I just copied this from one of my other projects. I don't know if it's the best way to do this. It's probably not.

type uiint16 interface {
	uint16 | int16
}

type uiint8 interface {
	uint8 | int8
}

type uiint64 interface {
	uint64 | int64
}

type uiint32 interface {
	uint32 | int32
}

func Int32ToBytes[T uiint32](value T) []byte {
	var buffer = make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, uint32(value))
	return buffer
}

func Int16ToBytes[T uiint16](value T) []byte {
	var buffer = make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, uint16(value))
	return buffer
}

func Int64ToBytes[T uiint64](value T) []byte {
	var buffer = make([]byte, 8)
	binary.LittleEndian.PutUint64(buffer, uint64(value))
	return buffer
}

func Int8ToBytes[T uiint8](value T) []byte {
	var buffer = make([]byte, 1)
	buffer[0] = uint8(value)
	return buffer
}

func BytesToInt32[T uiint32](buffer []byte) T {
	return T(binary.LittleEndian.Uint32(buffer))
}

func BytesToInt16[T uiint16](buffer []byte) T {
	return T(binary.LittleEndian.Uint16(buffer))
}

func BytesToInt64[T uiint64](buffer []byte) T {
	return T(binary.LittleEndian.Uint64(buffer))
}

func BytesToInt8[T uiint8](buffer []byte) T {
	return T(buffer[0])
}
