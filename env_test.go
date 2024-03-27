package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	const key = "env_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })
	require.Equal(t, "", os.Getenv("key"))

	require.Equal(t, "1234", String(key, "1234"))

	os.Setenv("env_test_key", "3456")
	require.Equal(t, "3456", String(key, "1234"))

	// explicitly set to empty string
	os.Setenv("env_test_key", "")
	require.Equal(t, "", String(key, "1234"))
}

func TestBool(t *testing.T) {
	const key = "env_bool_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })

	require.Equal(t, "", os.Getenv(key))

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
	const key = "env_int_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 1234, Int(key, 1234))

	os.Setenv(key, "3456")
	require.Equal(t, 3456, Int(key, 1234))

	// explicitly set to empty string
	os.Setenv(key, "")
	require.Equal(t, 1234, Int(key, 1234))

}

func TestInt64(t *testing.T) {
	const key = "env_int64_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, int64(9223372036854775807), Int64(key, 9223372036854775807))

	os.Setenv(key, "3456")
	require.Equal(t, int64(3456), Int64(key, 9223372036854775807))

	// explicitly set to empty string
	os.Setenv(key, "")
	require.Equal(t, int64(9223372036854775807), Int64(key, 9223372036854775807))
}

func TestDuration(t *testing.T) {
	const key = "env_duration_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })

	defaultValue := 300 * time.Second

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 300.0, Duration(key, defaultValue).Seconds())

	os.Setenv(key, "600s")
	require.Equal(t, 600.0, Duration(key, defaultValue).Seconds())

	// explicitly set to empty string
	os.Setenv(key, "")
	require.Equal(t, 300.0, Duration(key, defaultValue).Seconds())
}

func TestFloat64(t *testing.T) {
	const key = "env_float64_test_key"
	t.Cleanup(func() { os.Unsetenv(key) })

	require.Equal(t, "", os.Getenv(key))

	require.Equal(t, 123.4, Float64(key, 123.4))

	os.Setenv(key, "345.6")
	require.Equal(t, 345.6, Float64(key, 123.4))

	// explicitly set to empty string
	os.Setenv(key, "")
	require.Equal(t, 123.4, Float64(key, 123.4))
}
