package mcpegsde

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt" // We are using gophertunnel lib bc its made for bedrock edition. I think its a good idea to use it. I'm not sure if it's the best idea.
)

func splitNBT(data []byte) [][]byte {
	var entries [][]byte
	// we need to loop through the data and split it into individual NBT entries on the 0x04 0x00 bytes i think. i might be wrong. im dumb.
	var lastSplit = 0
	// we remove 1 from the index because we are seaking 1 byte ahead of the current byte
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x04 && data[i+1] == 0x00 {
			entries = append(entries, data[lastSplit:i])
			lastSplit = i
		}
	}
	return entries
}

func ExtractSingleNBT(data []byte) (BlockPaletteEntry, error) {
	var entry BlockPaletteEntry
	var err = nbt.Unmarshal(data, &entry)
	return entry, err
}

func ExtractBlockPallet(data []byte, SubChunk *LevelSubChunk) {
	// remove the first byte and use it to determine the length of the pallet
	var palletLength = int(data[0])
	// remove the first 3 bytes from the data
	data = data[3:]
	// resize the pallet to the correct length
	SubChunk.BlkPals = make([]BlockPaletteEntry, palletLength)
	// Split the data into individual NBT entries
	var NBTEntries = splitNBT(data)
	// loop through the NBT entries and extract them
	for i := 0; i < len(NBTEntries); i++ {
		var entry, err = ExtractSingleNBT(NBTEntries[i])
		if err != nil {
			panic(err)
		}
		SubChunk.BlkPals[i] = entry
	}
}
