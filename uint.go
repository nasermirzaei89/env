package env

// GetUint extracts uint value from env. If not set, returns default value.
func GetUint(key string, def uint) uint {
	return uint(GetUint64(key, uint64(def)))
}

// MustGetUint extracts uint value from env. If not set, it panics.
func MustGetUint(key string) uint {
	return uint(MustGetUint64(key))
}
