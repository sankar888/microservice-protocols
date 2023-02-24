package datagenerator

import (
	"net"

	"github.com/sankar888/protobuff-golang/device"
	"google.golang.org/protobuf/proto"
)

/**
Data generator generates proto.Message type of datarecords.
This generator should produce same output for same no of records. so we need to have same initial value
*/

type DataGenerator interface {
	Get() proto.Message
	RecordCount() uint64
}

var macAddr, _ = net.ParseMAC("00-B0-D0-63-C2-26")
var supportedFrequencyBands []string = []string{"2.4ghz", "5ghz"}
var operatingFrequencyBand string = supportedFrequencyBands[0]

type DeviceRecordGenerator struct {
	incUInt32 uint32
	incUInt64 uint64
	count     uint64
}

func NewDeviceRecordGenerator() *DeviceRecordGenerator {
	return &DeviceRecordGenerator{
		incUInt32: 1,
		incUInt64: 1,
		count:     0,
	}
}

func (gen *DeviceRecordGenerator) Get() proto.Message {
	var ssidStats *device.SSIDStats = &device.SSIDStats{
		BytesSent:       1000,
		BytesReceived:   1000,
		PacketsSent:     10000,
		PacketsReceived: 10000,
		ErrorsSent:      1,
		ErrorsReceived:  1,
	}

	var ssid *device.SSID = &device.SSID{
		Enable:      true,
		Status:      "UP",
		Alias:       "student",
		Name:        "wifi_2.4ghz_ssid_student",
		LowerLayers: []string{"Device.Wifi.Radio.1"},
		Bssid:       macAddr,
		MacAddress:  macAddr,
		Ssid:        "wifi_2.4ghz_ssid_student",
		Stats:       ssidStats,
	}

	var radioStats *device.RadioStats = &device.RadioStats{
		BytesSent:       1000,
		BytesReceived:   1000,
		PacketsSent:     10000,
		PacketsReceived: 10000,
		ErrorsSent:      1,
		ErrorsReceived:  1,
	}

	var radio *device.Radio = &device.Radio{
		Enable:                  true,
		Status:                  "UP",
		Alias:                   "radio",
		Name:                    "wifi_radio",
		Upstream:                false,
		MaxBitRate:              100,
		SupportedFrequencyBands: supportedFrequencyBands,
		OperatingFrequencyBand:  operatingFrequencyBand,
		Channel:                 16,
		Stats:                   radioStats,
	}

	var wifi *device.WiFi = &device.WiFi{
		RadioNumberOfEntries: 1,
		SsidNumberOfEntries:  1,
		Radio:                []*device.Radio{radio},
		Ssid:                 []*device.SSID{ssid},
	}

	var device *device.Device = &device.Device{
		Wifi: wifi,
	}
	gen.count++
	return device
}

func (gen *DeviceRecordGenerator) RecordCount() uint64 {
	return gen.count
}
