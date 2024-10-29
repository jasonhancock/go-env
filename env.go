package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// String returns a value from an environment variable defined by
// key. If key isn't set in the environment, returns defaultValue
func String(key string, defaultValue string) string {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			return pieces[1]
		}
	}
	return defaultValue
}

// Bool returns a value from an environment variable defined by
// key. If key isn't set in the environment, returns defaultValue
func Bool(key string, defaultValue bool) bool {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			result, err := strconv.ParseBool(strings.ToLower(pieces[1]))
			if err != nil {
				return defaultValue
			}
			return result
		}
	}
	return defaultValue
}

// Int returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an int,
// returns defaultValue
func Int(key string, defaultValue int) int {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.Atoi(pieces[1])
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// Int64 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an int64,
// returns defaultValue
func Int64(key string, defaultValue int64) int64 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseInt(pieces[1], 10, 64)
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// Duration returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as a time.Duration,
// returns defaultValue
func Duration(key string, defaultValue time.Duration) time.Duration {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := time.ParseDuration(pieces[1])
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// Float64 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as a float64,
// returns defaultValue
func Float64(key string, defaultValue float64) float64 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseFloat(pieces[1], 64)
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// Uint8 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an uint8,
// returns defaultValue
func Uint8(key string, defaultValue uint8) uint8 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseUint(pieces[1], 10, 8)
			if err != nil {
				return defaultValue
			}
			return uint8(val)
		}
	}
	return defaultValue
}

// Uint16 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an uint16,
// returns defaultValue
func Uint16(key string, defaultValue uint16) uint16 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseUint(pieces[1], 10, 16)
			if err != nil {
				return defaultValue
			}
			return uint16(val)
		}
	}
	return defaultValue
}

// Uint32 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an uint32,
// returns defaultValue
func Uint32(key string, defaultValue uint32) uint32 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseUint(pieces[1], 10, 32)
			if err != nil {
				return defaultValue
			}
			return uint32(val)
		}
	}
	return defaultValue
}

// Uint64 returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an uint64,
// returns defaultValue
func Uint64(key string, defaultValue uint64) uint64 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseUint(pieces[1], 10, 64)
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}
