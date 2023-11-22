package mcpegsde

import (
	"fmt"
	"strconv"
)

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

// A const array of all the key types used for quick lookup
var (
	KeyTypeArray = [16]uint8{
		Data3D,
		VersionNew,
		Data2D,
		Data2DLegacy,
		SubChunkTerrain,
		LegacyTerrain,
		BlockEntity,
		Entity,
		PendingTicks,
		BlockExtraData,
		BiomeState,
		FinalizedState,
		BorderBlocks,
		HardCodedSpawnAreas,
		Checksums,
		VersionOld,
	}
)

type LDBKey struct {
	Dimension int32
	X         int32
	Z         int32
	Type      uint8
	Y         int8
}

func (k *LDBKey) Key() []byte {
	var returnKey []byte
	returnKey = append(returnKey, Int32ToBytes(k.X)...)
	returnKey = append(returnKey, Int32ToBytes(k.Z)...)
	if k.Dimension != Overworld {
		returnKey = append(returnKey, Int32ToBytes(k.Dimension)...)
	}
	returnKey = append(returnKey, k.Type)
	// copy the raw bits in the int8 over to the byte
	returnKey = append(returnKey, byte(k.Y))
	return returnKey
}

func (k *LDBKey) KeyString() string {
	if k.Dimension == Overworld {
		return fmt.Sprintf("%032s | %032s | %08s | %08s", strconv.FormatInt(int64(k.X), 2), strconv.FormatInt(int64(k.Z), 2), strconv.FormatInt(int64(k.Type), 2), strconv.FormatInt(int64(k.Y), 2))
	} else {
		return fmt.Sprintf("%032s | %032s | %032s | %08s | %08s", strconv.FormatInt(int64(k.X), 2), strconv.FormatInt(int64(k.Z), 2), strconv.FormatInt(int64(k.Dimension), 2), strconv.FormatInt(int64(k.Type), 2), strconv.FormatInt(int64(k.Y), 2))
	}
}

func KeyTypeToString(key uint8) string {
	switch key {
	case Data3D:
		return "Data3D"
	case VersionNew:
		return "VersionNew"
	case Data2D:
		return "Data2D"
	case Data2DLegacy:
		return "Data2DLegacy"
	case SubChunkTerrain:
		return "SubChunkTerrain"
	case LegacyTerrain:
		return "LegacyTerrain"
	case BlockEntity:
		return "BlockEntity"
	case Entity:
		return "Entity"
	case PendingTicks:
		return "PendingTicks"
	case BlockExtraData:
		return "BlockExtraData"
	case BiomeState:
		return "BiomeState"
	case FinalizedState:
		return "FinalizedState"
	case BorderBlocks:
		return "BorderBlocks"
	case HardCodedSpawnAreas:
		return "HardCodedSpawnAreas"
	case Checksums:
		return "Checksums"
	case VersionOld:
		return "VersionOld"
	default:
		return "Unknown"
	}
}
