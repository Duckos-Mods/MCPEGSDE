package mcpegsde

type LevelSubChunk struct {
	Y                     int8
	BlkData               []uint16 // We will use these as indices for the block palette
	BlkPals               []BlockPaletteEntry
	SubChunkVersion       uint8
	PalletType            uint8 // Should allways be Persistant
	BitsPerBlock          uint8
	NumberOfStorageLayers uint8 // Idk what this is
	startOffset           uint8 // This is the offset from the start of the subchunk to the start of the block data
}

type LevelChunk struct {
	X, Z         int32
	SubChunks    [384 / 16]LevelSubChunk
	Dimension    int32
	ChunkVersion uint8
}

// GetSubChunk returns a pointer to the subchunk at Y, or nil if it doesn't exist
func (c *LevelChunk) GetSubChunk(Y int8) *LevelSubChunk {
	// Check if Y is less than -4 or greater than 20 (the range of subchunks) and return nil if it is
	if Y < -4 || Y > 19 {
		return nil
	}
	return &c.SubChunks[Y+4]
}

func (c *LevelSubChunk) SampleNextBlockIndexBassedOnBitsPerBlock(rawChunkData []byte, index int) uint16 {
	var returnIndex uint16
	/*
		We need to get the correct number of bits from the raw data.
		The index is the index of the block in the subchunk.
		We need to get the correct number of bits from the raw data.
		i have no idea how to do this
	*/

	return returnIndex
}

func ConvertLDBEntryToLevelSubChunk(rawChunkData []byte) (*LevelSubChunk, error) {
	var subChunk LevelSubChunk
	subChunk.Y = int8(rawChunkData[2])
	subChunk.NumberOfStorageLayers = rawChunkData[1]
	subChunk.SubChunkVersion = rawChunkData[0]

	subChunk.PalletType = rawChunkData[3] & 0x01                // we only want the first bit
	subChunk.BitsPerBlock = (rawChunkData[3] & 0b11111110) >> 1 // we want all bits except the first one
	// Write the block index data to the subchunk

	// This code sucks and i hate it. It needs to burn in hell.
	// WriteBitsToIndexMap(&subChunk.BlkData, rawChunkData[4:4096/4], subChunk.BitsPerBlock)

	return &subChunk, nil
}
