package mcpegsde

import "math"

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

func WriteBitsToIndexMap(blockData *[]uint16, seperatedRawData []byte, bitsPerBlock uint8) {
	var indices []uint16

	bitsPerWord := math.Floor(float64(32 / bitsPerBlock))
	wordsPerChunk := math.Ceil(float64(4096 / bitsPerWord))

	for wordi := 0; wordi < int(wordsPerChunk); wordi++ {
		word := BytesToInt32[uint32](seperatedRawData[wordi*4 : (wordi+1)*4])
		for i := 0; i < int(bitsPerWord); i++ {
			indices = append(indices, uint16(word>>(i*int(bitsPerBlock))))
		}
	}
}
