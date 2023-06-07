package device

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/samber/lo"
)

type Device struct {
	Manufacturer string `json:"manufacturer"`
	DeviceModel  string `json:"device_model"`
	Os           string `json:"os"`
	Version      int64  `json:"version"`
	UserAgent    string `json:"user-agent"`
}

type DeviceDataset struct {
	devices []*Device
}

func NewDeviceDataset() *DeviceDataset {
	inst := &DeviceDataset{}

	_deviceOnce.Do(func() {
		var devices []*Device
		err := json.Unmarshal([]byte(_deviceSet), &devices)
		if err != nil {
			panic("devies set unmarshal fail")
		}

		for _, d := range devices {
			d.ua()
		}

		inst.devices = devices
	})

	return inst
}

func (inst *DeviceDataset) All() []*Device {
	return inst.devices
}

func (inst *DeviceDataset) Random() *Device {
	return lo.Sample[*Device](inst.devices)
}

func (inst *Device) ua() {
	format := "Mozilla/5.0 (Linux; Android %d; %s Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/%s Mobile Safari/537.36"

	ua := fmt.Sprintf(format, inst.Version, inst.DeviceModel, inst.buildNumber(), inst.chromeVersion())

	inst.UserAgent = ua
}

func (inst *Device) chromeVersion() (chromeVersion string) {
	versions := []string{
		"Chrome 94.0.4606.61",
		"Chrome 93.0.4577.82",
		"Chrome 92.0.4515.131",
		"Chrome 91.0.4472.124",
		"Chrome 90.0.4430.93",
		"Chrome 89.0.4389.114",
		"Chrome 88.0.4324.182",
		"Chrome 87.0.4280.141",
		"Chrome 86.0.4240.198",
		"Chrome 85.0.4183.121",
		"Chrome 84.0.4147.135",
		"Chrome 83.0.4103.116",
		"Chrome 82.0.4085.28",
		"Chrome 81.0.4044.138",
		"Chrome 80.0.3987.163",
		"Chrome 79.0.3945.130",
		"Chrome 78.0.3904.108",
		"Chrome 77.0.3865.120",
		"Chrome 76.0.3809.132",
		"Chrome 75.0.3770.142",
		"Chrome 74.0.3729.169",
		"Chrome 73.0.3683.103",
		"Chrome 72.0.3626.121",
		"Chrome 71.0.3578.98",
		"Chrome 70.0.3538.110",
		"Chrome 69.0.3497.100",
		"Chrome 68.0.3440.106",
		"Chrome 67.0.3396.99",
		"Chrome 66.0.3359.181",
		"Chrome 65.0.3325.181",
		"Chrome 64.0.3282.167",
		"Chrome 63.0.3239.132",
		"Chrome 62.0.3202.94",
		"Chrome 61.0.3163.100",
		"Chrome 60.0.3112.113",
	}

	chromeVersion = lo.Sample[string](versions)

	return
}

func (inst *Device) buildNumber() (buildNumber string) {
	t := inst.randomTimeBefore(time.Now()).Format("060102")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sub := r.Intn(10) + 1

	buildNumber = fmt.Sprintf("%s.00%d", t, sub)

	return
}

func (inst *Device) randomTimeBefore(t time.Time) time.Time {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	daysToSubtract := r.Intn(365*4) + 150

	return t.Add(time.Duration(-daysToSubtract) * time.Hour * 24)
}

//go:embed device.json
var _deviceSet string
var _deviceOnce sync.Once
