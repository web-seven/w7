package types

const (
	// ModuleName defines the module name
	ModuleName = "wsvn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wsvn"
)

var (
	ParamsKey = []byte("p_wsvn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
