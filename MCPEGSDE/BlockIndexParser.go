package mcpegsde

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

func GetBlocksFromWordP1(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 32)
	for i := 0; i < 32; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x01)
	}
	return returnBlocks
}

func GetBlocksFromWordP2(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 16)
	for i := 0; i < 16; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x03)
	}
	return returnBlocks
}

func GetBlocksFromWordP3(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 10)
	for i := 0; i < 10; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x07)
	}
	return returnBlocks
}

func GetBlocksFromWordP4(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 8)
	for i := 0; i < 8; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x0f)
	}
	return returnBlocks
}

func GetBlocksFromWordP5(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 6)
	for i := 0; i < 6; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x1f)
	}
	return returnBlocks
}

func GetBlocksFromWordP6(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 5)
	for i := 0; i < 5; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0x3f)
	}
	return returnBlocks
}

func GetBlocksFromWordP8(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 4)
	for i := 0; i < 4; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0xff)
	}
	return returnBlocks
}

func GetBlocksFromWordP16(word uint32) []uint16 {
	var returnBlocks = make([]uint16, 2)
	for i := 0; i < 2; i++ {
		returnBlocks[i] = uint16(word >> uint(i) & 0xffff)
	}
	return returnBlocks
}

func GetBlocksFromBytes(SubChunk *LevelSubChunk, rawData []byte) {
	// cast the raw data to a uint32 array
	var rawWords = make([]uint32, len(rawData)/4)
	for i := 0; i < len(rawData); i += 4 {
		switch SubChunk.BitsPerBlock {
		case Palleted1:
			SubChunk.BlkData = GetBlocksFromWordP1(rawWords[i])
		case Palleted2:
			SubChunk.BlkData = GetBlocksFromWordP2(rawWords[i])
		case Palleted3:
			SubChunk.BlkData = GetBlocksFromWordP3(rawWords[i])
		case Palleted4:
			SubChunk.BlkData = GetBlocksFromWordP4(rawWords[i])
		case Palleted5:
			SubChunk.BlkData = GetBlocksFromWordP5(rawWords[i])
		case Palleted6:
			SubChunk.BlkData = GetBlocksFromWordP6(rawWords[i])
		case Palleted8:
			SubChunk.BlkData = GetBlocksFromWordP8(rawWords[i])
		case Palleted16:
			SubChunk.BlkData = GetBlocksFromWordP16(rawWords[i])
		default:
			panic("What the fuck is this pallet type")
		}
	}
}
