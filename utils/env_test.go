package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("MY_ENV", "value")
	defer os.Unsetenv("MY_ENV")

	assert.Equal(t, "value", GetEnv("MY_ENV", "default"))
	assert.Equal(t, "default", GetEnv("MY_ENV_UNKNOWN", "default"))
}

func TestGetDefaultServerUrl(t *testing.T) {
	os.Setenv("PUBLIC_DEFAULT_SERVER_URL", "testgetdefaultserverurl")
	defer os.Unsetenv("PUBLIC_DEFAULT_SERVER_URL")
	assert.Equal(t, "testgetdefaultserverurl", GetDefaultServerUrl())
	os.Setenv("PUBLIC_DEFAULT_SERVER_URL", "different")
	assert.NotEqual(t, "testgetdefaultserverurl", GetDefaultServerUrl())
}

func TestParseS3UrlToURL(t *testing.T) {
	os.Setenv("BUCKET_BASE_URL", "http://test.com/")
	defer os.Unsetenv("BUCKET_BASE_URL")

	path := "s3://stablecog/cc70edec-b6ff-42c5-8726-957bbd8fc212.jpeg"
	parsed, err := ParseS3UrlToURL(path)
	assert.Nil(t, err)
	assert.Equal(t, "http://test.com/cc70edec-b6ff-42c5-8726-957bbd8fc212.jpeg", parsed)
}
