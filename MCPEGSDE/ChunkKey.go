package mcpegsde

// Dim enum for dimension
const (
	Overworld = iota
	Nether
	TheEnd
)

// Key type
const (
	Data3D              uint8 = 43  // 0x2b (+)
	VersionNew          uint8 = 44  // 0x2c (,)
	Data2D              uint8 = 45  // 0x2d (-), height map + biomes
	Data2DLegacy        uint8 = 46  // 0x2e (.)
	SubChunkTerrain     uint8 = 47  // 0x2f (/)
	LegacyTerrain       uint8 = 48  // 0x30 (0)
	BlockEntity         uint8 = 49  // 0x31 (1)
	Entity              uint8 = 50  // 0x32 (2)
	PendingTicks        uint8 = 51  // 0x33 (3)
	BlockExtraData      uint8 = 52  // 0x34 (4)
	BiomeState          uint8 = 53  // 0x35 (5)
	FinalizedState      uint8 = 54  // 0x36 (6)
	BorderBlocks        uint8 = 56  // Education Edition Feature
	HardCodedSpawnAreas uint8 = 57  // 0x39 (8)
	Checksums           uint8 = 59  // 0x3b (;)
	VersionOld          uint8 = 118 // 0x76 (v)
)

type LDBKey struct {
	Dimension int32
	X         int32
	Z         int32
	Type      uint8
	Y         uint8
}

func (k *LDBKey) Key() []byte {
	var returnKey []byte
	returnKey = append(returnKey, Int32ToBytes(k.X)...)
	returnKey = append(returnKey, Int32ToBytes(k.Z)...)
	if k.Dimension != Overworld {
		returnKey = append(returnKey, Int32ToBytes(k.Dimension)...)
	}
	returnKey = append(returnKey, k.Type)
	returnKey = append(returnKey, k.Y)
	return returnKey
}
