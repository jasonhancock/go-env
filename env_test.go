package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	const key = "env_test_key"
	require.Equal(t, "", os.Getenv("key"))

	require.Equal(t, "1234", String(key, "1234"))

	t.Setenv("env_test_key", "3456")
	require.Equal(t, "3456", String(key, "1234"))

	// explicitly set to empty string
	t.Setenv("env_test_key", "")
	require.Equal(t, "", String(key, "1234"))
}

func TestBool(t *testing.T) {
	const key = "env_bool_test_key"

	require.Equal(t, "", os.Getenv(key))

	require.True(t, Bool(key, true))

	t.Setenv(key, "3456")
	require.True(t, Bool(key, true))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.True(t, Bool(key, true))

	t.Setenv(key, "1")
	require.True(t, Bool(key, false))

	t.Setenv(key, "T")
	require.True(t, Bool(key, false))

	t.Setenv(key, "TRUE")
	require.True(t, Bool(key, false))

	t.Setenv(key, "TrUe")
	require.True(t, Bool(key, false))
}

func TestInt(t *testing.T) {
	const key = "env_int_test_key"

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 1234, Int(key, 1234))

	t.Setenv(key, "3456")
	require.Equal(t, 3456, Int(key, 1234))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, 1234, Int(key, 1234))
}

func TestInt64(t *testing.T) {
	const key = "env_int64_test_key"

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, int64(9223372036854775807), Int64(key, 9223372036854775807))

	t.Setenv(key, "3456")
	require.Equal(t, int64(3456), Int64(key, 9223372036854775807))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, int64(9223372036854775807), Int64(key, 9223372036854775807))
}

func TestDuration(t *testing.T) {
	const key = "env_duration_test_key"

	defaultValue := 300 * time.Second

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 300.0, Duration(key, defaultValue).Seconds())

	t.Setenv(key, "600s")
	require.Equal(t, 600.0, Duration(key, defaultValue).Seconds())

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, 300.0, Duration(key, defaultValue).Seconds())
}

func TestFloat64(t *testing.T) {
	const key = "env_float64_test_key"

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 123.4, Float64(key, 123.4))

	t.Setenv(key, "345.6")
	require.Equal(t, 345.6, Float64(key, 123.4))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, 123.4, Float64(key, 123.4))
}

func TestUint8(t *testing.T) {
	const key = "env_uint8_test_key"
	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, uint8(123), Uint8(key, 123))

	t.Setenv(key, "124")
	require.Equal(t, uint8(124), Uint8(key, 123))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, uint8(123), Uint8(key, 123))

	// set out of bounds
	t.Setenv(key, "256")
	require.Equal(t, uint8(123), Uint8(key, 123))
}

func TestUint16(t *testing.T) {
	const key = "env_uint16_test_key"
	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, uint16(123), Uint16(key, 123))

	t.Setenv(key, "124")
	require.Equal(t, uint16(124), Uint16(key, 123))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, uint16(123), Uint16(key, 123))

	// set out of bounds
	t.Setenv(key, "65536")
	require.Equal(t, uint16(123), Uint16(key, 123))
}

func TestUint32(t *testing.T) {
	const key = "env_uint32_test_key"
	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, uint32(123), Uint32(key, 123))

	t.Setenv(key, "124")
	require.Equal(t, uint32(124), Uint32(key, 123))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, uint32(123), Uint32(key, 123))

	// set out of bounds
	t.Setenv(key, "4294967296")
	require.Equal(t, uint32(123), Uint32(key, 123))
}

func TestUint64(t *testing.T) {
	const key = "env_uint64_test_key"
	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, uint64(123), Uint64(key, 123))

	t.Setenv(key, "124")
	require.Equal(t, uint64(124), Uint64(key, 123))

	// explicitly set to empty string
	t.Setenv(key, "")
	require.Equal(t, uint64(123), Uint64(key, 123))

	// set out of bounds
	t.Setenv(key, "18446744073709551616")
	require.Equal(t, uint64(123), Uint64(key, 123))
}
