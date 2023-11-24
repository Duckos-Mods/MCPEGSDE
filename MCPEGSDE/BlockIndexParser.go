package mcpegsde

import (
	"reflect"
	"unsafe"
)

const (
	Palleted1  = 1  // 32 blocks per word
	Palleted2  = 2  // 16 blocks per word
	Palleted3  = 3  // 10 blocks per word and 2 bits of padding per word i think each byte has 2 bits of padding at the start.
	Palleted4  = 4  // 8 blocks per word
	Palleted5  = 5  // 6 blocks per word and 2 bits of padding per word.
	Palleted6  = 6  // 5 blocks per word and 2 bits of padding per word.
	Palleted8  = 8  // 4 blocks per word
	Palleted16 = 16 // 2 blocks per word
)

// Word is a uint32

func calculateBlocksPerWord(bitsPerBlock uint8) int {
	return int(32 / bitsPerBlock)
}

func getMask(blocksPerWord uint16) uint16 {
	// Calculate the number of bits to set in the mask
	numBitsSet := (1 << blocksPerWord) - 1

	// Generate the mask by setting the desired number of bits
	mask := uint16(numBitsSet)

	return mask
}

func maskWordToBytes(word uint32) (byte, byte, byte, byte) {
	var LeftByte = byte(word >> 24)
	var MidLeft = byte(word >> 16)
	var MidRight = byte(word >> 8)
	var RightByte = byte(word)
	return LeftByte, MidLeft, MidRight, RightByte
}

func maskBytesToWord(LeftByte, MidLeft, MidRight, RightByte byte) uint32 {
	var returnWord uint32
	returnWord |= uint32(LeftByte) << 24
	returnWord |= uint32(MidLeft) << 16
	returnWord |= uint32(MidRight) << 8
	returnWord |= uint32(RightByte)
	return returnWord
}

func unpackBlocksPN(blocksPerWord int, rawData []byte) []uint16 {
	var returnData []uint16
	for i := 0; i < len(rawData)/4; i++ {
		var word = maskBytesToWord(rawData[i*4], rawData[i*4+1], rawData[i*4+2], rawData[i*4+3])
		for b := 0; b < blocksPerWord; b++ {
			returnData = append(returnData, uint16(word>>b&uint32(getMask(uint16(blocksPerWord)))))
		}
	}
	return returnData
}

// This function wraps some calls to internal functions so you dont have to do it yourself
func GetBlocksFromBytes(SubChunk *LevelSubChunk, rawData []byte) {
	SubChunk.BlkData = unpackBlocksPN(calculateBlocksPerWord(SubChunk.BitsPerBlock), rawData)
}

func packBlocksPN(SubChunk *LevelSubChunk, blocksPerWord int) []byte {

	var ReturnPallet = make([]uint32, len(SubChunk.BlkData)/2)
	for i := 0; i < len(SubChunk.BlkData)/blocksPerWord; i++ {
		var CorrectIndex = i * 32
		var word uint32
		for b := 0; b < blocksPerWord; b++ {
			word = uint32(SubChunk.BlkData[CorrectIndex+b] >> b & getMask(uint16(blocksPerWord)))
		}
		ReturnPallet[i] = word
	}

	// The code below is a deamon from hell and i hate it.
	// It does work though so i guess thats good.
	// Just dont look at it too hard or it will break.
	// It is a hacky way to convert a uint32 slice to a byte slice.
	var byteslice []byte
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&ReturnPallet))
	header.Len *= 4
	header.Cap *= 4
	byteslice = *(*[]byte)(unsafe.Pointer(&header))
	return byteslice
}

// This function wraps some calls to internal functions so you dont have to do it yourself
func PackBlocksFromBlockData(SubChunk *LevelSubChunk) []byte {
	return packBlocksPN(SubChunk, calculateBlocksPerWord(SubChunk.BitsPerBlock))
}
