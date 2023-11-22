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
	X, Z      int32
	SubChunks []LevelSubChunk
	Dimension int32
}

// GetSubChunk returns a pointer to the subchunk at Y, or nil if it doesn't exist
// TODO: This is probably not the best way to do this. I'm not sure what the best way to do this is.
// Most likely, I'll have to do some sort of sort then just jump to the correct subchunk.
// But im lazy and this works for now.
func (c *LevelChunk) GetSubChunk(Y int8) *LevelSubChunk {
	for _, subchunk := range c.SubChunks {
		if subchunk.Y == Y {
			return &subchunk
		}
	}
	return nil
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
	WriteBitsToIndexMap(&subChunk.BlkData, rawChunkData[4:4096/4], subChunk.BitsPerBlock)

	return &subChunk, nil
}
