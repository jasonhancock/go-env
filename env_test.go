package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	val := os.Getenv("env_test_key")
	require.Equal(t, "", val)

	val = String("env_test_key", "1234")
	require.Equal(t, "1234", "1234", val)

	os.Setenv("env_test_key", "3456")
	val = String("env_test_key", "1234")
	require.Equal(t, "3456", val)

	// explicitly set to empty string
	os.Setenv("env_test_key", "")
	val = String("env_test_key", "1234")
	require.Equal(t, "", val)
}

func TestBool(t *testing.T) {
	key := "env_bool_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	require.True(t, Bool(key, true))

	os.Setenv(key, "3456")
	require.True(t, Bool(key, true))

	// explicitly set to empty string
	os.Setenv(key, "")
	require.True(t, Bool(key, true))

	os.Setenv(key, "1")
	require.True(t, Bool(key, false))

	os.Setenv(key, "T")
	require.True(t, Bool(key, false))

	os.Setenv(key, "TRUE")
	require.True(t, Bool(key, false))

	os.Setenv(key, "TrUe")
	require.True(t, Bool(key, false))
}

func TestInt(t *testing.T) {
	key := "env_int_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := Int(key, 1234)
	require.Equal(t, 1234, val)

	os.Setenv(key, "3456")
	val = Int(key, 1234)
	require.Equal(t, 3456, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = Int(key, 1234)
	require.Equal(t, 1234, val)
}

func TestDuration(t *testing.T) {
	key := "env_duration_test_key"
	defaultValue := 300 * time.Second

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := Duration(key, defaultValue)
	require.Equal(t, 300.0, val.Seconds())

	os.Setenv(key, "600s")
	val = Duration(key, defaultValue)
	require.Equal(t, 600.0, val.Seconds())

	// explicitly set to empty string
	os.Setenv(key, "")
	val = Duration(key, defaultValue)
	require.Equal(t, 300.0, val.Seconds())
}

func TestFloat64(t *testing.T) {
	key := "env_float64_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := Float64(key, 123.4)
	require.Equal(t, 123.4, val)

	os.Setenv(key, "345.6")
	val = Float64(key, 123.4)
	require.Equal(t, 345.6, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = Float64(key, 123.4)
	require.Equal(t, 123.4, val)
}
