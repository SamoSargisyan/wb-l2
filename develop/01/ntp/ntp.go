package ntp

import (
	"github.com/beevik/ntp"
	"time"
)

func GetTime(host string) (string, error) {
	currentTime, err := ntp.Time(host)
	if err != nil {
		return "", err
	}
	return currentTime.Format(time.UnixDate), nil
}
