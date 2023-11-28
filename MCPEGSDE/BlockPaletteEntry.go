package mcpegsde

type BlockPaletteEntry struct {
	Name    string                   `nbt:"name"`
	States  []map[string]interface{} `nbt:"states"`
	Version int64                    `nbt:"version"`
}
