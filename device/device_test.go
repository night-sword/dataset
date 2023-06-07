package device

import (
	"fmt"
	"testing"
	"time"
)

func TestDeviceRepo_Random(t *testing.T) {
	dd := NewDeviceDataset()
	device := dd.Random()
	fmt.Println(device)
}

func TestDevice_randomTimeBefore(t *testing.T) {
	tt := NewDeviceDataset().Random().randomTimeBefore(time.Now())
	fmt.Println(tt)
}

func TestDevice_buildNumber(t *testing.T) {
	tt := NewDeviceDataset().Random().buildNumber()
	fmt.Println(tt)
}

func TestDevice_ua(t *testing.T) {
	dd := NewDeviceDataset().Random()
	dd.ua()
	fmt.Println(dd.UserAgent)
}
