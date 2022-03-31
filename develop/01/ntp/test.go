package ntp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const server = "0.ru.pool.ntp.org"

func TestGetTime(t *testing.T) {
	time, err := GetTime(server)
	require.NoError(t, err)
	require.NotEmpty(t, time)
}

func TestGetTimeInvalid(t *testing.T) {
	_, err := GetTime("bad server")
	require.Error(t, err)
}
